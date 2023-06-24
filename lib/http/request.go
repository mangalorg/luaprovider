package http

import (
	"bufio"
	"bytes"
	"github.com/mangalorg/luaprovider/util"
	"github.com/philippgille/gokv"
	lua "github.com/yuin/gopher-lua"
	"net/http"
	"net/http/httputil"
	"strings"
)

const requestTypeName = libName + "_request"

func requestNew(L *lua.LState) int {
	method := L.CheckString(1)
	url := L.CheckString(2)
	body := L.OptString(3, "")

	request, err := http.NewRequestWithContext(
		L.Context(),
		method,
		url,
		strings.NewReader(body),
	)
	if err != nil {
		L.RaiseError(err.Error())
	}

	util.Push(L, request, requestTypeName)
	return 1
}

func requestHeader(L *lua.LState) int {
	request := util.Check[*http.Request](L, 1)
	key := L.CheckString(2)

	if L.GetTop() == 3 {
		value := L.CheckString(3)
		request.Header.Set(key, value)
		return 0
	}

	L.Push(lua.LString(request.Header.Get(key)))
	return 1
}

func requestCookie(L *lua.LState) int {
	request := util.Check[*http.Request](L, 1)
	key := L.CheckString(2)

	if L.GetTop() == 3 {
		value := L.CheckString(3)
		cookie := &http.Cookie{Name: key, Value: value}
		request.AddCookie(cookie)
		return 0
	}

	cookie, _ := request.Cookie(key)
	if cookie == nil {
		L.Push(lua.LNil)
		return 1
	}

	L.Push(lua.LString(cookie.Value))
	return 1
}

func requestContentLength(L *lua.LState) int {
	request := util.Check[*http.Request](L, 1)

	if L.GetTop() == 2 {
		value := L.CheckInt64(2)
		request.ContentLength = value
		return 0
	}

	L.Push(lua.LNumber(request.ContentLength))
	return 1
}

func requestSend(L *lua.LState, client *http.Client, store gokv.Store) int {
	request := util.Check[*http.Request](L, 1)

	dumpedRequest, errRequestDump := httputil.DumpRequestOut(request, true)
	dumpedRequestString := string(dumpedRequest)

	if errRequestDump == nil {
		var dumpedResponse []byte

		found, err := store.Get(dumpedRequestString, &dumpedResponse)
		if err != nil {
			_ = store.Delete(dumpedRequestString)
			goto doRequest
		}

		if !found {
			goto doRequest
		}

		response, err := http.ReadResponse(
			bufio.NewReader(bytes.NewReader(dumpedResponse)),
			request,
		)

		if err != nil {
			_ = store.Delete(dumpedRequestString)
			goto doRequest
		}

		util.Push(L, response, responseTypeName)
		return 1
	}

doRequest:
	response, err := client.Do(request)
	if err != nil {
		L.RaiseError(err.Error())
		return 0
	}

	// only cache the response if it was successful
	if errRequestDump == nil && response.StatusCode == http.StatusOK {
		dumpedResponse, err := httputil.DumpResponse(response, true)
		if err != nil {
			goto exit
		}

		_ = store.Set(dumpedRequestString, dumpedResponse)
	}

exit:
	util.Push(L, response, responseTypeName)
	return 1
}
