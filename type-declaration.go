package main

import "fmt"

func main() {
	type Id string
	type Married bool

	var personId Id = "69696969"
	var isMarried Married = false
	fmt.Println(personId)
	fmt.Println(isMarried)
}