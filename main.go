package main

import "fmt"

func main() {

	//one way to initialize map
	//var colors map[string]string

	colors := make(map[string]string) //built in function of Golang

	colors["white"] = "#ffffff"
	colors["red"] = "#ff0000"
	colors["green"] = "4bf745"

	delete(colors, "white")

	//another way to initialize map
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "4bf745",
	// }

	//fmt.Println(colors)

	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("%v, %v ", color, hex) //need printf for format directive %v
	}
}
