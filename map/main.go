package main

import "fmt"

func main() {
	// var colors map[string]string

	colors := make(map[int]string)
	colors[10] = "#fffff"

	delete(colors, 10)

	// colors := map[string]string{
	// 	"red": "#ff0000",
	// 	"green": "#4bf745",
	// }

	fmt.Println(colors)
}