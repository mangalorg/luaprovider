package luaprovider

import (
	"github.com/mangalorg/libmangal"
	"regexp"
)

var fileExtensionRegex = regexp.MustCompile(`^\.[a-zA-Z0-9][a-zA-Z0-9.]*[a-zA-Z0-9]$`)

var _ libmangal.Page = (*luaPage)(nil)

type luaPage struct {
	Extension string `gluamapper:"extension"`

	// URL is the url of the page image
	URL string `gluamapper:"url"`

	Headers map[string]string `gluamapper:"headers"`
	Cookies map[string]string `gluamapper:"cookies"`

	chapter *luaChapter
}

func (p luaPage) String() string {
	return p.URL
}

func (p luaPage) Chapter() libmangal.Chapter {
	return p.chapter
}

func (p luaPage) GetExtension() string {
	return p.Extension
}
