package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// strings
	var name = "Syahrul Safarudin Hasan"
	fmt.Println(strings.ToUpper(name))
	fmt.Println(strings.ToLower(name))
	fmt.Println(strings.Contains(name, "Hasan"))

	// sort
	var fruits = []string{"Apple", "Grape", "Melon", "Banana"}
	fmt.Println(fruits, len(fruits))

	sort.Strings(fruits)
	sort.Sort(sort.Reverse(sort.StringSlice(fruits)))
	fmt.Println(fruits, len(fruits))
}
