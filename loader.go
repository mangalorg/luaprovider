package luaprovider

import (
	"bytes"
	"context"
	"fmt"
	"github.com/mangalorg/libmangal"
	"github.com/philippgille/gokv"
	"github.com/philippgille/gokv/syncmap"
	lua "github.com/yuin/gopher-lua"
	"net/http"
)

var _ libmangal.ProviderLoader = (*loader)(nil)

type Options struct {
	HTTPClient        *http.Client
	HTTPStoreProvider func() (gokv.Store, error)
	PackagePaths      []string
}

func DefaultOptions() Options {
	return Options{
		HTTPClient: &http.Client{},
		HTTPStoreProvider: func() (gokv.Store, error) {
			return syncmap.NewStore(syncmap.DefaultOptions), nil
		},
	}
}

// NewLoader creates new lua provider loader for the given script.
//
// It won't run the script itself.
func NewLoader(script []byte, info libmangal.ProviderInfo, options Options) (libmangal.ProviderLoader, error) {
	if err := info.Validate(); err != nil {
		return nil, err
	}

	return loader{
		options: options,
		info:    info,
		script:  script,
	}, nil
}

type loader struct {
	options Options
	info    libmangal.ProviderInfo
	script  []byte
}

func (l loader) Info() libmangal.ProviderInfo {
	return l.info
}

func (l loader) String() string {
	return l.info.Name
}

func (l loader) Load(ctx context.Context) (libmangal.Provider, error) {
	provider := &provider{
		info:    l.info,
		options: l.options,
	}

	state, store, err := newState(l.options)
	if err != nil {
		return nil, err
	}

	provider.store = store
	provider.state = state
	provider.state.SetContext(ctx)
	lfunc, err := provider.state.Load(bytes.NewReader(l.script), l.info.Name)
	if err != nil {
		return nil, err
	}

	err = provider.state.CallByParam(lua.P{
		Fn:      lfunc,
		NRet:    0,
		Protect: true,
	})
	if err != nil {
		return nil, err
	}

	for name, fn := range map[string]**lua.LFunction{
		methodSearchMangas:   &provider.fnSearchMangas,
		methodMangaVolumes:   &provider.fnMangaVolumes,
		methodVolumeChapters: &provider.fnVolumeChapters,
		methodChapterPages:   &provider.fnChapterPages,
	} {
		var found bool
		*fn, found = provider.state.GetGlobal(name).(*lua.LFunction)

		if !found {
			return nil, fmt.Errorf("missing function: %s", name)
		}
	}

	return provider, nil
}
