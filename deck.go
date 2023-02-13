package main

import (
	"fmt"
	"os"
	"strings"
	_ "strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubes"}
	cardValues := []string{"Ace", "One", "Two", "Three"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (deck deck) print() {
	for i, card := range deck {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
func (deck deck) toString() string {
	return strings.Join(deck, ",")
}
func (deck deck) saveFile(filename string) error {
	return os.WriteFile(filename, []byte(deck.toString()), 0666)
}

func readFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	s := strings.SplitAfter(string(bs), ",")
	return deck(s)
}
