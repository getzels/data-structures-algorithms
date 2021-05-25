package main

import (
	"os"
	"testing"
)

var d deck = newDeck()

func TestNewDeck(t *testing.T) {

	if len(d) != 16 {
		t.Errorf("Error validating the deck len most be 16, but got %v", len(d))
	}
}

func TestNewDeckFirstElement(t *testing.T) {
	if d[0] != "Spades of Ace" {
		t.Errorf("Error the with the first value, expected Spades of Ace, but got %v", d[0])
	}
}

func TestNewDeckLastElement(t *testing.T) {
	if d[len(d)-1] != "Clubs of Four" {
		t.Errorf("Error the with the first value, expected Clubs of Four, but got %v", d[len(d)-1])
	}
}

func TestDeckToFile(t *testing.T) {

	os.Remove("_testdecktofile")

	d := newDeck()

	t.Log("lenght of deck: ", len(d))

	err := d.deckToFile("_testdecktofile")

	if err == 1 {
		t.Errorf("Error try to save the deck to file (_testdecktofile), %v", err)
	}
}

func TestNewDeckFromFile(t *testing.T) {
	d := newDeckFromFile("_testdecktofile")

	if d == nil {
		t.Errorf("Error reading the file.. %v", d)
	} else if len(d) != 16 {
		t.Errorf("Error validating the deck len most be 16, but got %v", len(d))
	}

	os.Remove("_testdecktofile")
}
