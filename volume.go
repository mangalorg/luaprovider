package luaprovider

import (
	"github.com/mangalorg/libmangal"
	lua "github.com/yuin/gopher-lua"
)

type Volume struct {
	Number int `gluamapper:"number"`

	manga *Manga
	table *lua.LTable
}

func (v Volume) IntoLValue() lua.LValue {
	return v.table
}

func (v Volume) Info() libmangal.VolumeInfo {
	return libmangal.VolumeInfo{
		Number:    v.Number,
		MangaInfo: v.manga.Info,
	}
}
