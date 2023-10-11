package engine

type Config struct {
	App    any          `json:"app"`
	Window WindowConfig `json:"window"`
}

type WindowConfig struct {
	Title  string `json:"title"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type PageData struct {
	Title string `json:"title"`
	Tag   string `json:"tag"`
}

// Button, Sidebar, List, etc
type Element struct {
	Type string `json:"type"`
	Data any    `json:"data"`
}

type Elements []Element

type Page struct {
	e        EventHandler `json:"-"`
	Data     PageData     `json:"data"`
	Elements []Element    `json:"elements"`

	OnEnter func() `json:"-"`
	OnLeave func() `json:"-"`
}

type Pages []*Page

type ButtonData struct {
	Label string `json:"label"`
	Tag   string `json:"tag"`
}

type Button struct {
	e         EventHandler `json:"-"`
	Data      ButtonData   `json:"data"`
	OnPressed func()       `json:"-"`
}

type Buttons []*Button

type Sidebar struct {
	Title             string  `json:"title"`
	Tag               string  `json:"tag"`
	Buttons           Buttons `json:"buttons"`
	LogoImageFilename string  `json:"logo"`
}
