package app

import (
	"tool/internal/pkg/engine"
)

type application struct {
	e   engine.EventHandler
	cfg AppConfig
}

// AppConfig - settings that apply exclusively to this package
type AppConfig struct{}
