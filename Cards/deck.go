package main

import (
	"fmt"
	"os"
	"strings"
)

// Define a type
type Person []string

// Method with a value receiver
func (p Person) PrintPerson() {

	for i, card := range p {
		fmt.Println("Print P", i, p, card)

	}
}

// create a function this work on based on our Person type

func NewDeckFunc() Person {
	Cards := Person{}

	cardsSuits := []string{"Pankaj", "suri", "Arjun", "Meena"}
	cardsSet := []string{"Jaipur", "Bundi", "New Delhi", "Roorkee"}

	for _, suits := range cardsSuits {
		for _, setsValue := range cardsSet {
			// fmt.Println("Print For loop Value", suits, setsValue)
			Cards = append(Cards, suits+" of  "+setsValue)
		}
	}

	return Cards
}

// join to string help of string JOIn  prebuilt function

// func Join(s []string, sep string) string  syntax

func (p Person) StringJoinFunc() string {

	return strings.Join([]string(p), ",")
}

// write File if don't exist if exist then updated

func (p Person) SavedFile(filename string) error {

	return os.WriteFile(filename, []byte(p.StringJoinFunc()), 0644)
}

// ReadFile if exist help of os.ReadFile Prebuilkt function

func ReadStringFile(fileName string) Person {
	bs, err := os.ReadFile(fileName)

	if err != nil {

		fmt.Println("Error File Read ==>>", err)
		os.Exit(1) // if this value non zero denots return error if value zero then don't have error
	}

	s := strings.Split(string(bs), " ")
	return Person(s)

}
