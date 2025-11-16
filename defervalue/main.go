package main

import "fmt"

func main() {
	// Go using call by value so a copy of the variable being passed in is what is used at runtime
	// meaning that if we use a value receiver variable in a defer function and change later, defer won't see
	// that new value unless the variable is a pointer receiver
	valueVar := 5
	pointerVar := &valueVar

	defer func(valueNum int, pointerNum *int) {
		fmt.Printf("Value Var: %v\n", valueNum)
		fmt.Printf("Pointer Var: %v\n", *pointerNum)
	}(valueVar, pointerVar)

	valueVar = 10
}
