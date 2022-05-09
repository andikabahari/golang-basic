package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, value := range nums {
		total += value
	}
	return total
}

func main() {
	fmt.Println(sum(1,2,3))
}