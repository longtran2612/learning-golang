package main

import (
	_ "fmt"
)

func main() {
	cards := newDeck()
	err := cards.saveFile("file_saved")
	if err != nil {
		return
	}
	cards2 := readFile("file_saved")
	cards2.print()

}
