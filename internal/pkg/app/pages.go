package app

import "tool/internal/pkg/engine"

const mainPageTag = "main"

func (a *def) GetPages() engine.Pages {
	return engine.Pages{
		a.getMainPage(),
	}
}

func (a *def) getMainPage() *engine.Page {
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

	return engine.NewPage(engine.PageData{
		Title: "Main Page",
		Tag:   mainPageTag,
	}, a.e).SetElements([]interface{}{
		sidebar,
	})
}

func (a *def) GetDefaultPageTag() string {
	return mainPageTag
}
