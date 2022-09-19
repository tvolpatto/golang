package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create  a new type of 'deck'
// which is a slice of strings
type deck []string

// Create a new deck
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}

	// we use the underscore if we don't need the indexes
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)

		}
	}

	return cards
}

// Print the Deck
// function with a receiver
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)

	}

}

// slice range sintax:
// deck[startIndexIncluding:upTotheIndexNotIncluded]
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// convert the []deck to a []string  to join and return a single string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// convert to []byte and save to a file
func (d deck) saveToFile(filename string) error {

	return ioutil.WriteFile(filename, []byte(d.toString()), 06666)
}

// load from a file
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("ERROR: ", err)
		os.Exit(1)
	}

	//convert the []byte to a string and to a []string
	s := strings.Split(string(bs), ",")

	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPos := r.Intn(len(d) - 1)

		d[i], d[newPos] = d[newPos], d[i]
	}

}
