package urls

import (
	"net/url"

	luadoc "github.com/mangalorg/gopher-luadoc"
	"github.com/mangalorg/luaprovider/util"
	lua "github.com/yuin/gopher-lua"
)

const (
	libName        = "urls"
	valuesTypeName = libName + "_values"
	urlTypeName    = libName + "_url"
)

func Lib() *luadoc.Lib {
	classValues := &luadoc.Class{
		Name:        valuesTypeName,
		Description: "Values maps a string key to a list of values. It is typically used for query parameters and form values. Unlike in the `http.header` map, the keys in a `values` map are case-sensitive.",
		Methods: []*luadoc.Method{
			{
				Name:        "add",
				Description: "Adds the key and value to the values. It appends to any existing values associated with key.",
				Value:       urlValuesAdd,
				Params: []*luadoc.Param{
					{
						Name:        "key",
						Description: "The key to add. It must not be empty.",
						Type:        luadoc.String,
					},
					{
						Name:        "value",
						Description: "The value to add. It must not be empty.",
						Type:        luadoc.String,
					},
				},
			},
			{
				Name:        "set",
				Description: "Sets the key to value. It replaces any existing values associated with key.",
				Value:       urlValuesSet,
				Params: []*luadoc.Param{
					{
						Name:        "key",
						Description: "The key to add. It must not be empty.",
						Type:        luadoc.String,
					},
					{
						Name:        "value",
						Description: "The value to add. It must not be empty.",
						Type:        luadoc.String,
					},
				},
			},
			{
				Name:        "get",
				Description: "Gets the first value associated with the given key. If there are no values associated with the key, Get returns \"\".",
				Value:       urlValuesGet,
				Params: []*luadoc.Param{
					{
						Name:        "key",
						Description: "The key to get.",
						Type:        luadoc.String,
					},
				},
				Returns: []*luadoc.Param{
					{
						Name:        "value",
						Description: "The first value associated with the given key.",
						Type:        luadoc.String,
					},
				},
			},
			{
				Name:        "has",
				Description: "Returns true if the values contains the specified key, false otherwise.",
				Value:       urlValuesHas,
				Params: []*luadoc.Param{
					{
						Name:        "key",
						Description: "The key to check.",
						Type:        luadoc.String,
					},
				},
				Returns: []*luadoc.Param{
					{
						Name:        "has",
						Description: "True if the values contains the specified key, false otherwise.",
						Type:        luadoc.Boolean,
					},
				},
			},
			{
				Name:        "del",
				Description: "Deletes the values associated with key.",
				Value:       urlValuesDel,
				Params: []*luadoc.Param{
					{
						Name:        "key",
						Description: "The key to delete.",
						Type:        luadoc.String,
					},
				},
			},
			{
				Name:        "string",
				Description: "Encodes the values into \"URL encoded\" form sorted by key.",
				Value:       urlValuesString,
				Returns: []*luadoc.Param{
					{
						Name:        "encoded",
						Description: "The URL encoded form of the values.",
						Type:        luadoc.String,
					},
				},
			},
			{
				Name:        "parse",
				Description: "Creates a values from the URL encoded form. It is the inverse operation of string.",
				Value:       urlValuesParse,
				Params: []*luadoc.Param{
					{
						Name:        "encoded",
						Description: "The URL encoded form of the values.",
						Type:        luadoc.String,
					},
				},
				Returns: []*luadoc.Param{
					{
						Name:        "values",
						Description: "The values created from the URL encoded form.",
						Type:        valuesTypeName,
					},
				},
			},
		},
	}

	classURL := &luadoc.Class{
		Name:        urlTypeName,
		Description: "Structured URL",
		Methods: []*luadoc.Method{
			{
				Name:        "hostname",
				Description: "Return hostname without port numbers.",
				Value:       urlURLHostname,
				Returns: []*luadoc.Param{
					{
						Name:        "hostname",
						Description: "URLs hostname",
						Type:        luadoc.String,
					},
				},
			},
			{
				Name:        "parse",
				Description: "Parses a URL in the context of the receiver. The provided URL may be relative or absolute.",
				Value:       urlURLParse,
				Params: []*luadoc.Param{
					{
						Name:        "ref",
						Description: "",
						Type:        luadoc.String,
					},
				},
				Returns: []*luadoc.Param{
					{
						Name:        "url",
						Description: "",
						Type:        urlTypeName,
					},
				},
			},
			{
				Name:        "string",
				Description: "Reassembles the URL into a valid URL string.",
				Value:       urlURLString,
				Returns: []*luadoc.Param{
					{
						Name:        "url",
						Description: "",
						Type:        luadoc.String,
					},
				},
			},
			{
				Name:        "join_path",
				Description: "Returns a new URL with the provided path elements joined to any existing path and the resulting path cleaned of any ./ or ../ elements. Any sequences of multiple / characters will be reduced to a single /.",
				Value:       urlURLJoinPath,
				Params: []*luadoc.Param{
					{
						Name:        "...",
						Description: "",
						Type:        luadoc.String,
					},
				},
				Returns: []*luadoc.Param{
					{
						Name:        "url",
						Description: "",
						Type:        urlTypeName,
					},
				},
			},
			{
				Name:        "query",
				Description: "",
				Value:       urlURLQuery,
				Params: []*luadoc.Param{
					{
						Name:        "query",
						Description: "",
						Type:        valuesTypeName,
						Optional:    true,
					},
				},
				Returns: []*luadoc.Param{
					{
						Name:        "query",
						Description: "",
						Type:        valuesTypeName,
						Optional:    true,
					},
				},
			},
			{
				Name:        "copy",
				Description: "Copy URL",
				Value:       urlURLCopy,
				Returns: []*luadoc.Param{
					{
						Name:        "url",
						Description: "A copied URL",
						Type:        urlTypeName,
					},
				},
			},
		},
	}

	return &luadoc.Lib{
		Name:        libName,
		Description: "URLs is a library for working with URLs.",
		Funcs: []*luadoc.Func{
			{
				Name:        "parse",
				Description: "Parses URL",
				Value:       urlParse,
				Params: []*luadoc.Param{
					{
						Name:        "raw_url",
						Description: "URL string to parse",
						Type:        luadoc.String,
					},
				},
				Returns: []*luadoc.Param{
					{
						Name:        "url",
						Description: "Parsed URL",
						Type:        urlTypeName,
					},
				},
			},
			{
				Name:        "values",
				Description: "Creates a new values.",
				Value:       urlValues,
				Returns: []*luadoc.Param{
					{
						Name:        "values",
						Description: "The new values.",
						Type:        valuesTypeName,
					},
				},
			},
			{
				Name:        "path_escape",
				Description: `Escapes the string so it can be safely placed inside a URL path segment, replacing special characters (including /) with %XX sequences as needed.`,
				Value:       urlPathEscape,
				Params: []*luadoc.Param{
					{
						Name:        "path",
						Description: "The path to escape.",
						Type:        luadoc.String,
					},
				},
				Returns: []*luadoc.Param{
					{
						Name:        "escaped",
						Description: "The escaped path.",
						Type:        luadoc.String,
					},
				},
			},
			{
				Name:        "path_unescape",
				Description: `Unescapes a string; the inverse operation of path_escape. It converts each 3-byte encoded substring of the form "%AB" into the hex-decoded byte 0xAB.`,
				Value:       urlPathUnescape,
				Params: []*luadoc.Param{
					{
						Name:        "escaped",
						Description: "The escaped path.",
						Type:        luadoc.String,
					},
				},
				Returns: []*luadoc.Param{
					{
						Name:        "path",
						Description: "The unescaped path.",
						Type:        luadoc.String,
					},
				},
			},
			{
				Name:        "query_escape",
				Description: `Escapes the string so it can be safely placed inside a URL query parameter, replacing special characters (including /) with %XX sequences as needed.`,
				Value:       urlQueryEscape,
				Params: []*luadoc.Param{
					{
						Name:        "query",
						Description: "The query to escape.",
						Type:        luadoc.String,
					},
				},
				Returns: []*luadoc.Param{
					{
						Name:        "escaped",
						Description: "The escaped query.",
						Type:        luadoc.String,
					},
				},
			},
			{
				Name:        "query_unescape",
				Description: `Unescapes a string; the inverse operation of query_escape. It converts each 3-byte encoded substring of the form "%AB" into the hex-decoded byte 0xAB.`,
				Value:       urlQueryUnescape,
				Params: []*luadoc.Param{
					{
						Name:        "escaped",
						Description: "The escaped query.",
						Type:        luadoc.String,
					},
				},
				Returns: []*luadoc.Param{
					{
						Name:        "query",
						Description: "The unescaped query.",
						Type:        luadoc.String,
					},
				},
			},
		},
		Classes: []*luadoc.Class{
			classValues,
			classURL,
		},
	}
}

func urlPathEscape(L *lua.LState) int {
	s := L.CheckString(1)
	L.Push(lua.LString(url.PathEscape(s)))
	return 1
}

func urlPathUnescape(L *lua.LState) int {
	s := L.CheckString(1)
	s, err := url.PathUnescape(s)
	if err != nil {
		L.RaiseError(err.Error())
		return 0
	}
	L.Push(lua.LString(s))
	return 1
}

func urlQueryEscape(L *lua.LState) int {
	s := L.CheckString(1)
	L.Push(lua.LString(url.QueryEscape(s)))
	return 1
}

func urlQueryUnescape(L *lua.LState) int {
	s := L.CheckString(1)
	s, err := url.QueryUnescape(s)
	if err != nil {
		L.RaiseError(err.Error())
		return 0
	}
	L.Push(lua.LString(s))
	return 1
}

func urlValues(L *lua.LState) int {
	util.Push(L, &url.Values{}, valuesTypeName)
	return 1
}

func urlValuesAdd(L *lua.LState) int {
	v := util.Check[*url.Values](L, 1)
	key := L.CheckString(2)
	value := L.CheckString(3)
	v.Add(key, value)
	return 0
}

func urlValuesSet(L *lua.LState) int {
	v := util.Check[*url.Values](L, 1)
	key := L.CheckString(2)
	value := L.CheckString(3)
	v.Set(key, value)
	return 0
}

func urlValuesGet(L *lua.LState) int {
	v := util.Check[*url.Values](L, 1)
	key := L.CheckString(2)
	L.Push(lua.LString(v.Get(key)))
	return 1
}

func urlValuesString(L *lua.LState) int {
	v := util.Check[*url.Values](L, 1)
	L.Push(lua.LString(v.Encode()))
	return 1
}

func urlValuesParse(L *lua.LState) int {
	s := L.CheckString(1)
	v, err := url.ParseQuery(s)
	if err != nil {
		L.RaiseError(err.Error())
		return 0
	}
	util.Push(L, &v, valuesTypeName)
	return 1
}

func urlValuesDel(L *lua.LState) int {
	v := util.Check[*url.Values](L, 1)
	key := L.CheckString(2)
	v.Del(key)
	return 0
}

func urlValuesHas(L *lua.LState) int {
	v := util.Check[*url.Values](L, 1)
	key := L.CheckString(2)
	L.Push(lua.LBool(v.Has(key)))
	return 1
}

func urlParse(L *lua.LState) int {
	rawURL := L.CheckString(1)

	u, err := url.Parse(rawURL)
	if err != nil {
		L.RaiseError(err.Error())
		return 0
	}

	util.Push(L, u, urlTypeName)
	return 1
}

func urlURLHostname(L *lua.LState) int {
	u := util.Check[*url.URL](L, 1)

	L.Push(lua.LString(u.Hostname()))
	return 1
}

func urlURLQuery(L *lua.LState) int {
	u := util.Check[*url.URL](L, 1)

	if L.GetTop() == 1 {
		values := u.Query()
		util.Push(L, &values, valuesTypeName)
		return 1
	} else { // >= 2
		values := util.Check[*url.Values](L, 2)
		u.RawQuery = values.Encode()
		return 0
	}
}

func urlURLJoinPath(L *lua.LState) int {
	u := util.Check[*url.URL](L, 1)

	top := L.GetTop()
	elems := make([]string, top-1)

	for i := 2; i <= top; i++ {
		elems[i-2] = L.CheckString(i)
	}

	util.Push(L, u.JoinPath(elems...), urlTypeName)
	return 1
}

func urlURLParse(L *lua.LState) int {
	u := util.Check[*url.URL](L, 1)
	ref := L.CheckString(2)

	parsed, err := u.Parse(ref)
	if err != nil {
		L.RaiseError(err.Error())
		return 0
	}

	util.Push(L, parsed, urlTypeName)
	return 1
}

func urlURLString(L *lua.LState) int {
	u := util.Check[*url.URL](L, 1)

	L.Push(lua.LString(u.String()))
	return 1
}

func urlURLCopy(L *lua.LState) int {
	u := util.Check[*url.URL](L, 1)

	copied := *u
	util.Push(L, &copied, urlTypeName)
	return 1
}
