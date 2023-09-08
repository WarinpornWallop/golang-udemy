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
		lastName:  "Warin",
		contactInfo: contactInfo{
			email:   "tonnamwarin@gmail.com",
			zipCode: 12365,
		},
	}
	fmt.Println(tonnam)
	// %+v is a format specifier that tells Printf to print the value of tonnam with additional information 
	// about the fields in the struct or the keys and values in a map. 
	// It is often used for debugging and inspecting data structures 
	// because it includes field names and key-value pairs when printing structs and maps.
	fmt.Printf("%+v", tonnam)
	// print -> {firstName:Tonnam lastName:Warin contactInfo:{email:tonnamwarin@gmail.com zipCode:12365}}
}
