package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) updateAge2(newAge int) {
	p.age = newAge
}

func (p *person) updateAge(newAge int) {
	p.age = newAge
}

func main() {
	var tom = &person{name: "Tom", age: 24}
	fmt.Println("before", tom.age)
	tom.updateAge(33)
	tom.updateAge(35)
	tom.updateAge(40)
	fmt.Println("after", tom.age)
}
