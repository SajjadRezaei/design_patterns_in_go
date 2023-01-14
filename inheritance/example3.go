package inheritance

import "fmt"

type iBase interface {
	say()
}

type base3 struct {
	color string
}

func (b *base3) say() {
	fmt.Println("Hi from say function")
}

type child3 struct {
	base3
	style string
}

func check2(b iBase) {
	b.say()
}

func Run3() {
	base := base3{
		color: "BV",
	}
	ch := &child3{
		base3: base,
		style: "asdasd",
	}
	ch.say()
	fmt.Println("the color is " + ch.color)
	check2(ch)
}
