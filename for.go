package main

import "fmt"

func main() {
	counter := 1
	for counter <= 10 {
		fmt.Println("Iteration", counter)
		counter++
	}

	names := []string{"One", "Two", "Three"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}
}