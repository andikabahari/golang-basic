package main

import "fmt"

func main() {
	var result = 60 + 9
	fmt.Println(result)

	var a = 6
	var b = 9
	var c = a * b
	fmt.Println(c)

	var i = 10
	i += 10
	fmt.Println(i)

	i++
	fmt.Println(i)

	var negative = -69
	var positive = +69
	fmt.Println(negative)
	fmt.Println(positive)

	var isTrue = !false
	fmt.Println(isTrue)
}