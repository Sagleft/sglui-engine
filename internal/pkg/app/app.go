package app

import (
	"tool/internal/pkg/engine"
)

type App interface {
	GetPages() []engine.Page
	GetDefaultPageTag() string

	SetEventHandler(engine.EventHandler)
}

func (a *def) GetPages() []engine.Page {
	// TODO
	return nil
}

func (a *def) GetDefaultPageTag() string {
	return "" // TODO
}

func (a *def) SetEventHandler(e engine.EventHandler) {
	a.e = e
}
