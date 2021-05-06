package main

import "fmt"

type Cat struct{}

func (Cat) voice() {
	fmt.Println("Meow")
}

func (Cat) run() int {
	fmt.Println("Meow")

	return 10
}

func (Cat) eat() {
	fmt.Println("Meow")
}
