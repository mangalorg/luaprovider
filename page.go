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

	// Data is the raw data of the page image.
	// It will have a higher priority than URL if it is not empty.
	// string is used instead of []byte because lua cannot handle []byte.
	Data string `gluamapper:"data"`

	Headers map[string]string `gluamapper:"headers"`
	Cookies map[string]string `gluamapper:"cookies"`

	chapter *Chapter
}

func (p *Page) Chapter() libmangal.Chapter {
	return p.chapter
}

func (p *Page) fillDefaults() {
}

func (p *Page) GetExtension() string {
	return p.Extension
}
