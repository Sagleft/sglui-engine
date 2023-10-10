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
	return &Page{e: e, data: data}
}

func (p *Page) SetElements(el []any) *Page {
	p.Elements = el
	return p
}

func NewButton(data ButtonData, e EventHandler) *Button {
	return &Button{e: e, data: data}
}

func NewSidebar(btns []*Button) *Sidebar {
	return &Sidebar{Buttons: btns}
}
