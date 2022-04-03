package main

import "fmt"

func main() {
	var name1 = "Jazz"
	var name2 = "Fusion"
	var result = name1 == name2
	fmt.Println(result)

	var value1 = 69
	var value2 = 420
	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}