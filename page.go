package luaprovider

import (
	"github.com/mangalorg/libmangal"
	"regexp"
)

var fileExtensionRegex = regexp.MustCompile(`^\.[a-zA-Z0-9][a-zA-Z0-9.]*[a-zA-Z0-9]$`)

type Page struct {
	Extension string `gluamapper:"extension"`

	// URL is the url of the page image
	URL string `gluamapper:"url"`

	Headers map[string]string `gluamapper:"headers"`
	Cookies map[string]string `gluamapper:"cookies"`

	chapter *Chapter
}

func (p Page) String() string {
	return p.URL
}

func (p Page) Chapter() libmangal.Chapter {
	return p.chapter
}

func (p Page) GetExtension() string {
	return p.Extension
}
