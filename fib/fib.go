package main

import "fmt"

func factIter(n int) int {
	res := 1
	for i := 2; i <= n; i++ {
		res *= i
	}
	return res
}

func main() {
	n := 5
	fmt.Println(factIter(n))
}
