package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, noOfcards int) (deck, deck) {
	return d[:noOfcards], d[noOfcards:]
}

func newDeck() deck {
	var suits = []string{"Club", "Spades", "Diamond", "Heart"}

	deck := deck{}

	for _, suit := range suits {
		for i := 1; i <= 13; i++ {
			deck = append(deck, suit+" of "+getCardNo(i))
		}
	}
	return deck
}

func getCardNo(i int) string {
	cardNo := ""
	if i == 1 {
		cardNo = "A"
	} else if i == 11 {
		cardNo = "J"
	} else if i == 12 {
		cardNo = "Q"
	} else if i == 13 {
		cardNo = "K"
	} else {
		cardNo = fmt.Sprintf("%d", i)
	}

	return cardNo
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) shuffle() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func loadFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(bs), ","))
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}
