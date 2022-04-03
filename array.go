package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Awesome"
	names[1] = "Jazz"
	names[2] = "Fusion"
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		69,
		420,
		599,
	}
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(names))
	fmt.Println(len(values))

	var empty [10]string
	fmt.Println(len(empty))
}