package main

import (
	"fmt"
)

func main() {

	// different ways of declaring a map
	colors := make(map[string]string)

	// var colors map [string]string{}

	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#00ff00",
	// }

	colors["white"] = "#ffffff"

	delete(colors, "white")

	fmt.Println(colors)
}
