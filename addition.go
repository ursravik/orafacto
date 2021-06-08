package main

import "fmt"

// main function
func main() {
	a := 3
	b := 4

	fmt.Println("!3 + !4 is =", factorial(a)+factorial(b))
}

// function to find factorial
func factorial(number int) int {
	var result int = 1

	for i := 1; i <= number; i++ {
		result = result * i
	}

	return result
}