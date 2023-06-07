package luaprovider

import (
	"errors"
	"fmt"
	"regexp"
)

var fileExtensionRegex = regexp.MustCompile(`^\.[a-zA-Z0-9][a-zA-Z0-9.]*[a-zA-Z0-9]$`)

type Page struct {
	Extension string

	// URL is the url of the page image
	URL string

	// Data is the raw data of the page image.
	// It will have a higher priority than URL if it is not empty.
	// string is used instead of []byte because lua cannot handle []byte.
	Data string

	Headers, Cookies map[string]string

	chapter *Chapter
}

func (p *Page) Validate() error {
	if p.URL == "" && p.Data == "" {
		return errors.New("either URL or Data must be set")
	}

	if !fileExtensionRegex.MatchString(p.Extension) {
		return fmt.Errorf("invalid page extension: %s", p.Extension)
	}

	return nil
}

func (p *Page) fillDefaults() {
	if p.Extension == "" {
		p.Extension = ".jpg"
	}

	if p.Headers == nil {
		p.Headers = make(map[string]string)
		p.Headers["Referer"] = p.chapter.URL
		p.Headers["Accept"] = "image/webp,image/apng,image/*,*/*;q=0.8"

		// TODO: generate random user-agent
		p.Headers["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36"
	}
}

func (p *Page) GetExtension() string {
	return p.Extension
}
