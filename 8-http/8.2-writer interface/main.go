package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)


func main() {
	// Send an HTTP GET request to "http://google.com"
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Copy the response body to the standard output (stdout)
	// This effectively prints the response body to the console
	io.Copy(os.Stdout, resp.Body)
}
