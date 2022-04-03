package main

import "fmt"

func main() {
	var months = [...]string{
		"Jan",
		"Feb",
		"Mar",
		"Apr",
		"May",
		"Jun",
		"Jul",
		"Aug",
		"Sep",
		"Oct",
		"Nov",
		"Dec",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// months[5] = "Edited"
	// fmt.Println(slice1) // changed

	// slice1[0] = "Edited"
	// fmt.Println(months) // changed

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Test") // new array
	fmt.Println(slice3)

	slice3[1] = "Not Dec"
	fmt.Println(slice3)
	fmt.Println(slice2) // not changed
	fmt.Println(months) // not changed

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Jazz"
	newSlice[1] = "Fusion"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(newSlice)

	// thisIsArray := [5]int{1,2,3,4,5}
	// thisIsArray := [...]int{1,2,3,4,5}
	// thisIsSlice := []int{1,2,3,4,5}
}