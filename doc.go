package luaprovider

import (
	_ "embed"
	"github.com/mangalorg/luaprovider/lib"
)

//go:embed template.lua
var luaTemplate string

func LuaTemplate() string {
	return luaTemplate
}

func LuaDoc() string {
	return lib.
		Lib(
			newState(DefaultOptions()),
			lib.DefaultOptions(),
		).
		LuaDoc()
}
