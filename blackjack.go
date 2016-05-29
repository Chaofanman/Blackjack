package main

import (
	"fmt"
	"math/rand"
)

type Card struct {
	Value int
	Suit  string
}

type Player struct {
	Hand  []Card
	Score int
}

type Deck []Card
type House []Card
type User []Card

func newDeck() (deck Deck) {
	Value := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	Suit := []string{"Spades", "Clubs", "Diamonds", "Hearts"}

	for i := 0; i < len(Value); i++ {
		for j := 0; j < len(Suit); j++ {
			card := Card{Value[i], Suit[j]}
			deck = append(deck, card)
		}
	}
	return
}

func Shuffle(deck Deck) {
	for i := 1; i < len(deck); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			deck[r], deck[i] = deck[i], deck[r]
		}
	}
}

func Deal(deck Deck, player Player, count *int) (card, Card) {
	currentCard := deck[*count]
	player = append(player, currentCard)
	*count++
	return currentCard
}

func main() {
	deck := newDeck()
	Shuffle(deck)
	count := 0

	user := User

	Deal(deck, User, &count)
	fmt.Printf("%+v | count: %d\n", topCard, count)

}
