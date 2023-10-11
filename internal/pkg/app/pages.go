package app

import (
	"tool/internal/pkg/consts"
	"tool/internal/pkg/engine"
)

const mainPageTag = "main"

func (a *application) GetPages() engine.Pages {
	return engine.Pages{
		a.getMainPage(),
	}
}

func (a *application) getMainPage() *engine.Page {
	btn1 := engine.NewButton(engine.ButtonData{
		Label: "Test btn 1",
		Tag:   "btn1",
	}, a.e)
	btn2 := engine.NewButton(engine.ButtonData{
		Label: "Test btn 1",
		Tag:   "btn2",
	}, a.e)

	sidebarBtns := engine.Buttons{btn1, btn2}
	sidebar := engine.NewSidebar(sidebarBtns)
	sidebar.Title = "My sidebar"
	sidebar.Tag = "sidebar1"

	return engine.NewPage(engine.PageData{
		Title: "Main Page",
		Tag:   mainPageTag,
	}, a.e).SetElements(engine.Elements{
		engine.Element{
			Type: consts.TypeSidebar,
			Data: sidebar,
		},
	})
}

func (a *application) GetDefaultPageTag() string {
	return mainPageTag
}
