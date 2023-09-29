package logic

import "fmt"

type App interface {
	Do()
}

func Test() {
	fmt.Println("original engine func")
}
