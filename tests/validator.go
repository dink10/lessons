package tests

import "fmt"

// ValidateString - len >= 3, isString
func ValidateString(input interface{}) error {
	v, ok := input.(string)
	if !ok {
		return fmt.Errorf("input param is not a string")
	}

	if len(v) < 3 {
		return fmt.Errorf("lenght less than 3")
	}

	return nil
}

func sum(a, b int) int {
	return a + b
}
