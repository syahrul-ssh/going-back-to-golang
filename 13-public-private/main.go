package main

import (
	"13-public-private/library"
	"fmt"
)

func main() {
	library.SayHello()
	//library.introduce("John Doe") // This will cause an error because introduce is unexported

	var person = library.Person{Name: "Alice", Age: 30} // This will work because Person is exported
	fmt.Println("Person struct created:", person)
}
