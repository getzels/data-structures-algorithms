package main

import "fmt"

func main() {
	alex := person{firstName: "Getzels", lastName: "De Los Santos"}

	fmt.Println(alex)

	alex.lastName = "Cabrera"
	alex.firstName = "Xavier"

	fmt.Println(alex)
}

type person struct {
	firstName string
	lastName  string
}
