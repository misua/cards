package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"
	card := newCard()
	//card = "Five of Diamons" if you want to assign new value you dont need to put :

	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
