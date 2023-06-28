package js

import (
	"github.com/mangalorg/luaprovider/util"
	"github.com/robertkrimen/otto"
	lua "github.com/yuin/gopher-lua"
)

const valueTypeName = libName + "_value"

func vmValueExport(L *lua.LState) int {
	value := util.Check[*otto.Value](L, 1)
	nativeValue, err := value.Export()
	if err != nil {
		L.RaiseError(err.Error())
		return 0
	}

	lvalue, err := util.ToLValue(L, nativeValue)
	if err != nil {
		L.RaiseError(err.Error())
		return 0
	}

	L.Push(lvalue)
	return 1
}

func vmValueIsNull(L *lua.LState) int {
	value := util.Check[*otto.Value](L, 1)
	L.Push(lua.LBool(value.IsNull()))
	return 1
}

func vmValueIsUndefined(L *lua.LState) int {
	value := util.Check[*otto.Value](L, 1)
	L.Push(lua.LBool(value.IsUndefined()))
	return 1
}

func vmValueIsNumber(L *lua.LState) int {
	value := util.Check[*otto.Value](L, 1)
	L.Push(lua.LBool(value.IsNumber()))
	return 1
}

func vmValueIsBoolean(L *lua.LState) int {
	value := util.Check[*otto.Value](L, 1)
	L.Push(lua.LBool(value.IsBoolean()))
	return 1
}

func vmValueIsString(L *lua.LState) int {
	value := util.Check[*otto.Value](L, 1)
	L.Push(lua.LBool(value.IsString()))
	return 1
}

func vmValueIsObject(L *lua.LState) int {
	value := util.Check[*otto.Value](L, 1)
	L.Push(lua.LBool(value.IsObject()))
	return 1
}

func vmValueIsNaN(L *lua.LState) int {
	value := util.Check[*otto.Value](L, 1)
	L.Push(lua.LBool(value.IsNaN()))
	return 1
}

func vmValueIsFunction(L *lua.LState) int {
	value := util.Check[*otto.Value](L, 1)
	L.Push(lua.LBool(value.IsFunction()))
	return 1
}

func vmValueString(L *lua.LState) int {
	value := util.Check[*otto.Value](L, 1)
	L.Push(lua.LString(value.String()))
	return 1
}
