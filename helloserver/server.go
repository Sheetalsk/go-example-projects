package main

import (
	"fmt"
	"time"
)

func main() {
	// This is a simple program to demonstrate static checks
	var unusedVar string // Unused variable, staticcheck will catch this

	fmt.Println("Hello, Static Check!")
	time.Sleep(1 * time.Second)
	fmt.Println("Goodbye, Static Check!")

	if len(unusedVar) == 0 {
		return nil
	 }
	 return nil // Can be simplified

	 t := time.Now()
	fmt.Println(t.Sub(t)) 
}
