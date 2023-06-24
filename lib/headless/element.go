package headless

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
	"github.com/mangalorg/luaprovider/util"
	lua "github.com/yuin/gopher-lua"
)

const elementTypeName = libName + "_element"

func elementInput(state *lua.LState) int {
	element := util.Check[*rod.Element](state, 1)
	text := state.CheckString(2)

	if err := element.Input(text); err != nil {
		state.RaiseError(err.Error())
	}

	return 0
}

func elementClick(state *lua.LState) int {
	element := util.Check[*rod.Element](state, 1)
	count := state.OptNumber(2, 1)

	if err := element.Click(proto.InputMouseButtonLeft, int(count)); err != nil {
		state.RaiseError(err.Error())
	}

	return 0
}

func elementHTML(state *lua.LState) int {
	element := util.Check[*rod.Element](state, 1)

	HTML, err := element.HTML()
	if err != nil {
		state.RaiseError(err.Error())
		return 0
	}

	state.Push(lua.LString(HTML))
	return 1
}
