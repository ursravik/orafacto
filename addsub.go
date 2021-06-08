package main

import "fmt"

// main function
func main() {
	a := 5
	b := 6
	c := 7
	fmt.Println("!5 + !6 - !7 is =", factorial(a)+factorial(b)-factorial(c))
}

// function to find factorial
func factorial(number int) int {
	var result int = 1

	for i := 1; i <= number; i++ {
		result = result * i
	}

	return result
}