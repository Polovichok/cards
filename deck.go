package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Jack", "Queen", "King", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}
	for _, suit := range cardValues {
		for _, value := range cardSuits {
			cards = append(cards, suit+" of "+value)
		}
	}
	cards = append(cards, "Colorful Joker", "Black and White Joker")
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
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	byteSize, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	sliceString := strings.Split(string((byteSize)), ",")
	return deck(sliceString)
}

func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	for i := range d {
		newPosition := random.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
