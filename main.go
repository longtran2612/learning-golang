package main

import _ "fmt"

func main() {
	cards := newDeck()
	hand, remainingDeck := deal(cards, 3)
	hand.print()
	remainingDeck.print()
}
