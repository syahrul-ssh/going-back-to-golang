package main

import (
	"fmt"
	"math"
)

func main() {
	// aritmatic operator
	var a = 10
	var b = 3

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	var c = 10.0
	var d = 3.0

	fmt.Println(c + d)
	fmt.Println(c - d)
	fmt.Println(c * d)
	fmt.Println(c / d)

	var e = 10
	var f = 3

	fmt.Println(math.Pow(float64(e), float64(f)))

	// comparison operator
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)

	// logical operator
	var isSuccess = true
	var isFailed = false

	fmt.Println(isSuccess && isFailed)
	fmt.Println(isSuccess || isFailed)
	fmt.Println(!isSuccess)
}
