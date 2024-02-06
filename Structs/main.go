package main

import "fmt"

type ContactInfo struct {
	email   string
	zipCode int64
}

// Struct definition
type Person struct {
	firstName string
	lastName  string
	// Custom type
	contact ContactInfo
}

// Receiver function with pointer receiver
func (p *Person) updateName(newName string) {
	(*p).firstName = newName
}

// Receiver function with value receiver
func (p Person) PrintResult() {
	fmt.Println("Print user name", p)
}

func main() {
	// Declare a variable of type Person
	var users Person

	// Update struct values
	users.firstName = "Pankaj Kumar"
	users.lastName = "Meena"

	// Create a new instance of Person
	UsersName := Person{
		firstName: "Pankaj",
		lastName:  "Meena",
		contact: ContactInfo{
			email:   "hel@gmail.com",
			zipCode: 12004,
		},
	}

	UserNamePointer := &users

	// Print the new instance
	fmt.Println("Hello", UsersName, UserNamePointer)

	// Print the internal key-value pairs with key names
	fmt.Printf("%+v\n", users)

	// Call the receiver function to update name
	users.updateName("Arjun Kumar")

	// Call the receiver function to print the result
	users.PrintResult()
}
