package main

import (
	"fmt"
	"strings"
)

type student struct {
	name  string
	grade int
}

func (s student) sayHello() {
	fmt.Println("hello", s.name)
}

func (s student) getNameAt(i int) string {
	return strings.Split(s.name, " ")[i-1]
}

// method pointer
func (s student) changeName1(name string) {
	fmt.Println("on changeName1, name is changed to", name)
	s.name = name
}

func (s *student) changeName2(name string) {
	fmt.Println("on changeName2, name is changed to", name)
	s.name = name
}

func main() {
	var s1 = student{"john doe", 10}
	s1.sayHello()

	var name = s1.getNameAt(2)
	fmt.Println("nickname:", name)

	s1.changeName1("Jason Doe")
	fmt.Println("s1 name after changeName1:", s1.name)

	s1.changeName2("Jason Martin")
	fmt.Println("s1 name after changeName2:", s1.name)

	var s2 = &student{"jane doe", 11}
	s2.sayHello()

	name = s2.getNameAt(1)
	fmt.Println("nickname:", name)

	s2.changeName1("Jane Martin")
	fmt.Println("s2 name after changeName1:", s2.name)

	s2.changeName2("Jane Smith")
	fmt.Println("s2 name after changeName2:", s2.name)
}
