package luaprovider

import (
	"encoding/json"
	"github.com/mangalorg/libmangal"
	lua "github.com/yuin/gopher-lua"
)

var _ libmangal.Manga = (*luaManga)(nil)

type luaManga struct {
	Title         string `gluamapper:"title"`
	AnilistSearch string `gluamapper:"anilist_search"`
	URL           string `gluamapper:"url"`
	ID            string `gluamapper:"id"`
	Cover         string `gluamapper:"cover"`
	Banner        string `gluamapper:"banner"`

	table *lua.LTable
}

func (m luaManga) String() string {
	return m.Title
}

func (m luaManga) IntoLValue() lua.LValue {
	return m.table
}

func (m luaManga) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Info())
}

func (m luaManga) Info() libmangal.MangaInfo {
	return libmangal.MangaInfo{
		Title:         m.Title,
		AnilistSearch: m.AnilistSearch,
		URL:           m.URL,
		ID:            m.ID,
		Cover:         m.Cover,
		Banner:        m.Banner,
	}
}
