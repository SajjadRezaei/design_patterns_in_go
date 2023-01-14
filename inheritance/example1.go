package inheritance

import "fmt"

type base struct {
	color string
}

func (b *base) say() {
	fmt.Println("Hi from say function")
}

type child struct {
	base
	style string
}

func Run() {
	b := base{color: "Red"}
	ch := &child{
		base:  b,
		style: "somestyle",
	}
	ch.say()
	//ch.color = "green"
	fmt.Println("the color is " + ch.color)
}
