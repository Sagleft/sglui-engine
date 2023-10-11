package engine

// Engine should implement this
type EventHandler interface {
	// control
	ChangePage(pageTag string)
	CloseApp()

	// notify
	ShowNotify(string)
	ShowError(string)
}

func NewPage(data PageData, e EventHandler) *Page {
	return &Page{e: e, Data: data}
}

func (p *Page) SetElements(es Elements) *Page {
	p.Elements = es
	return p
}

func NewButton(data ButtonData, e EventHandler) *Button {
	return &Button{e: e, Data: data}
}

func NewSidebar(btns []*Button) *Sidebar {
	return &Sidebar{Buttons: btns}
}
