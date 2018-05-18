package main

import (
	"fmt"
)

func main() {
	// map[string]string is mean key is string type, and value is string type
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// colors := make(map[string]string)
	// colors["white"] = "#ffffff"

	// colors := make(map[int]string)
	// colors[10] = "#000000"

	// delete(colors, 10) // will delete the key and value

	// fmt.Println(colors)
	printMap(colors)
}

func printMap(argument map[string]string) {
	for key, value := range argument {
		fmt.Println("key is: ", key, " value is: ", value)
	}
}
