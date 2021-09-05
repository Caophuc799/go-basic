package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	// cards.saveToFile("./text.txt")
	// hand, remainingDeck := deal(cards, 5)
	// test := newDeckFromFile("./text.txt")
	// test.print()
	// hand.print()
	// remainingDeck.print()

}

func newCard() string {
	return "Five of diamonds"
}
