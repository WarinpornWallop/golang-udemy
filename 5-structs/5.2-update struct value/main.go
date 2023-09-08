package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	//if not declare, it follow key by below
	// alex := person{"Alex","Anderson"}
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)
}