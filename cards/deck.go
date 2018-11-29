package main

import (
	"fmt"
	"strings"
)

type deck []string

func newDeck() deck {
	//fruitFillter := deck{}
	cards := deck{}

	cardSuits := []string{"Apple", "Banana", "Orange", "Grasp"}
	cardValues := []string{"One", "Two", "Three", "Four"}

	for i, suit := range cardSuits {
		// for _, value := range cardValues {
		cards = append(cards, cardValues[i]+" of "+suit)
		// }
	}
	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
func (d deck) saveToFile(filename string) {

}
