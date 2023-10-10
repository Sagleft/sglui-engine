package main

import (
	"context"
	"embed"
	"fmt"
	"tool/internal/pkg/app"
	"tool/internal/pkg/engine"

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
	engineCfg, err := engine.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("load config: %w", err)
	}

	a, err := app.New(engineCfg.App)
	if err != nil {
		return nil, fmt.Errorf("create app: %w", err)
	}

	if err := e.runEnginge(engineCfg); err != nil {
		return nil, fmt.Errorf("run app: %w", err)
	}
	return a, nil
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (e *defaultEngine) startup(ctx context.Context) {
	e.ctx = ctx
}

func (e *defaultEngine) runEnginge(engineCfg engine.Config) error {
	if err := wails.Run(&options.App{
		Title:  engineCfg.Window.Title,
		Width:  engineCfg.Window.Width,
		Height: engineCfg.Window.Height,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup:                e.startup,
		Bind:                     []interface{}{e},
		EnableDefaultContextMenu: false,
	}); err != nil {
		return fmt.Errorf("run wails app: %w", err)
	}
	return nil
}
