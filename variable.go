package main

import "fmt"

func main() {
	var name string
	
	name = "Jazz"
	fmt.Println(name)

	name = "Fusion"
	fmt.Println(name)

	var fullName = "Jazz Fusion"
	fmt.Println(fullName)

	var age = 69
	fmt.Println(age)

	city := "Bandung"
	fmt.Println(city)

	var (
		firstName = "Jazz"
		lastName = "Fusion"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}