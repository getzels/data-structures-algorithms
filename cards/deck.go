package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
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

func newDeckFromFile(filename string) deck {
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal("Error trying to read the file.", err)
		os.Exit(1)
	}

	deckstring := strings.Split(string(file), ",")

	return deck(deckstring)
}

func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
