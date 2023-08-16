package luaprovider

import (
	_ "embed"

	"github.com/mangalorg/luaprovider/lib"
	lua "github.com/yuin/gopher-lua"
)

//go:embed template.lua
var luaTemplate string

//go:embed .luarc.json
var luarcJSON string

// LuaTemplate will generate template for the valid lua script used by this provider.
func LuaTemplate() string {
	return luaTemplate
}

func LuarcJSON() string {
	return luarcJSON
}

// LuaDoc will generate library documentation so that language servers can benefit from it.
//
// It's optimized for the https://github.com/LuaLS/lua-language-server
func LuaDoc() string {
	return lib.
		Lib(
			lua.NewState(),
			lib.DefaultOptions(),
		).
		LuaDoc()
}
