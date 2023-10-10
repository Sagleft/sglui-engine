package app

import "tool/internal/pkg/engine"

type def struct {
	e   engine.EventHandler
	cfg AppConfig
}

type AppConfig struct{}
