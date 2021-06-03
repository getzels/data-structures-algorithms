package main

import "fmt"

var numbers = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func main() {

	for i := 0; i < len(numbers); i++ {
		if numbers[i]%2 == 0 {
			fmt.Println(numbers[i], " is Even")
		} else {
			fmt.Println(numbers[i], " is Odd")
		}

	}
}
