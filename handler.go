package main

import (
	"context"
	"embed"
	"fmt"
	"tool/internal/pkg/engine"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

// WailsHandler ..
type WailsHandler struct {
	ctx    context.Context
	engine engine.Engine
}

//go:embed all:engine/dist
var assets embed.FS

func main() {
	// Create an instance of the h structure
	h, err := NewHandler()
	if err != nil {
		println("create app:", err.Error())
	}

	// Create application with options
	if err := h.run(); err != nil {
		// TBD: test & update error handler
		// https://github.com/Sagleft/sglui-engine/issues/1
		println("run app:", err.Error())
	}
}

// NewHandler creates a new Wals app struct
func NewHandler() (*WailsHandler, error) {
	e, err := engine.New()
	if err != nil {
		return nil, fmt.Errorf("init engine: %w", err)
	}

	return &WailsHandler{engine: e}, nil
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *WailsHandler) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *WailsHandler) run() error {
	if err := wails.Run(&options.App{
		Title:  "App name",
		Width:  1280,
		Height: 720,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: a.startup,
		Bind:      []interface{}{a},
	}); err != nil {
		return fmt.Errorf("run wails app: %w", err)
	}
	return nil
}

// Greet returns a greeting for the given name
func (a *WailsHandler) Greet(name string) string {
	return fmt.Sprintf(
		"Result: %s",
		"hello, world",
		//logic.Test(),
	)
}
