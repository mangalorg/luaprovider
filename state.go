package luaprovider

import (
	"github.com/mangalorg/luaprovider/lib"
	"github.com/philippgille/gokv/syncmap"
	lua "github.com/yuin/gopher-lua"
	"net/http"
)

func DefaultOptions() *Options {
	return &Options{
		HTTPClient: &http.Client{},
		HTTPStore:  syncmap.NewStore(syncmap.DefaultOptions),
	}
}

func newState(options *Options) *lua.LState {
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

	lib.Preload(state, &lib.Options{
		HTTPClient: options.HTTPClient,
		FS:         options.FS,
		HTTPStore:  options.HTTPStore,
	})

	return state
}
