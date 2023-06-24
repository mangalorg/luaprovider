package headless

import (
	"github.com/go-rod/rod"
	"github.com/mangalorg/luaprovider/util"
	lua "github.com/yuin/gopher-lua"
)

const browserTypeName = libName + "_browser"

func newBrowser(state *lua.LState) int {
	browser := rod.New()

	if err := browser.Connect(); err != nil {
		state.RaiseError(err.Error())
		return 0
	}

	util.Push(state, browser, browserTypeName)
	return 1
}

func browserPage(state *lua.LState) int {
	browser := util.Check[*rod.Browser](state, 1)
	URL := state.CheckString(2)

	page := browser.MustPage(URL)

	if err := page.WaitLoad(); err != nil {
		state.RaiseError(err.Error())
		return 0
	}

	util.Push(state, page, pageTypeName)
	return 1
}
