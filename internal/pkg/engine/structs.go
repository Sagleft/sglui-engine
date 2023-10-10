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

type Page struct {
	e    EventHandler
	data PageData

	Elements []any
	OnEnter  func()
	OnLeave  func()
}

type ButtonData struct {
	Label string `json:"label"`
	Tag   string `json:"tag"`
}

type Button struct {
	e         EventHandler
	data      ButtonData
	OnPressed func()
}

type Sidebar struct {
	Title             string    `json:"title"`
	Tag               string    `json:"tag"`
	Buttons           []*Button `json:"buttons"`
	LogoImageFilename string    `json:"logo"`
}
