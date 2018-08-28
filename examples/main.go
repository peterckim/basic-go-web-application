package main

func main() {
	// var card string = "Ace of Spades"
	// compiler infers what datatype

	cards := newDeck()
	cards.shuffle()
	cards.print()

	// cards := newDeckFromFile("my_cards")
	// cards.print()

	// hand, remainingDeck := deal(cards, 5)

	// hand.print()
	// remainingDeck.print()
}
