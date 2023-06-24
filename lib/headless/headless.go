package headless

import (
	luadoc "github.com/mangalorg/luaprovider/doc"
)

const libName = "headless"

func Lib() *luadoc.Lib {
	classBrowser := &luadoc.Class{
		Name:        browserTypeName,
		Description: "",
		Methods: []*luadoc.Method{
			{
				Name:        "page",
				Description: "Visit the page",
				Value:       browserPage,
				Params: []*luadoc.Param{
					{
						Name:        "url",
						Description: "url of the page to visit",
						Type:        luadoc.String,
					},
				},
				Returns: []*luadoc.Param{
					{
						Name:        "page",
						Description: "",
						Type:        pageTypeName,
					},
				},
			},
		},
	}

	classPage := &luadoc.Class{
		Name:        pageTypeName,
		Description: "",
		Methods: []*luadoc.Method{
			{
				Name:        "html",
				Description: "",
				Value:       pageHTML,
				Returns: []*luadoc.Param{
					{
						Name:        "html",
						Description: "",
						Type:        luadoc.String,
					},
				},
			},
			{
				Name:        "navigate",
				Description: "",
				Value:       pageNavigate,
				Params: []*luadoc.Param{
					{
						Name:        "url",
						Description: "URL to navigate",
						Type:        luadoc.String,
					},
				},
			},
			{
				Name:        "element",
				Description: "Get HTML element by CSS selector",
				Value:       pageElement,
				Params: []*luadoc.Param{
					{
						Name:        "selector",
						Description: "CSS selector for the element",
						Type:        luadoc.String,
					},
				},
				Returns: []*luadoc.Param{
					{
						Name:        "element",
						Description: "HTML element",
						Type:        elementTypeName,
					},
				},
			},
		},
	}

	classElement := &luadoc.Class{
		Name:        elementTypeName,
		Description: "",
		Methods: []*luadoc.Method{
			{
				Name:        "html",
				Description: "",
				Value:       elementHTML,
				Returns: []*luadoc.Param{
					{
						Name:        "html",
						Description: "",
						Type:        luadoc.String,
					},
				},
			},
			{
				Name:        "click",
				Description: "",
				Value:       elementClick,
			},
			{
				Name:        "input",
				Description: "",
				Value:       elementInput,
				Params: []*luadoc.Param{
					{
						Name:        "text",
						Description: "",
						Type:        luadoc.String,
					},
				},
			},
		},
	}

	return &luadoc.Lib{
		Name:        libName,
		Description: "Headless browser",
		Funcs: []*luadoc.Func{
			{
				Name:        "browser",
				Description: "Creates a new headless browser",
				Value:       newBrowser,
				Returns: []*luadoc.Param{
					{
						Name:        "browser",
						Description: "headless browser instance",
						Type:        browserTypeName,
					},
				},
			},
		},
		Classes: []*luadoc.Class{
			classBrowser,
			classPage,
			classElement,
		},
	}
}
