package luaprovider

import (
	"errors"
	"github.com/mangalorg/libmangal"
	lua "github.com/yuin/gopher-lua"
)

type Manga struct {
	Title   string `gluamapper:"title"`
	Anilist string `gluamapper:"anilist"`
	URL     string `gluamapper:"url"`
	ID      string `gluamapper:"id"`
	Cover   string `gluamapper:"cover"`

	table *lua.LTable
}

func (m Manga) IntoLValue() lua.LValue {
	return m.table
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

func (m Manga) Info() libmangal.MangaInfo {
	return libmangal.MangaInfo{
		Title:   m.Title,
		Anilist: m.Anilist,
		URL:     m.URL,
		ID:      m.ID,
		Cover:   m.Cover,
	}
}
