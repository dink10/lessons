package animals

import "fmt"

type Animal struct {
	name   string
	height int
	weight int
}

func (a Animal) SaySomething() {
	fmt.Printf("%s said something\n", a.name)
}
