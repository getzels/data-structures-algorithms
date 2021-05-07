package main

import "fmt"

//Create a deck type as a slice of string
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuit := []string{"Spades", "Diamonds", "Heart", "Clubs"}
	cardValue := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuit {
		for _, value := range cardValue {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}

func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}
}
