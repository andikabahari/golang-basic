package main

import "fmt"

func main() {
	var exam = 80
	var attendances = 80

	var passExam = exam >= 80
	var passAttendances = attendances >= 80
	fmt.Println(passExam)
	fmt.Println(passAttendances)

	var pass = passExam && passAttendances
	fmt.Println(pass)
}