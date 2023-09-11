package main

import (
	"fmt"
	"io"
	"os"
)

func main(){
	//go run main.go tonnamwarineiei
	//[<path>/main.exe tonnamwarineiei]
	//fmt.Println(os.Args)

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
//buld main file
//go build main.go
//main.exe main.go <- run for read code