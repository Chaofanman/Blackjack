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

func newDeck() (deck Deck) {
	Value := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	Suit := []string{"Spades", "Clubs", "Diamonds", "Hearts"}

	//Initialize every card in the deck
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

//Pass pointer to be able to dereference
func Deal(deck Deck, count *int) (card Card) {
	currentCard := deck[*count]
	*count++
	return currentCard
}

//Allows to dynamically add to the Player Hand
//which is an array of cards
func (player *Player) addCard(card Card) []Card {
	player.Hand = append(player.Hand, card)
	return player.Hand
}

func (player *Player) countScore() int {
	for i := 0; i < len(player.Hand); i++ {
		player.Score += player.Hand[i].Value
	}
	return player.Score
}

func main() {
	//Initialize new deck using a function
	deck := newDeck()
	Shuffle(deck)
	count := 0

	newCard := Deal(deck, &count)
	newCard2 := Deal(deck, &count)

	fmt.Println(newCard)

	//Initialize a struct Player
	user := Player{}

	user.addCard(newCard)
	user.addCard(newCard2)

	user.countScore()

	fmt.Println(user.Hand)
	fmt.Println(user.Score)

}
