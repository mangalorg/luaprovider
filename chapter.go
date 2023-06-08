package luaprovider

import (
	"errors"
	"github.com/mangalorg/libmangal"
	lua "github.com/yuin/gopher-lua"
)

type Chapter struct {
	Title  string `gluamapper:"title"`
	URL    string `gluamapper:"url"`
	Number string `gluamapper:"number"`

	volume *Volume
	table  *lua.LTable
}

func (c Chapter) IntoLValue() lua.LValue {
	return c.table
}

func (c Chapter) Validate() error {
	if c.Title == "" {
		return errors.New("title must be non-empty")
	}

	return nil
}

func (c Chapter) Info() libmangal.ChapterInfo {
	return libmangal.ChapterInfo{
		Title:      c.Title,
		URL:        c.URL,
		Number:     c.Number,
		VolumeInfo: c.volume.Info,
	}
}
