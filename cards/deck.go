package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

//Create a deck type as a slice of string
type deck []string

func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, sideHand int) (deck, deck) {
	return d[:sideHand], d[sideHand:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

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

func (d deck) deckToFile(filename string) int {
	error := ioutil.WriteFile(filename, []byte(d.toString()), 0644)

	if error != nil {
		log.Fatal("Error try to save the file", error)
		return 1
	}

	return 0
}

func newDeckFromFile(filename string) {
	file, err := ioutil.ReadFile(filename)

	if err != nil {

	}
}
