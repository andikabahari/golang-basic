package main

import "fmt"

func main() {
	const value = 69
	fmt.Println(value)

	const (
		firstName = "Jazz"
		lastName string = "Fusion"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}