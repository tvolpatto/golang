package main

import (
	"fmt"
)

func main() {
	// colors := map[string]string{
	// 	"red":  "#FF0000",
	// 	"blue": "#EE1244",
	// }

	//var colors map[string]string

	colors := make(map[string]string)

	colors["white"] = "#AAAAAA"

	delete(colors, "white")

	fmt.Println(colors)
}
