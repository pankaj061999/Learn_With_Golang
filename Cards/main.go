// main.go
package main

import (
	"fmt"
)

// Here create a reusable package
func ConcatString(s1, s2 string) string {
	return s1 + " " + s2
}

func main() {
	var FirstName string = "Pankaj Kumar"
	var LastName string = "Meena"

	FullName := ConcatString(FirstName, LastName)

	var Card string = "arjiun"
	Card = "arjun"

	Ram := 123

	greeting := GreetingName()

	fmt.Println(FullName)
	fmt.Println(Card)
	fmt.Println(Ram)
	fmt.Println(greeting)

	// create a slice and a array
	// array
	BlogList := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(BlogList)

	// slice
	BlogData := []int{1, 2, 4, 9}
	// add element this slice
	BlogData = append(BlogData, 625)

	// Print all blogData value through for loop
	for i := 0; i < len(BlogData); i++ {
		fmt.Println(BlogData[i], i)
	}

	for i, BlogDataelem := range BlogData {
		fmt.Println("Second method", BlogDataelem, i)
	}
	fmt.Println(BlogData)
}
