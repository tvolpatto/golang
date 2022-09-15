package main

func main() {

	//Slice declaration
	cards := newDeck()

	//adding elements to a slice
	cards = append(cards, "Two of Hearts")

	cards.print()

}

func newCard() string {
	return "Ace of Spades"
}
