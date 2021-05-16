package util

import "fmt"

var Variable string

func init() {
	Variable = "Hello %s!"

	fmt.Println(Variable)
}
