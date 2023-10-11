package consts

type Constants struct{}

func NewProvider() *Constants {
	return &Constants{}
}

func (Constants) TypeButton() string {
	return TypeButton
}

func (Constants) TypeSidebar() string {
	return TypeSidebar
}
