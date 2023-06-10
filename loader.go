package luaprovider

import (
	"bytes"
	"context"
	"fmt"
	"github.com/mangalorg/libmangal"
	"github.com/philippgille/gokv"
	"github.com/philippgille/gokv/syncmap"
	"github.com/pkg/errors"
	lua "github.com/yuin/gopher-lua"
	"gopkg.in/yaml.v3"
	"io"
	"net/http"
)

type Options struct {
	HTTPClient *http.Client
	HTTPStore  gokv.Store
}

func DefaultOptions() Options {
	return Options{
		HTTPClient: &http.Client{},
		HTTPStore:  syncmap.NewStore(syncmap.DefaultOptions),
	}
}

func NewLoader(reader io.Reader, options Options) (libmangal.ProviderLoader[Manga, Volume, Chapter, Page], error) {
	contents, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	info, err := extractInfo(contents)
	if err != nil {
		return nil, err
	}

	return Loader{
		options: options,
		info:    info,
		script:  contents,
	}, nil
}

func extractInfo(script []byte) (libmangal.ProviderInfo, error) {
	var (
		infoLines  [][]byte
		infoPrefix = []byte("--->")
	)

	for _, line := range bytes.Split(script, []byte("\n")) {
		if bytes.HasPrefix(line, infoPrefix) {
			infoLines = append(infoLines, bytes.TrimPrefix(line, infoPrefix))
		}
	}

	info := libmangal.ProviderInfo{}
	err := yaml.Unmarshal(bytes.Join(infoLines, []byte("\n")), &info)

	if err != nil {
		return libmangal.ProviderInfo{}, err
	}

	if err := info.Validate(); err != nil {
		return libmangal.ProviderInfo{}, errors.Wrap(err, "info")
	}

	return info, nil
}

type Loader struct {
	options Options
	info    libmangal.ProviderInfo
	script  []byte
}

func (l Loader) Info() libmangal.ProviderInfo {
	return l.info
}

func (l Loader) Load(ctx context.Context) (libmangal.Provider[Manga, Volume, Chapter, Page], error) {
	provider := Provider{
		info:    l.info,
		options: l.options,
	}

	provider.state = newState(l.options)
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
