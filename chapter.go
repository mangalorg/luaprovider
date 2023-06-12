package luaprovider

import (
	"github.com/mangalorg/libmangal"
	lua "github.com/yuin/gopher-lua"
)

var _ libmangal.Chapter = (*luaChapter)(nil)

type luaChapter struct {
	Title  string  `gluamapper:"title"`
	URL    string  `gluamapper:"url"`
	Number float32 `gluamapper:"number"`

	volume *luaVolume
	table  *lua.LTable
}

func (c luaChapter) String() string {
	return c.Title
}

func (c luaChapter) Volume() libmangal.Volume {
	return c.volume
}

func (c luaChapter) IntoLValue() lua.LValue {
	return c.table
}

func (c luaChapter) Info() libmangal.ChapterInfo {
	return libmangal.ChapterInfo{
		Title:  c.Title,
		URL:    c.URL,
		Number: c.Number,
	}
}
