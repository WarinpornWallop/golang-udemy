package main

import (
	"fmt"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	// Send an HTTP GET request to "http://google.com"
	resp, err := http.Get("http://google.com")

	// Check if there was an error during the HTTP request
	if err != nil {
		// Print the error message
		fmt.Println("Error:", err)

		// Exit the program with an exit status code of 1
		os.Exit(1)
	}

	// Create a byte slice (byte array) with a fixed size (99999)
	bs := make([]byte, 99999)

	// Read the response body into the byte slice
	resp.Body.Read(bs)

	// Print the response body as a string
	fmt.Println(string(bs))
}