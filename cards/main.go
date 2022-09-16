package main

func main() {

	cards := newDeckFromFile("myFoIle")

	cards.print()
}

func newCard() string {
	return "Ace of Spades"
}
