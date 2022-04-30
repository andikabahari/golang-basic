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

	song := make(map[string]string)
	song["title"] = "Giant Steps"
	song["composer"] = "John Coltrane"
	song["any"] = "..."
	fmt.Println(song)

	delete(song, "any")
	fmt.Println(song)
}