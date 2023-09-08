package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

//Create a new type of 'deck'
//which is a slice of strings
type deck []string
//Create Deck
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards

}
//show card
//recieve deck type in this func ; d = parameter
func (d deck) print(){
	for i, card := range d {
		fmt.Println(i,card)
	}
}
//devided deck in hand and on desk
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d),",")
}
//0666 = everyone can read and write (permision)
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename) //err = value of type 'error'. If nothing went wrong, it will have a value of 'nil'
	if err != nil {
		// Option 1 -> log the error and return a call to newDeck()
		// Option 1 -> log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		//rand = random number for index
		//newPosition := rand.Intn(len(d) - 1)
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}