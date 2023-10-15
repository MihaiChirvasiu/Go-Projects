package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo // or contactInfo without the name contact
}

func main() {
	// var personAlex person -> initializes the variable personAlex of type person with Zero Values for both fields
	// personAlex.firstName = "Alex"
	// personAlex.lastName = "Anderson" we can assign values to the fields like in any other language
	personAlex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(personAlex)
	fmt.Printf("%+v", personAlex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{ // contactInfo: contactInfo
			email:   "jim.party@gmail.com",
			zipCode: 240158,
		},
	}
	jim.print()
	// jimPointer := &jim
	// jimPointer.updateFirstName("Jameson")
	jim.updateFirstName("jimmy")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateFirstName(newName string) {
	(*pointerToPerson).firstName = newName
}
