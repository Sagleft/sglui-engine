package app

type App interface{}

type def struct {
	cfg AppConfig
}

type AppConfig struct {
	Window WindowConfig `json:"window"`
}

type WindowConfig struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func New(config AppConfig) (App, error) {
	return &def{cfg: config}, nil
}
