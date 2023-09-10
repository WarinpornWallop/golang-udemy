package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	tonnam := person{
		firstName: "Tonnam",
		lastName:  "Warineiei",
		contactInfo: contactInfo{
			email:   "tonnamwarin@gmail.com",
			zipCode: 12345,
		},
	}
	// tonnamPointer := &tonnam
	// tonnamPointer.updateName("Tonnamwarineiei")
	//from 2 line above, you can use the below line that is a shortcut
	tonnam.updateName("Tonnamwarineiei2")
	tonnam.print()
}
// Define a method 'updateName' for the 'person' struct
// This method takes a pointer to a 'person' as its receiver and a new first name as a parameter
// It updates the 'firstName' field of the person with the new first name
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}