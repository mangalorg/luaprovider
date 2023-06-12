package luaprovider

import (
	"github.com/mangalorg/libmangal"
	lua "github.com/yuin/gopher-lua"
)

type Manga struct {
	Title         string `gluamapper:"title"`
	AnilistSearch string `gluamapper:"anilist_search"`
	URL           string `gluamapper:"url"`
	ID            string `gluamapper:"id"`
	Cover         string `gluamapper:"cover"`
	Banner        string `gluamapper:"banner"`

	table *lua.LTable
}

func (m Manga) String() string {
	return m.Title
}

func (m Manga) SeriesJson() (seriesJson libmangal.SeriesJson, ok bool) {
	return
}

func (m Manga) IntoLValue() lua.LValue {
	return m.table
}

func (m Manga) Info() libmangal.MangaInfo {
	return libmangal.MangaInfo{
		Title:         m.Title,
		AnilistSearch: m.AnilistSearch,
		URL:           m.URL,
		ID:            m.ID,
		Cover:         m.Cover,
		Banner:        m.Banner,
	}
}
