package main

import "fmt"

func main() {
	colors := make(map[int]string)
	// var colors map[string]string
	// colors := map[string]string{
	//	"red":   "ff0000",
	//	"green": "4bf745",
	//}
	colors[10] = "#ffffff"

	delete(colors, 10)

	fmt.Println(colors)
}
