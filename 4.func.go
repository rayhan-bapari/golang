package main

import "fmt"

// Basic Function
func sayHello() {
	fmt.Println("Hello, Go!")
}

// Function with Parameters
func greet(name string) {
	fmt.Println("Hello,", name)
}

// Function with Return Value
func add(a int, b int) int {
	return a + b
}

// Function with Multiple Return Values
func divide(a, b float64) (float64, float64) {
	quotient := a / b
	remainder := float64(int(a) % int(b))
	return quotient, remainder
}

// Variadic Function (variable number of arguments)
func sum(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

func main() {
	sayHello()

	greet("Rayhan")

	result := add(5, 3)
	fmt.Println("Sum is:", result)

	q, r := divide(10, 3)
	fmt.Println("Quotient:", q, "Remainder:", r)

	fmt.Println(sum(1, 2, 3))

	colors := [...]string{"red", "green", "blue"}

	for {
		var color []string

		for _, n := range colors {
			fmt.Println(n)
			color = append(color, n)
		}

		break 
	}
}
