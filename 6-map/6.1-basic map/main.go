package main

import "fmt"

func main() {
	//var colors map[string]string
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// }
	colors := make(map[int]string)
	// colors["white"] = "#ffffff"
	 colors[10] = "#ffffff"
	delete(colors, 10)
	fmt.Println(colors)
}
