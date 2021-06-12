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

	alex.updateName("Xavier")
	alex.lastName = "Cabrera"

	fmt.Printf("%+v", alex)
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

func (p *person) updateName(name string) {
	(*p).firstName = name
}
