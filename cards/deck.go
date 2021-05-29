package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// first loop through the cardSuits
	for _, suit := range cardSuits {
		// while we are looping through cardSuits, loop through the cardValues
		for _, value := range cardValues {
			// create a new card by join the string values of the current cardSuit and cardValue together
			// and appending it to the cardSlice
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// this will give me a slice of strings - []string(d)
	// and we will condense the slice of strings and join the values with a comma between each value
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	// if there is an error, let us run this block
	if err != nil {
		// log error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// if there is no error, we take the byteslice(bs) to create a deck
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	// create the source/seed
	// creating a new number every time the program starts
	// so we can have a random number
	source := rand.NewSource(time.Now().UnixNano())
	// create new rand object
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		// the swapping part of code
		// take whatever is at "d[newPosition]" on the right side of the = sign
		// and assign it to d[i] as notated on the left side of the = sign
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
