package lib

import (
	luadoc "github.com/mangalorg/luaprovider/doc"
	"github.com/mangalorg/luaprovider/lib/crypto"
	"github.com/mangalorg/luaprovider/lib/encoding"
	"github.com/mangalorg/luaprovider/lib/headless"
	"github.com/mangalorg/luaprovider/lib/html"
	httpLib "github.com/mangalorg/luaprovider/lib/http"
	"github.com/mangalorg/luaprovider/lib/js"
	"github.com/mangalorg/luaprovider/lib/levenshtein"
	"github.com/mangalorg/luaprovider/lib/regexp"
	"github.com/mangalorg/luaprovider/lib/strings"
	"github.com/mangalorg/luaprovider/lib/time"
	"github.com/mangalorg/luaprovider/lib/urls"
	"github.com/mangalorg/luaprovider/lib/util"
	"github.com/philippgille/gokv"
	"github.com/philippgille/gokv/syncmap"
	"github.com/spf13/afero"
	lua "github.com/yuin/gopher-lua"
	"net/http"
)

type Options struct {
	HTTPClient *http.Client
	HTTPStore  gokv.Store
	FS         afero.Fs
}

func DefaultOptions() *Options {
	return &Options{
		HTTPClient: &http.Client{},
		HTTPStore:  syncmap.NewStore(syncmap.DefaultOptions),
		FS:         afero.NewMemMapFs(),
	}
}

const libName = "sdk"

func Lib(L *lua.LState, options *Options) *luadoc.Lib {
	return &luadoc.Lib{
		Name:        libName,
		Description: `Contains various utilities for making HTTP requests, working with JSON, HTML, and more.`,
		Libs: []*luadoc.Lib{
			regexp.Lib(L),
			strings.Lib(),
			crypto.Lib(L),
			js.Lib(),
			html.Lib(),
			levenshtein.Lib(),
			util.Lib(),
			time.Lib(),
			urls.Lib(),
			encoding.Lib(L),
			headless.Lib(),
			httpLib.Lib(httpLib.LibOptions{
				HTTPClient: options.HTTPClient,
				HTTPStore:  options.HTTPStore,
			}),
		},
	}
}

func Preload(L *lua.LState, options *Options) {
	lib := Lib(L, options)
	L.PreloadModule(lib.Name, lib.Loader())
}
