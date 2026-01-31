package main

import "fmt"

// add returns the sum of two integers
func add(a, b int) int {
	return a + b
}

// multiply returns the product of two integers
func multiply(a, b int) int {
	return a * b
}

func main() {
	// Greeting
	fmt.Println("🚀 Welcome to Go Programming!")
	fmt.Println("===========================")

	// Variables
	name := "Prakhar"
	age := 25
	fmt.Printf("Hello, %s! You are %d years old.\n\n", name, age)

	// Simple calculations
	num1 := 10
	num2 := 5

	sum := add(num1, num2)
	product := multiply(num1, num2)

	fmt.Printf("Simple Math Operations:\n")
	fmt.Printf("%d + %d = %d\n", num1, num2, sum)
	fmt.Printf("%d × %d = %d\n\n", num1, num2, product)

	// Arrays and loops
	fmt.Println("Counting to 5:")
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println("\n\n✅ Program completed successfully!")
}
