package main

import "fmt"

func main() {

	firstMap := map[string]string{
		"red":   "#FF0000",
		"white": "#FFFFFF",
		"black": "#000000",
	}

	creatingMap()

	creatingEmptyMap()

	printMap(firstMap)

}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("Key: " + key + "Value: " + value)
	}
}

func creatingMap() {
	firstMap := map[string]string{
		"red":   "#FF0000",
		"white": "#FFFFFF",
	}

	fmt.Println(firstMap)
}

func creatingEmptyMap() {
	empMap := make(map[int]string)

	fmt.Println(empMap)

	empMap[10] = "Ten"

	fmt.Println(empMap)
}
