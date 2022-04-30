package main

import "fmt"

func main() {
	person := map[string]string {
		"firstName": "Jazz",
		"lastName": "Fusion",
	}

	fmt.Println(person)
	fmt.Println(person["firstName"])
	fmt.Println(person["lastName"])
}