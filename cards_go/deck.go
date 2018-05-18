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
type deck []string // equal deck === []string
type stringSlice []string
type color string

// d is a value recevier
func (d deck) print() {
	// var cardList = []string{}
	for i, card := range d {
		fmt.Println(i, card)
		// cardList = append(cardList, card+" ,")
	}
	// return cardList
}

// this func mean any time we call this func will return type deck's value
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Heart", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// d deck : is a recevier is talking about d is deck instance's value
// (deck, deck) : is talking about we will return two value and the value type is type deck
// handSize: is a parms and expect is a int
func deal(d deck, handSize int) (deck, deck) {
	// 0 ~ handSize, handSize ~ end
	return d[:handSize], d[handSize:]
}

func (c color) describeColor(describeWord string) string {
	return string(c) + " " + describeWord
}

func (d deck) toString() string {
	// let []string to string second arguments is a seperotor
	// first argument is what we need to translate is a slice
	return strings.Join([]string(d), ", ")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	byteSlice, err := ioutil.ReadFile(fileName)
	if err != nil {
		// option 1 - log error
		// option 2 - log error and entirly quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// string(byteSlice) we want a string and we give it a value not the string
	s := strings.Split(string(byteSlice), ", ")
	return deck(s)
}

func (d deck) shuffle() {
	// because golang generate the number use the same source,
	// so we need to give it a new source to complish when we call this method
	// every time we have a different result
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		// randomlize a number which in d.length
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
