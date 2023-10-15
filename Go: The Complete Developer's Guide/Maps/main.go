package main

import "fmt"

func main() {
	// var colors map[string]string initializes the colors map with its zero values
	// the above statement is equivalent with
	// colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}
	colors["white"] = "#ffffff"
	delete(colors, "white")
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
