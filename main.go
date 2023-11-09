package main

func main() {
	cards := newDeck()
	cards.shuffle()
	hand, remainingCards := deal(cards, 6)
	hand.print()
	remainingCards.print()
}
