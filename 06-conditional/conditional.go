package main

import (
	"fmt"
)

func main() {
	// if statement
	var score = 80

	if score > 80 {
		fmt.Println("A")
	} else if score > 70 {
		fmt.Println("B")
	} else if score > 60 {
		fmt.Println("C")
	} else if score > 50 {
		fmt.Println("D")
	} else {
		fmt.Println("E")
	}

	// switch statement
	var grade = "B"

	switch grade {
	case "A":
		fmt.Println("Excellent")
	case "B":
		fmt.Println("Good")
	case "C":
		fmt.Println("Fair")
	case "D", "E":
		fmt.Println("Poor")
	default:
		fmt.Println("Fail")
	}
}
