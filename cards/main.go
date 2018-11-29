package main

import (
	"fmt"
)

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 2)

	hand.print()
	fmt.Println("hehehehehe")
	remainingCards.print()
}
