package main

import (
	"fmt"
)

func main() {
	var person = map[string]string{
		"name":    "John",
		"country": "USA",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["country"])

	person["age"] = "30"
	fmt.Println(person)

	delete(person, "age")
	fmt.Println(person)
}
