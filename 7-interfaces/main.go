package main

import "fmt"

// Define the 'bot' interface with a 'getGreeting' method
type bot interface {
	getGreeting() string // 'getGreeting' method return string
}
// Define the 'englishBot' struct and 'spanishBot' struct
type englishBot struct{}
type spanishBot struct{}

func main() {
	// Create instances of the 'englishBot' and 'spanishBot'
	eb := englishBot{}
	sb := spanishBot{}

	// Call 'printGreeting' with both bots
	printGreeting(eb) // Prints "Hi there!"
	printGreeting(sb) // Prints "Hola!"
}

//same function but other parameter

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }
// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }


// bot is interface that have getGreeting function
// Function 'printGreeting' takes a 'bot' interface and prints its greeting
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}


func (englishBot) getGreeting() string {
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
