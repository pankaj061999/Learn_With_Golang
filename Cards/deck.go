package main

import "fmt"

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
