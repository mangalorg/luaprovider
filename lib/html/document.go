package html

import (
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/cixtor/readability"
	lua "github.com/yuin/gopher-lua"
	"strings"
)

const documentTypeName = libName + "_document"

func checkDocument(L *lua.LState, n int) *goquery.Document {
	ud := L.CheckUserData(n)
	if v, ok := ud.Value.(*goquery.Document); ok {
		return v
	}
	L.ArgError(1, "document expected")
	return nil
}

func pushDocument(L *lua.LState, document *goquery.Document) {
	ud := L.NewUserData()
	ud.Value = document
	L.SetMetatable(ud, L.GetTypeMetatable(documentTypeName))
	L.Push(ud)
}

func parse(L *lua.LState) int {
	value := L.CheckString(1)
	reader := strings.NewReader(value)
	document, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		L.RaiseError(err.Error())
		return 0
	}

	pushDocument(L, document)
	return 1
}

func documentFind(L *lua.LState) int {
	document := checkDocument(L, 1)
	selector := L.CheckString(2)

	selection := document.Find(selector)
	pushSelection(L, selection)
	return 1
}

func documentHtml(L *lua.LState) int {
	document := checkDocument(L, 1)
	html, err := document.Html()
	if err != nil {
		L.RaiseError(err.Error())
		return 0
	}

	L.Push(lua.LString(html))
	return 1
}

func documentSelection(L *lua.LState) int {
	document := checkDocument(L, 1)
	selection := document.Selection
	pushSelection(L, selection)
	return 1
}

func documentMarkdown(L *lua.LState) int {
	document := checkDocument(L, 1)
	converter := md.NewConverter("", true, nil)

	L.Push(lua.LString(converter.Convert(document.Selection)))
	return 1
}

func documentSimplified(L *lua.LState) int {
	document := checkDocument(L, 1)
	html, err := document.Html()
	if err != nil {
		L.RaiseError(err.Error())
		return 0
	}

	article, err := readability.New().Parse(strings.NewReader(html), "https://example.com")
	if err != nil {
		L.RaiseError(err.Error())
		return 0
	}

	document = goquery.NewDocumentFromNode(article.Node)

	pushDocument(L, document)
	return 1
}
