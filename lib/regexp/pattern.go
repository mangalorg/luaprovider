package regexp

import (
	"github.com/mangalorg/luaprovider/util"
	lua "github.com/yuin/gopher-lua"
	"regexp"
)

const patternTypeName = libName + "_pattern"

func regexpFindSubmatch(L *lua.LState) int {
	re := util.Check[*regexp.Regexp](L, 1)
	text := L.CheckString(2)
	matches := re.FindStringSubmatch(text)
	tbl := L.NewTable()
	for _, match := range matches {
		tbl.Append(lua.LString(match))
	}
	L.Push(tbl)
	return 1
}

func regexpMatch(L *lua.LState) int {
	re := util.Check[*regexp.Regexp](L, 1)
	text := L.CheckString(2)
	matched := re.MatchString(text)
	L.Push(lua.LBool(matched))
	return 1
}

func regexpReplaceAll(L *lua.LState) int {
	re := util.Check[*regexp.Regexp](L, 1)
	text := L.CheckString(2)
	replacement := L.CheckString(3)
	result := re.ReplaceAllString(text, replacement)
	L.Push(lua.LString(result))
	return 1
}

func regexpSplit(L *lua.LState) int {
	re := util.Check[*regexp.Regexp](L, 1)
	text := L.CheckString(2)
	tbl := L.NewTable()
	for _, match := range re.Split(text, -1) {
		tbl.Append(lua.LString(match))
	}
	L.Push(tbl)
	return 1
}

func regexpGroups(L *lua.LState) int {
	re := util.Check[*regexp.Regexp](L, 1)
	value := L.CheckString(2)

	// match all groups as map
	matches := re.FindStringSubmatch(value)
	if matches == nil {
		L.Push(lua.LNil)
		return 1
	}

	tbl := L.NewTable()
	for i, name := range re.SubexpNames() {
		if i == 0 {
			continue
		}
		tbl.RawSetString(name, lua.LString(matches[i]))
	}

	L.Push(tbl)
	return 1
}

func regexpReplaceAllFunc(L *lua.LState) int {
	re := util.Check[*regexp.Regexp](L, 1)
	text := L.CheckString(2)
	replacer := L.CheckFunction(3)

	result := re.ReplaceAllStringFunc(text, func(match string) string {
		L.Push(replacer)
		L.Push(lua.LString(match))
		L.Call(1, 1)
		return L.CheckString(-1)
	})

	L.Push(lua.LString(result))
	return 1
}
