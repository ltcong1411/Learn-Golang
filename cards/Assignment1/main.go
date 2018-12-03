package main

import (
	"fmt"
	"strconv"
)

type deck []int
type deckString []string

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(numbers)
	cards := checkCard(numbers)
	for _, d := range cards {
		fmt.Println(d)
	}
}

func checkCard(d deck) []string {
	inf := []string{}
	for _, card := range d {
		if card%2 == 1 {
			inf = append(inf, strconv.Itoa(card)+" is odd")
		} else {
			inf = append(inf, strconv.Itoa(card)+" is even")
		}
	}
	return inf
}
