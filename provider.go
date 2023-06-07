package luaprovider

import (
	"bytes"
	"context"
	"fmt"
	"github.com/mangalorg/libmangal"
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

	searchMangas, mangaChapters, chapterPages *lua.LFunction
}

func (p Provider) Info() libmangal.ProviderInfo {
	return *p.info
}

func (p Provider) SearchMangas(
	ctx context.Context,
	log libmangal.LogFunc,
	query string,
) ([]libmangal.Manga, error) {
	log(fmt.Sprintf("Searching mangas for %q", query))

	values, err := p.evalFunction(ctx, p.searchMangas, lua.LString(query))
	if err != nil {
		return nil, err
	}

	var mangas = make([]libmangal.Manga, len(values))
	for i, value := range values {
		log(fmt.Sprintf("Parsing manga #%03d", i+1))

		table, ok := value.(*lua.LTable)
		if !ok {
			// TODO: add more descriptive message
			return nil, fmt.Errorf("table expected")
		}

		var manga Manga
		if err = gluamapper.Map(table, &manga); err != nil {
			return nil, err
		}

		if err = manga.Validate(); err != nil {
			return nil, err
		}

		manga.table = table
		mangas[i] = &manga
	}

	log(fmt.Sprintf("Found %d mangas", len(mangas)))
	return mangas, nil
}

func (p Provider) MangaChapters(
	ctx context.Context,
	log libmangal.LogFunc,
	manga *Manga,
) ([]libmangal.Chapter, error) {
	log(fmt.Sprintf("Fetching chapters for %q", manga.Title))
	values, err := p.evalFunction(ctx, p.mangaChapters, manga.table)
	if err != nil {
		return nil, err
	}

	var chapters = make([]libmangal.Chapter, len(values))
	for i, value := range values {
		log(fmt.Sprintf("Parsing chapter #%04d", i))

		table, ok := value.(*lua.LTable)
		if !ok {
			// TODO: add more descriptive message
			return nil, fmt.Errorf("table expected")
		}

		var chapter Chapter
		if err = gluamapper.Map(table, &chapter); err != nil {
			return nil, err
		}

		if err = chapter.Validate(); err != nil {
			return nil, err
		}

		chapter.table = table
		chapter.manga = manga

		if chapter.Number == "" {
			chapter.Number = strconv.Itoa(i + 1)
		}

		chapters[i] = &chapter
	}

	log(fmt.Sprintf("Found %d chapters", len(chapters)))
	return chapters, nil
}

func (p Provider) ChapterPages(
	ctx context.Context,
	log libmangal.LogFunc,
	chapter *Chapter,
) ([]libmangal.Page, error) {
	log(fmt.Sprintf("Fetching pages for %q", chapter.Title))

	values, err := p.evalFunction(ctx, p.chapterPages, chapter.table)
	if err != nil {
		return nil, err
	}

	var pages = make([]libmangal.Page, len(values))
	for i, value := range values {
		log(fmt.Sprintf("Parsing page #%03d", i+1))

		table, ok := value.(*lua.LTable)
		if !ok {
			// TODO: add more descriptive message
			return nil, fmt.Errorf("table expected")
		}

		var page Page
		if err = gluamapper.Map(table, &page); err != nil {
			return nil, err
		}

		page.fillDefaults()

		if err = page.Validate(); err != nil {
			return nil, err
		}

		pages[i] = &page
	}

	log(fmt.Sprintf("Found %d pages", len(pages)))
	return pages, nil
}

func (p Provider) GetImage(
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
	request, _ := http.NewRequestWithContext(ctx, http.MethodGet, page.URL, nil)

	if page.Headers != nil {
		for key, value := range page.Headers {
			request.Header.Set(key, value)
		}
	}

	if page.Cookies != nil {
		for key, value := range page.Cookies {
			request.AddCookie(&http.Cookie{Name: key, Value: value})
		}
	}

	response, err := p.options.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}

	log("Got response")

	defer response.Body.Close()

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

func (p Provider) evalFunction(
	ctx context.Context,
	fn *lua.LFunction,
	input lua.LValue,
) (output []lua.LValue, err error) {
	p.state.SetContext(ctx)
	err = p.state.CallByParam(lua.P{
		Fn:      fn,
		NRet:    1,
		Protect: true,
	}, input)

	if err != nil {
		return nil, err
	}

	p.
		state.
		CheckTable(-1).
		ForEach(func(_, value lua.LValue) {
			output = append(output, value)
		})

	return
}
