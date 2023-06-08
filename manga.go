package luaprovider

import (
	"errors"
	lua "github.com/yuin/gopher-lua"
)

type Manga struct {
	Title string `gluamapper:"title"`
	URL   string `gluamapper:"url"`
	ID    string `gluamapper:"id"`
	Cover string `gluamapper:"cover"`

	table *lua.LTable
}

func (m Manga) Validate() error {
	if m.ID == "" {
		return errors.New("id must be non-empty")
	}

	if m.Title == "" {
		return errors.New("title must be non-empty")
	}

	return nil
}

func (m Manga) GetTitle() string {
	return m.Title
}

func (m Manga) GetURL() string {
	return m.URL
}

func (m Manga) GetID() string {
	return m.ID
}

func (m Manga) GetCoverURL() string {
	return m.Cover
}
