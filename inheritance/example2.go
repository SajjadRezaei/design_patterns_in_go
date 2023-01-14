package inheritance

import "fmt"

type base2 struct {
	color string
}

func (b *base2) say2() {
	fmt.Println("Hi from say function")
}

type child2 struct {
	base2
	style string
}

func check(b base) {
	b.say()
}

func Run2() {
	b2 := base2{color: "Red"}
	ch := &child2{
		base2: b2,
		style: "somestyle",
	}
	ch.say2()
	//ch.color = "green"
	fmt.Println("the color is " + ch.color)

	//check(child) //important: cannot use child (type *child) as type base in argument to check
}
