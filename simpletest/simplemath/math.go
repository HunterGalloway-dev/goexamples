// Package simplemath simple math functions and handles them
package simplemath

import "fmt"

func Add(x, y int) int {
	return x + y
}

func Subtract(x, y int) int {
	return x-y
}

func Multiply(x, y int) int { 
	return x*y
}

func Divide(x, y int) (int, int, error) {
	if y == 0 {
		return 0, 0, fmt.Errorf("cannot divide by zero")
	}

	return x / y, x % y, nil
}
