package app

import "tool/internal/pkg/engine"

func (a *application) createButton(data engine.ButtonData) *engine.Button {
	return engine.NewButton(data, a.e)
}
