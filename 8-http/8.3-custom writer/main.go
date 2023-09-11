package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Define a logWriter type
type logWriter struct{}

func main() {
	// Send an HTTP GET request to "http://google.com"
	resp, err := http.Get("http://google.com")
	if err != nil {
		// If there's an error during the HTTP request, print the error and exit with status code 1
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Create an instance of logWriter
	lw := logWriter{}

	// Copy the response body to the logWriter instance
	// This effectively invokes the Write method of logWriter
	io.Copy(lw, resp.Body)
}

// Implement the Write method for logWriter, satisfying the io.Writer interface
// This method will be called by io.Copy for each chunk of data read from the response body
func (logWriter) Write(bs []byte) (int, error) {
	// Print the received data as a string
	fmt.Println(string(bs))

	// Print the number of bytes written (length of the byte slice)
	fmt.Println("Just wrote this many bytes:", len(bs))

	// Return the number of bytes written and nil error
	return len(bs), nil
}
