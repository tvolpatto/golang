package main

import "fmt"

func main() {

	//Slice declaration
	cards := []string{"Nine of Diamonds", newCard()}

	//adding elements to a slice
	cards = append(cards, "Two of Hearts")

	for i, card := range cards {
		fmt.Println(i, card)

	}

}

func newCard() string {
	return "Ace of Spades"
}
