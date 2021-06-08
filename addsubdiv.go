package main

import "fmt"

// main function
func main() {
	a := 10
	b := 5
	c := 3
	d := 4
	fmt.Println("10! - 5!/3! + 4! is =", factorial(a)-factorial(b)/factorial(c)+factorial(d))
}

// function to find factorial
func factorial(number int) int {
	var result int = 1

	for i := 1; i <= number; i++ {
		result = result * i
	}

	return result
}