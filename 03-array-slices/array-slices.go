package main 

import "fmt"

func main() {
	// array
	var fruits [4]string
	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[2] = "Grape"
	fruits[3] = "Melon"
	fmt.Println(fruits)

	fruits2 := [4]string{"Apple", "Banana", "Grape", "Melon"}
	fmt.Println(fruits2, len(fruits2))

	fruits2[0] = "Pineapple"
	fmt.Println(fruits2, len(fruits2))

	// slice
	names := []string{"Syahrul", "Safarudin", "Hasan"}
	fmt.Println(names, len(names))

	names = append(names, "Fauzan")
	fmt.Println(names, len(names))

	// slice from array
	var fruits3 = [4]string{"Apple", "Banana", "Grape", "Melon"}
	var sliceFruits = fruits3[0:2]
	fmt.Println(sliceFruits, len(sliceFruits))

	sliceFruits = append(sliceFruits, "Pineapple")
	fmt.Println(sliceFruits, len(sliceFruits))
}