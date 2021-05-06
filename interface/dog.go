package main

import "fmt"

type Dog struct{}

func (Dog) voice() {
	fmt.Println("Woof")
}

func (Dog) run() int {
	fmt.Println("Meow")

	return 10
}

func (Dog) eat() {
	fmt.Println("Meow")
}
