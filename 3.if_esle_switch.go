package main

import (
	"fmt"
	"time"
)

func main() {
	age := 18

	if age <= 18 {
		fmt.Println("You can't visit here....")
	} else {
		fmt.Println("Welcome")
	}

	score := 85

	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}

	now := time.Now()

	day := now.Weekday()

	fmt.Println("Today is:", day)

	switch day {
	case time.Sunday:
		fmt.Println("Start of the work week")
	case time.Thursday:
		fmt.Println("Last of work week")
	case time.Friday, time.Saturday:
		fmt.Println("Weekend! Relax!")
	default:
		fmt.Println("Midweek")
	}
}
