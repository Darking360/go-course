package main

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")
	readCards := newDeckFromFile("my_cards")
	readCards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
