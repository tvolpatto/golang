package main

func main() {

	//Slice declaration
	cards := newDeck()

	// calling deal
	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()

}

func newCard() string {
	return "Ace of Spades"
}
