package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	addColors(colors)

	delete(colors, "green")

	printColorsMap(colors)
}

func addColors(colors map[string]string) {
	colors["grey"] = "#bbbbbb"
	colors["white"] = "#ffffff"
}

func printColorsMap(colors map[string]string) {
	for color, hex := range colors {
		fmt.Printf("Hex code for color %v is %v\n", color, hex)
	}
}
