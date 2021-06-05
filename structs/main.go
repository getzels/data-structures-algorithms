package main

import "fmt"

func main() {
	alex := person{firstName: "Getzels",
		lastName: "De Los Santos",
		contact: contactInfo{
			email:   "get.delossantos@gmail.com",
			zipCode: 11076,
		},
	}

	fmt.Println(alex)

	alex.lastName = "Cabrera"
	alex.firstName = "Xavier"

	fmt.Println(alex)
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}
