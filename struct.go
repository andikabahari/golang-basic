package main

import "fmt"

type Customer struct {
	Name, Address	string
	Age				int
}

func main() {
	var john Customer
	john.Name = "John Doe"
	john.Address = "Bandung, West Java"
	john.Age = 99
	fmt.Println(john)

	pat := Customer{
		Name:		"Pat M.",
		Address:	"Bandung, West Java",
		Age:		99,
	}
	fmt.Println(pat)
}