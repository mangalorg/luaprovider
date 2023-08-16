package luaprovider

import (
	"path/filepath"
	"strings"

	"github.com/mangalorg/luaprovider/lib"
	"github.com/philippgille/gokv"
	"github.com/samber/lo"
	lua "github.com/yuin/gopher-lua"
)

func newState(options Options) (*lua.LState, gokv.Store, error) {
	libs := []lua.LGFunction{
		lua.OpenBase,
		lua.OpenTable,
		lua.OpenString,
		lua.OpenMath,
		lua.OpenPackage,
		lua.OpenIo,
		lua.OpenCoroutine,
		lua.OpenChannel,
	}

	state := lua.NewState(lua.Options{
		SkipOpenLibs: true,
	})

	for _, injectLib := range libs {
		injectLib(state)
	}

	store, err := options.HTTPStoreProvider()
	if err != nil {
		return nil, nil, err
	}

	lib.Preload(state, &lib.Options{
		HTTPClient: options.HTTPClient,
		HTTPStore:  store,
	})

	pkg := state.GetGlobal("package").(*lua.LTable)

	paths := lo.Map(options.PackagePaths, func(path string, _ int) string {
		return filepath.Join(path, "?.lua")
	})

	pkg.RawSetString("path", lua.LString(strings.Join(paths, ";")))

	return state, store, nil
}
