package main

import "fmt"

func main() {
	n := 1
	switch n {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Unknown")
	}

	name := "Example"
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Too many")
	case false:
		fmt.Println("OK")
	}
}