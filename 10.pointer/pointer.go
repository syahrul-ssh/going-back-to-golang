package main

import (
	"fmt"
)

func main() {
	// declaring a pointer
	var numberA int = 4
	var pointerA *int = &numberA

	fmt.Println("Value of numberA: ", numberA)
	fmt.Println("Address of numberA: ", &numberA)
	fmt.Println("Value of pointerA: ", pointerA)
	fmt.Println("Value pointed by pointerA: ", *pointerA)

	// changing the value using pointer
	*pointerA = 10
	fmt.Println("\nAfter changing the value using pointer:")
	fmt.Println("Value of numberA: ", numberA)
	fmt.Println("Value pointed by pointerA: ", *pointerA)

	// pointer as a parameter
	fmt.Println("\nUsing pointer as a parameter:")
	var numberB int = 5
	fmt.Println("Before change: ", numberB)
	changeValue(&numberB)
	fmt.Println("After change: ", numberB)
}

func changeValue(num *int) {
	*num = 20
}
