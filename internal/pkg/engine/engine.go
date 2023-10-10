package engine

import (
	"fmt"
	"tool/internal/pkg/app"
	"tool/internal/pkg/consts"

	swissknife "github.com/Sagleft/swiss-knife"
)

type Engine interface {
	CreateApp() (app.App, error)
}

type defaultEngine struct{}

func New() (Engine, error) {
	return &defaultEngine{}, nil
}

func (e *defaultEngine) CreateApp() (app.App, error) {
	var cfg app.AppConfig

	if err := swissknife.ParseStructFromJSONFile(
		consts.AppConfigFilePath,
		&cfg,
	); err != nil {
		return nil, fmt.Errorf("read config file: %w", err)
	}

	a, err := app.New(cfg)
	if err != nil {
		return nil, fmt.Errorf("create app: %w", err)
	}
	return a, nil
}
