package app

import (
	"fmt"
	"tool/internal/pkg/engine"
)

func New(cfgRaw any) (App, error) {
	var cfg AppConfig
	if err := engine.ReconvertStruct(cfgRaw, &cfg); err != nil {
		return nil, fmt.Errorf("convert app config: %w", err)
	}

	return &application{cfg: cfg}, nil
}
