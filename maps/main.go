package main

import (
	"fmt"
)

func main() {
	//var colors map[string]string

	//colors := make(map[string]string)

	colors := map[string]string{
		"red":  "#FF0000",
		"blue": "#EE1244",
	}

	colors["white"] = "#AAAAAA"

	//delete(colors, "white")

	printMap(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Printf(" The hex color for %v is %v \n", key, value)

	}
}
