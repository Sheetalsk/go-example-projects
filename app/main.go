package main

import "fmt"

func main() {
	// Unused variable
	var unusedVar string

	// Redundant operation (len of a string literal)
	str := "hello"
	fmt.Println(len(str)) // This is fine
	fmt.Println(len("hello")) // This is redundant, staticcheck will catch this

	// Inefficient slicing (slicing the whole array)
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arr[:]) // staticcheck can detect this

	// An unused function
	unusedFunction()

	// Print a message
	fmt.Println("This is a test for staticcheck.")
}

// Unused function
func unusedFunction() {
	fmt.Println("This function is unused.")
}
