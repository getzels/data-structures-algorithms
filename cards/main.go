package main

var card string = "Ace of Spaces"

func main() {

	cards := newDeckFromFile("deckFile.txt")

	cards.print()
}
