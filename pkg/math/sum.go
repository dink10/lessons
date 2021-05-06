package math

import "fmt"

type Base struct {
	Reader
	Writer
}

type Reader struct {
	Name string
}

type Writer struct {
	Name string
}

func Sum(x, y int) int {

	log(func() { fmt.Println(x + y) })

	return x + y
}

func log(f func()) {
	f()
}
