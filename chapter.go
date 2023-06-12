package luaprovider

import (
	"github.com/mangalorg/libmangal"
	lua "github.com/yuin/gopher-lua"
)

type Chapter struct {
	Title  string  `gluamapper:"title"`
	URL    string  `gluamapper:"url"`
	Number float32 `gluamapper:"number"`

	volume *Volume
	table  *lua.LTable
}

func (c Chapter) String() string {
	return c.Title
}

func (c Chapter) Volume() libmangal.Volume {
	return c.volume
}

func (c Chapter) ComicInfoXml() (comicInfo libmangal.ComicInfoXml, ok bool) {
	return
}

func (c Chapter) IntoLValue() lua.LValue {
	return c.table
}

func (c Chapter) Info() libmangal.ChapterInfo {
	return libmangal.ChapterInfo{
		Title:  c.Title,
		URL:    c.URL,
		Number: c.Number,
	}
}
