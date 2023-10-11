package app

import (
	"tool/internal/pkg/engine"
)

type App interface {
	GetPages() engine.Pages
	GetDefaultPageTag() string

	SetEventHandler(engine.EventHandler)
}

func (a *application) SetEventHandler(e engine.EventHandler) {
	a.e = e
}
