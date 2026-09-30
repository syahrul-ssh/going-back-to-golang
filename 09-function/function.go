package main

import (
	"fmt"
)

func main() {
	// function without parameter and return value
	fmt.Println("Function without parameter and return value")
	sayHello()

	// function with parameter and without return value
	fmt.Println("\nFunction with parameter and without return value")
	sayHelloTo("John")

	// function with parameter and return value
	fmt.Println("\nFunction with parameter and return value")
	result := add(5, 10)
	fmt.Println("Result: ", result)

	// function with multiple parameters and return values
	fmt.Println("\nFunction with multiple parameters and return values")
	sum, product := calculate(5, 10)
	fmt.Println("Sum: ", sum)
	fmt.Println("Product: ", product)

	// function with variadic parameters
	fmt.Println("\nFunction with variadic parameters")
	avg := average(1.0, 2.0, 3.0, 4.0, 5.0)
	fmt.Println("Average: ", avg)
}

func sayHello() {
	fmt.Println("Hello!")
}

func sayHelloTo(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func add(a int, b int) int {
	return a + b
}

func calculate(a int, b int) (int, int) {
	sum := a + b
	product := a * b
	return sum, product
}

func average(numbers ...float64) float64 {
	var total float64
	for _, number := range numbers {
		total += number
	}
	return total / float64(len(numbers))
}
