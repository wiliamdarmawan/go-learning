package main

func main() {
	// cards := newDeck()
	// cards.saveToFile("my_cards")
	
	test := newDeckFromFile("my_cards")
	test.print()
}

func newCard() string {
	return "Five of Diamonds"
}