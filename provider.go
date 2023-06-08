package luaprovider

import (
	"bytes"
	"context"
	"fmt"
	"github.com/mangalorg/libmangal"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/yuin/gluamapper"
	lua "github.com/yuin/gopher-lua"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type Provider struct {
	info    *libmangal.ProviderInfo
	options *Options
	state   *lua.LState

	fnSearchMangas,
	fnMangaVolumes,
	fnVolumeChapters,
	fnChapterPages *lua.LFunction
}

func (p Provider) Info() libmangal.ProviderInfo {
	return *p.info
}

type IntoLValue interface {
	IntoLValue() lua.LValue
}

func loadItems[Input IntoLValue, Output any](
	ctx context.Context,
	log libmangal.LogFunc,
	state *lua.LState,
	lfunc *lua.LFunction,
	convert func(int, *lua.LTable) (Output, error),
	args ...Input,
) ([]Output, error) {
	state.SetContext(ctx)
	err := state.CallByParam(lua.P{
		Fn:      lfunc,
		NRet:    1,
		Protect: true,
	}, lo.Map(args, func(arg Input, _ int) lua.LValue {
		return arg.IntoLValue()
	})...)

	if err != nil {
		return nil, err
	}

	output := state.CheckTable(-1)

	var values []lua.LValue
	output.ForEach(func(_ lua.LValue, value lua.LValue) {
		values = append(values, value)
	})

	var items = make([]Output, len(values))
	for i, value := range values {
		log(fmt.Sprintf("Parsing item %d", i))
		table, ok := value.(*lua.LTable)
		if !ok {
			return nil, errors.Wrapf(fmt.Errorf("expected table, got %s", value.Type().String()), "parsing item %d", i)
		}

		item, err := convert(i, table)
		if err != nil {
			return nil, errors.Wrapf(err, "parsing item %d", i)
		}

		items[i] = item
	}

	log(fmt.Sprintf("Found %d items", len(items)))
	return items, nil
}

type luaString string

func (l luaString) IntoLValue() lua.LValue {
	return lua.LString(l)
}

func (p Provider) SearchMangas(
	ctx context.Context,
	log libmangal.LogFunc,
	query string,
) ([]*Manga, error) {
	log(fmt.Sprintf("Searching mangas with %q", query))

	return loadItems(
		ctx,
		log,
		p.state,
		p.fnSearchMangas,
		func(i int, table *lua.LTable) (*Manga, error) {
			var manga *Manga
			if err := gluamapper.Map(table, &manga); err != nil {
				return nil, err
			}

			if err := manga.Validate(); err != nil {
				return nil, err
			}

			manga.table = table
			return manga, nil
		},
		luaString(query),
	)
}

func (p Provider) MangaVolumes(
	ctx context.Context,
	log libmangal.LogFunc,
	manga *Manga,
) ([]*Volume, error) {
	log(fmt.Sprintf("Fetching volumes for %q", manga.Title))

	return loadItems(
		ctx,
		log,
		p.state,
		p.fnMangaVolumes,
		func(_ int, table *lua.LTable) (*Volume, error) {
			var volume *Volume

			if err := gluamapper.Map(table, &volume); err != nil {
				return nil, err
			}

			if volume.Number <= 0 {
				return nil, fmt.Errorf("invalid volume number: %d", volume.Number)
			}

			volume.table = table
			volume.manga = manga
			return volume, nil
		},
		manga,
	)
}

func (p Provider) VolumeChapters(
	ctx context.Context,
	log libmangal.LogFunc,
	volume *Volume,
) ([]*Chapter, error) {
	log(fmt.Sprintf("Fetching chapters for volume %d", volume.Number))

	return loadItems(
		ctx,
		log,
		p.state,
		p.fnVolumeChapters,
		func(i int, table *lua.LTable) (*Chapter, error) {
			var chapter *Chapter
			if err := gluamapper.Map(table, &chapter); err != nil {
				return nil, err
			}

			if err := chapter.Validate(); err != nil {
				return nil, err
			}

			if chapter.Number == "" {
				chapter.Number = strconv.Itoa(i + 1)
			}

			chapter.table = table
			chapter.volume = volume
			return chapter, nil
		},
		volume,
	)
}

func (p Provider) ChapterPages(
	ctx context.Context,
	log libmangal.LogFunc,
	chapter *Chapter,
) ([]*Page, error) {
	log(fmt.Sprintf("Fetching pages for %q", chapter.Title))

	return loadItems(
		ctx,
		log,
		p.state,
		p.fnChapterPages,
		func(i int, table *lua.LTable) (*Page, error) {
			var page *Page
			if err := gluamapper.Map(table, &page); err != nil {
				return nil, err
			}

			page.chapter = chapter
			page.fillDefaults()

			if err := page.Validate(); err != nil {
				return nil, err
			}

			return page, nil
		},
		chapter,
	)
}

func (p Provider) GetPageImage(
	ctx context.Context,
	log libmangal.LogFunc,
	page *Page,
) (io.Reader, error) {
	log("Getting image for page")

	if page.Data != "" {
		log("Page already contains image, returning")
		return strings.NewReader(page.Data), nil
	}

	log(fmt.Sprintf("Making HTTP GET request for %q", page.URL))
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, page.URL, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range page.Headers {
		request.Header.Set(key, value)
	}

	for key, value := range page.Cookies {
		request.AddCookie(&http.Cookie{Name: key, Value: value})
	}

	response, err := p.options.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	log("Got response")

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	log("Everything is OK, reading")
	var buffer []byte

	// check content length
	if response.ContentLength > 0 {
		buffer = make([]byte, response.ContentLength)
		_, err = io.ReadFull(response.Body, buffer)
	} else {
		buffer, err = io.ReadAll(response.Body)
	}

	if err != nil {
		return nil, err
	}

	return bytes.NewReader(buffer), nil
}
