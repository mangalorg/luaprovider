package luaprovider

import "github.com/mangalorg/luaprovider/lib"

func LuaDoc() string {
	return lib.
		Lib(
			newState(DefaultOptions()),
			lib.DefaultOptions(),
		).
		LuaDoc()
}
