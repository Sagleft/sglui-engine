package main

import (
	"fmt"
	"tool/internal/pkg/app"
)

// WailsHandler ..
type WailsHandler struct {
	engine Engine
	app    app.App
}

func main() {
	// Create an instance of the h structure
	h, err := NewHandler()
	if err != nil {
		println("create app:", err.Error())
	}

	// Create application with options
	h.app, err = h.engine.CreateApp()
	if err != nil {
		// TBD: test & update error handler
		// https://github.com/Sagleft/sglui-engine/issues/1
		println("run app:", err.Error())
	}
}

// NewHandler creates a new Wals app struct
func NewHandler() (*WailsHandler, error) {
	e, err := NewEngine()
	if err != nil {
		return nil, fmt.Errorf("init engine: %w", err)
	}

	return &WailsHandler{engine: e}, nil
}
