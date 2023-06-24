package http

import (
	"github.com/mangalorg/luaprovider/util"
	lua "github.com/yuin/gopher-lua"
	"io"
	"net/http"
)

const responseTypeName = libName + "_response"

func responseStatus(L *lua.LState) int {
	response := util.Check[*http.Response](L, 1)
	L.Push(lua.LNumber(response.StatusCode))
	return 1
}

func responseBody(L *lua.LState) int {
	response := util.Check[*http.Response](L, 1)

	var (
		buffer []byte
		err    error
	)

	buffer, err = io.ReadAll(response.Body)
	if err != nil {
		L.RaiseError("failed to read response body: %s", err.Error())
		return 0
	}

	L.Push(lua.LString(buffer))
	return 1
}

func responseHeader(L *lua.LState) int {
	response := util.Check[*http.Response](L, 1)
	key := L.CheckString(2)

	L.Push(lua.LString(response.Header.Get(key)))
	return 1
}

func responseCookies(L *lua.LState) int {
	response := util.Check[*http.Response](L, 1)

	cookies := L.NewTable()
	for _, cookie := range response.Cookies() {
		c := L.NewTable()
		c.RawSetString("name", lua.LString(cookie.Name))
		c.RawSetString("value", lua.LString(cookie.Value))

		cookies.Append(c)
	}

	L.Push(cookies)
	return 1
}

func responseContentLength(L *lua.LState) int {
	response := util.Check[*http.Response](L, 1)
	L.Push(lua.LNumber(response.ContentLength))
	return 1
}
