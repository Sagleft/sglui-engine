package main

import (
	"context"
	"embed"
	"fmt"
	"tool/internal/pkg/app"
	"tool/internal/pkg/consts"

	swissknife "github.com/Sagleft/swiss-knife"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:engine/dist
var assets embed.FS

type Engine interface {
	CreateApp() (app.App, error)
}

type defaultEngine struct {
	ctx context.Context
}

func NewEngine() (Engine, error) {
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

	if err := e.runApp(); err != nil {
		return nil, fmt.Errorf("run app: %w", err)
	}
	return a, nil
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (e *defaultEngine) startup(ctx context.Context) {
	e.ctx = ctx
}

func (e *defaultEngine) runApp() error {
	if err := wails.Run(&options.App{
		Title:  "App name",
		Width:  1280,
		Height: 720,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: e.startup,
		Bind:      []interface{}{e},
	}); err != nil {
		return fmt.Errorf("run wails app: %w", err)
	}
	return nil
}
