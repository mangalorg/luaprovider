package headless

import (
	"github.com/go-rod/rod"
	"github.com/mangalorg/luaprovider/util"
	lua "github.com/yuin/gopher-lua"
)

const pageTypeName = libName + "_page"

func pageElement(state *lua.LState) int {
	page := util.Check[*rod.Page](state, 1)
	selector := state.CheckString(2)

	element, err := page.Element(selector)
	if err != nil {
		state.RaiseError(err.Error())
		return 0
	}

	util.Push(state, element, elementTypeName)
	return 1
}

func pageNavigate(state *lua.LState) int {
	page := util.Check[*rod.Page](state, 1)
	URL := state.CheckString(2)

	if err := page.Navigate(URL); err != nil {
		state.RaiseError(err.Error())
	}

	return 0
}

func pageHTML(state *lua.LState) int {
	page := util.Check[*rod.Page](state, 1)
	HTML, err := page.HTML()
	if err != nil {
		state.RaiseError(err.Error())
		return 0
	}

	state.Push(lua.LString(HTML))
	return 1
}
