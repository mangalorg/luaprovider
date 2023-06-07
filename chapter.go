package luaprovider

import (
	"errors"
	"github.com/mangalorg/libmangal"
	lua "github.com/yuin/gopher-lua"
)

type Chapter struct {
	Title  string
	URL    string
	Number string

	manga *Manga
	table *lua.LTable
}

func (c Chapter) Validate() error {
	if c.Title == "" {
		return errors.New("title must be non-empty")
	}

	return nil
}

func (c Chapter) GetTitle() string {
	return c.Title
}

func (c Chapter) GetURL() string {
	return c.URL
}

func (c Chapter) GetNumber() string {
	return c.Number
}

func (c Chapter) GetManga() libmangal.Manga {
	return c.manga
}
