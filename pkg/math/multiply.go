package math

import "fmt"

func Multiply(x, y int) int {
	log(func() { fmt.Println(x * y) })

	return x * y
}
