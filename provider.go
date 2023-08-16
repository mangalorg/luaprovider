package luaprovider

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/mangalorg/libmangal"
	"github.com/philippgille/gokv"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/yuin/gluamapper"
	lua "github.com/yuin/gopher-lua"
)

var _ libmangal.Provider = (*provider)(nil)

type provider struct {
	store   gokv.Store
	info    libmangal.ProviderInfo
	options Options
	state   *lua.LState
	logger  *libmangal.Logger

	fnSearchMangas,
	fnMangaVolumes,
	fnVolumeChapters,
	fnChapterPages *lua.LFunction
}

func (p *provider) String() string {
	return p.info.Name
}

func (p *provider) Close() error {
	p.state.Close()
	return p.store.Close()
}

func (p *provider) Info() libmangal.ProviderInfo {
	return p.info
}

type IntoLValue interface {
	IntoLValue() lua.LValue
}

// loadItems will run the given lua function,
// perform type checking and apply conversion function for each item.
func loadItems[Input IntoLValue, Output any](
	ctx context.Context,
	logger *libmangal.Logger,
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
		logger.Log(fmt.Sprintf("Parsing item %d", i))
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

	logger.Log(fmt.Sprintf("Found %d items", len(items)))
	return items, nil
}

// luaString wraps string to make IntoLValue method available.
// Required by the loadItems function.
type luaString string

func (l luaString) IntoLValue() lua.LValue {
	return lua.LString(l)
}

func (p *provider) SetLogger(logger *libmangal.Logger) {
	p.logger = logger
}

func (p *provider) SearchMangas(
	ctx context.Context,
	query string,
) ([]libmangal.Manga, error) {
	p.logger.Log(fmt.Sprintf("Searching mangas with %q", query))

	return loadItems(
		ctx,
		p.logger,
		p.state,
		p.fnSearchMangas,
		func(i int, table *lua.LTable) (libmangal.Manga, error) {
			var manga luaManga
			if err := gluamapper.Map(table, &manga); err != nil {
				return luaManga{}, err
			}

			if manga.ID == "" {
				return luaManga{}, errors.New("id must be non-empty")
			}

			if manga.Title == "" {
				return luaManga{}, errors.New("title must be non-empty")
			}

			manga.table = table
			return manga, nil
		},
		luaString(query),
	)
}

func (p *provider) MangaVolumes(
	ctx context.Context,
	manga libmangal.Manga,
) ([]libmangal.Volume, error) {
	m, ok := manga.(luaManga)
	if !ok {
		return nil, fmt.Errorf("unexpected manga type: %T", manga)
	}

	return p.mangaVolumes(ctx, m)
}

func (p *provider) mangaVolumes(
	ctx context.Context,
	manga luaManga,
) ([]libmangal.Volume, error) {
	p.logger.Log(fmt.Sprintf("Fetching volumes for %q", manga.Title))

	return loadItems(
		ctx,
		p.logger,
		p.state,
		p.fnMangaVolumes,
		func(_ int, table *lua.LTable) (libmangal.Volume, error) {
			var volume luaVolume

			if err := gluamapper.Map(table, &volume); err != nil {
				return luaVolume{}, err
			}

			if volume.Number <= 0 {
				return luaVolume{}, fmt.Errorf("invalid volume number: %d", volume.Number)
			}

			volume.table = table
			volume.manga = &manga
			return volume, nil
		},
		manga,
	)
}

func (p *provider) VolumeChapters(
	ctx context.Context,
	volume libmangal.Volume,
) ([]libmangal.Chapter, error) {
	v, ok := volume.(luaVolume)
	if !ok {
		return nil, fmt.Errorf("unexpected volume type: %T", volume)
	}

	return p.volumeChapters(ctx, v)
}

func (p *provider) volumeChapters(
	ctx context.Context,
	volume luaVolume,
) ([]libmangal.Chapter, error) {
	p.logger.Log(fmt.Sprintf("Fetching chapters for volume %d", volume.Number))

	return loadItems(
		ctx,
		p.logger,
		p.state,
		p.fnVolumeChapters,
		func(i int, table *lua.LTable) (libmangal.Chapter, error) {
			var chapter luaChapter
			if err := gluamapper.Map(table, &chapter); err != nil {
				return luaChapter{}, err
			}

			if chapter.Title == "" {
				return luaChapter{}, errors.New("title must be non-empty")
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

func (p *provider) ChapterPages(
	ctx context.Context,
	chapter libmangal.Chapter,
) ([]libmangal.Page, error) {
	c, ok := chapter.(luaChapter)
	if !ok {
		return nil, fmt.Errorf("unexpected chapter type: %T", chapter)
	}

	return p.chapterPages(ctx, c)
}

func (p *provider) chapterPages(
	ctx context.Context,
	chapter luaChapter,
) ([]libmangal.Page, error) {
	p.logger.Log(fmt.Sprintf("Fetching pages for %q", chapter.Title))

	return loadItems(
		ctx,
		p.logger,
		p.state,
		p.fnChapterPages,
		func(i int, table *lua.LTable) (libmangal.Page, error) {
			var page luaPage
			if err := gluamapper.Map(table, &page); err != nil {
				return luaPage{}, err
			}

			page.chapter = &chapter

			if page.URL == "" {
				return luaPage{}, errors.New("url must be set")
			}

			if page.Extension == "" {
				page.Extension = ".jpg"
			}

			if !fileExtensionRegex.MatchString(page.Extension) {
				return luaPage{}, fmt.Errorf("invalid page extension: %s", page.Extension)
			}

			if page.Headers == nil {
				page.Headers = make(map[string]string)
				page.Headers["Referer"] = page.chapter.URL
				page.Headers["Accept"] = "image/webp,image/apng,image/*,*/*;q=0.8"

				// TODO: generate random user-agent
				page.Headers["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36"
			}

			return page, nil
		},
		chapter,
	)
}

func (p *provider) GetPageImage(
	ctx context.Context,
	page libmangal.Page,
) ([]byte, error) {
	page_, ok := page.(luaPage)
	if !ok {
		return nil, fmt.Errorf("unexpected page type: %T", page)
	}

	return p.getPageImage(ctx, page_)
}

func (p *provider) getPageImage(
	ctx context.Context,
	page luaPage,
) ([]byte, error) {
	p.logger.Log(fmt.Sprintf("Making HTTP GET request for %q", page.URL))
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

	p.logger.Log("Got response")

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	image, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return image, nil
}
