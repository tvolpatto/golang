package main

func main() {

	//Slice declaration
	cards := newDeck()

	cards.saveToFile("myFIle")
}

func newCard() string {
	return "Ace of Spades"
}
