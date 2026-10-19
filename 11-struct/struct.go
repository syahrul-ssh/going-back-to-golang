package main

import (
	"fmt"
)

type Person struct {
	Name    string
	Age     int
	Country string
}

func main() {
	// creating a struct instance
	person1 := Person{
		Name:    "Alice",
		Age:     30,
		Country: "USA",
	}

	fmt.Println("Person 1:", person1)

	// accessing struct fields
	fmt.Println("Name:", person1.Name)
	fmt.Println("Age:", person1.Age)
	fmt.Println("Country:", person1.Country)

	// modifying struct fields
	person1.Age = 31
	fmt.Println("\nAfter modifying Age:")
	fmt.Println("Person 1:", person1)

	// using pointers with structs
	person2 := &person1
	fmt.Println("\nPerson 2 (using pointer):", *person2)

	// modifying struct fields using pointer
	person2.Age = 32
	fmt.Println("\nAfter modifying Age using pointer:")
	fmt.Println("Person 1:", person1)
	fmt.Println("Person 2:", *person2)
}
