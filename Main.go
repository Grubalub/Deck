package main

import (
	"fmt"

	"github.com/Grubalub/Cards"
)

func main() {
	cards := Cards.New(Cards.Deck(3), Cards.Shuffle)
	var card Cards.Card
	for i := 0; i < 10; i++ {
		card, cards = cards[0], cards[1:]
		fmt.Println(card)
	}
}
