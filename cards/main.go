package main

func main() {
	cards := newDeck()
	// cards.saveToFile("my_cards")
	// cards := newDeckFromFile("my_card")
	cards.shuffle()
	cards.print()
}
