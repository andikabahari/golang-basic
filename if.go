package main

import "fmt"

func main() {
	name := "John Doe"
	if name == "Ben" {
		fmt.Println(name)
	} else {
		fmt.Println("You are not " + name)
	}

	if length := len(name); length < 20 {
		fmt.Println("OK")
	} else {
		fmt.Println("Too long")
	}
}