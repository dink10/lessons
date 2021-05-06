package main

import (
	"fmt"
)

func main() {
	fmt.Println(extra(5, "Hello"))
}

func extra(index int, str string) (sum int, st string) {

	sum = index + 10
	st = str + " World"

	if sum == 15 {
		return 105, ""
	}

	return sum, st
}
