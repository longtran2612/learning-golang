package main

import (
	_ "fmt"
)

func main() {
	cards := newDeck()
	err := cards.saveFile("tesst save file")
	if err != nil {
		return
	}

}
