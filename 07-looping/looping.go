package main

import (
	"fmt"
)

func main() {
	// for loop
	fmt.Println("For Loop")
	for i := 0; i < 5; i++ {
		fmt.Println("Angka ", i)
	}

	// while loop (using for)
	fmt.Println("\nWhile Loop using for")
	j := 0
	for j < 5 {
		fmt.Println("Angka ", j)
		j++
	}

	// infinite loop (using for)
	fmt.Println("\nInfinite Loop using for")
	k := 0
	for {
		if k >= 5 {
			break
		}
		fmt.Println(k)
		k++
	}

	// for range loop
	fmt.Println("\nFor Range Loop")
	// looping through a string
	fmt.Println("\nLooping through a string")
	var strings = "Hello, World!"
	for index, char := range strings {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}

	// looping through an array
	fmt.Println("\nLooping through an array")
	var numbers = [5]int{1, 2, 3, 4, 5}
	for index, number := range numbers {
		fmt.Printf("Index: %d, Number: %d\n", index, number)
	}

	// looping through a slice
	fmt.Println("\nLooping through a slice")
	var sliceNumbers = numbers[1:4]
	for index, number := range sliceNumbers {
		fmt.Printf("Index: %d, Number: %d\n", index, number)
	}

	// looping through a map
	fmt.Println("\nLooping through a map")
	var person = map[string]string{
		"name":    "John",
		"country": "USA",
	}
	for key, value := range person {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	for range person {
		fmt.Println("Looping through map")
	}
}
