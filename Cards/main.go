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

	// greeting := GreetingName()

	fmt.Println(FullName)
	fmt.Println(Card)
	fmt.Println(Ram)
	// fmt.Println(greeting)

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

	// create a var this take two argument

	Namereturn := Person{"Pankaj My Name", "arjun My second Name"}

	Namereturn.PrintPerson()

	// call common function based on deck types

	result := NewDeckFunc()
	result.PrintPerson()

	// create a slice and extract value based on syntax

	res := []string{"Pankaj", "arjun", "Roy", "Tony"}
	fmt.Println(res[0:2]) //[Pankaj arjun]

	fmt.Println(res[0:]) //[Pankaj arjun Roy Tony]

	fmt.Println(res[:])  //[Pankaj arjun Roy Tony]
	fmt.Println(res[1:]) //[arjun Roy Tony]
	// fmt.Println(res[-1:])  .\main.go:71:19: invalid argument: index -2 (constant of type int) must not be negative

	// bytes slices
	City := "Jaipur Pnkaj"
	fmt.Println([]byte(City)) //[74 97 105 112 117 114 32 80 110 107 97 106]

	// join string
	finalJOinString := NewDeckFunc()
	resultvalue := finalJOinString.StringJoinFunc()

	fmt.Println(resultvalue)

	// save string into file pankaj.txt file
	finalSaveStringintoFIle := finalJOinString.SavedFile("pankaj.txt")
	fmt.Println("Print saved file", finalSaveStringintoFIle)

	// Read File strings into Pankaj.txt file
	readFile := ReadStringFile("pankaj.txt")

	fmt.Println("Read File Value ==>>", readFile)

	// shuffling name

	resultShuffling := NewDeckFunc()
	resultShuffling.shuffiingFunc()
	resultShuffling.PrintPerson()

	// return random nuber
	resultShuffling.randomGenrateFunc()
	resultShuffling.PrintPerson()

	// check slice nuber even or odd

	numbers := ResultNuber{2, 7, 4, 9, 6}
	messages := numbers.checkEvenOrOddNuber()

	for _, message := range messages {
		fmt.Println(message)
	}
}
