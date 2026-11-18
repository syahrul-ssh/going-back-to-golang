package library

import "fmt"

type Person struct {
	Name string
	Age  int
}

func SayHello() {
	fmt.Println("Hello from library")
	introduce("John Doe") // This will work because introduce is called within the same package
}

func introduce(name string) {
	fmt.Println("Hello, my name is", name)
}
