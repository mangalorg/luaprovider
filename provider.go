package luaprovider

import (
	"context"
	"fmt"
	"github.com/mangalorg/libmangal"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/yuin/gluamapper"
	lua "github.com/yuin/gopher-lua"
	"io"
	"net/http"
)

type Provider struct {
	info    libmangal.ProviderInfo
	options Options
	state   *lua.LState

	fnSearchMangas,
	fnMangaVolumes,
	fnVolumeChapters,
	fnChapterPages *lua.LFunction
}

func (p Provider) Info() libmangal.ProviderInfo {
	return p.info
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
) ([]Manga, error) {
	log(fmt.Sprintf("Searching mangas with %q", query))

	return loadItems(
		ctx,
		log,
		p.state,
		p.fnSearchMangas,
		func(i int, table *lua.LTable) (Manga, error) {
			var manga Manga
			if err := gluamapper.Map(table, &manga); err != nil {
				return Manga{}, err
			}

			if manga.ID == "" {
				return Manga{}, errors.New("id must be non-empty")
			}

			if manga.Title == "" {
				return Manga{}, errors.New("title must be non-empty")
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
	manga Manga,
) ([]Volume, error) {
	log(fmt.Sprintf("Fetching volumes for %q", manga.Title))

	return loadItems(
		ctx,
		log,
		p.state,
		p.fnMangaVolumes,
		func(_ int, table *lua.LTable) (Volume, error) {
			var volume Volume

			if err := gluamapper.Map(table, &volume); err != nil {
				return Volume{}, err
			}

			if volume.Number <= 0 {
				return Volume{}, fmt.Errorf("invalid volume number: %d", volume.Number)
			}

			volume.table = table
			volume.manga = &manga
			return volume, nil
		},
		manga,
	)
}

func (p Provider) VolumeChapters(
	ctx context.Context,
	log libmangal.LogFunc,
	volume Volume,
) ([]Chapter, error) {
	log(fmt.Sprintf("Fetching chapters for volume %d", volume.Number))

	return loadItems(
		ctx,
		log,
		p.state,
		p.fnVolumeChapters,
		func(i int, table *lua.LTable) (Chapter, error) {
			var chapter Chapter
			if err := gluamapper.Map(table, &chapter); err != nil {
				return Chapter{}, err
			}

			if chapter.Title == "" {
				return Chapter{}, errors.New("title must be non-empty")
			}

			if chapter.Number == 0 {
				chapter.Number = float32(i)
			}

			chapter.table = table
			chapter.volume = &volume
			return chapter, nil
		},
		volume,
	)
}

func (p Provider) ChapterPages(
	ctx context.Context,
	log libmangal.LogFunc,
	chapter Chapter,
) ([]Page, error) {
	log(fmt.Sprintf("Fetching pages for %q", chapter.Title))

	return loadItems(
		ctx,
		log,
		p.state,
		p.fnChapterPages,
		func(i int, table *lua.LTable) (Page, error) {
			var page Page
			if err := gluamapper.Map(table, &page); err != nil {
				return Page{}, err
			}

			page.chapter = &chapter

			if page.Extension == "" {
				page.Extension = ".jpg"
			}

			if page.Headers == nil {
				page.Headers = make(map[string]string)
				page.Headers["Referer"] = page.chapter.URL
				page.Headers["Accept"] = "image/webp,image/apng,image/*,*/*;q=0.8"

				// TODO: generate random user-agent
				page.Headers["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36"
			}

			if page.URL == "" {
				return Page{}, errors.New("url must be set")
			}

			if !fileExtensionRegex.MatchString(page.Extension) {
				return Page{}, fmt.Errorf("invalid page extension: %s", page.Extension)
			}

			return page, nil
		},
		chapter,
	)
}

func (p Provider) GetPageImage(
	ctx context.Context,
	log libmangal.LogFunc,
	page Page,
) ([]byte, error) {
	log("Getting image for page")

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
	var image []byte

	// check content length
	if response.ContentLength > 0 {
		image = make([]byte, response.ContentLength)
		_, err = io.ReadFull(response.Body, image)
	} else {
		image, err = io.ReadAll(response.Body)
	}

	if err != nil {
		return nil, err
	}

	return image, nil
}
