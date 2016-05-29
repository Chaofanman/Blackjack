package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Card struct {
	Value int
	Suit  string
}

type Player struct {
	Hand  []Card
	Name  string
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
	player.Score = player.countScore()

	return player.Hand
}

func (player *Player) countScore() int {
	score := 0
	//Value of Jack, King, Queen is 10
	for i := 0; i < len(player.Hand); i++ {
		if player.Hand[i].Value >= 10 {
			score += 10
		} else {
			if player.Hand[i].Value == 1 {
				score += player.handleAce()
			} else {
				score += player.Hand[i].Value
			}
		}
	}
	return score
}

func (player *Player) handleAce() int {
	var score int
	if player.Score < 21 {
		score += 11
	} else {
		score += 1
	}
	return score
}

func whoWins(user Player, house Player) {
	switch {
	case (user.Score > house.Score && user.Score <= 21) && house.Score > 21:
		fmt.Println("YOU WIN WIN WIN WIN !")
	case (user.Score < house.Score && user.Score <= 21) && house.Score > 21:
		fmt.Println("YOU WIN WIN WIN WIN !")
	case user.Score > 21:
		fmt.Println("YOU LOSE LOSE LOSE LOSE")
	case user.Score < house.Score && house.Score <= 21:
		fmt.Println("YOU LOSE LOSE LOSE LOSE")
	case user.Score == house.Score:
		fmt.Println("Tie")
	}
}

func printState(user Player, house Player) {
	fmt.Println("\nYou have the: ")
	for i := 0; i < len(user.Hand); i++ {
		//Mimicing %v for Ace, Jack, Queen, and King
		if user.Hand[i].Value == 1 {
			fmt.Printf("{Ace %s}\n", user.Hand[i].Suit)
		}
		if user.Hand[i].Value == 11 {
			fmt.Printf("{Jack %s}\n", user.Hand[i].Suit)
		}
		if user.Hand[i].Value == 12 {
			fmt.Printf("{Queen %s}\n", user.Hand[i].Suit)
		}
		if user.Hand[i].Value == 13 {
			fmt.Printf("{King %s}\n", user.Hand[i].Suit)
		}
		if user.Hand[i].Value >= 2 && user.Hand[i].Value <= 10 {
			fmt.Printf("%v\n", user.Hand[i])
		}
	}
	fmt.Println("\nHouse has: ")
	//Bad practice
	for i := 0; i < len(house.Hand); i++ {
		//Mimicing %v for Ace, Jack, Queen, and King
		if house.Hand[i].Value == 1 {
			fmt.Printf("{Ace %s}\n", house.Hand[i].Suit)
		}
		if house.Hand[i].Value == 11 {
			fmt.Printf("{Jack %s}\n", house.Hand[i].Suit)
		}
		if house.Hand[i].Value == 12 {
			fmt.Printf("{Queen %s}\n", house.Hand[i].Suit)
		}
		if house.Hand[i].Value == 13 {
			fmt.Printf("{King %s}\n", house.Hand[i].Suit)
		}
		if house.Hand[i].Value >= 2 && house.Hand[i].Value <= 10 {
			fmt.Printf("%v\n", house.Hand[i])
		}
	}
	fmt.Printf("\nYour score is: %d\n", user.Score)
	fmt.Printf("The House's score is: %d\n", house.Score)
}

func BLACKJACK(deck Deck, user Player, house Player) {
	count := 0
	var choice int
	done := false

	for done == false {
		if count < 4 {
			if count%2 == 0 {
				user.addCard(Deal(deck, &count))
			}
			if count%2 == 1 {
				house.addCard(Deal(deck, &count))
			}
		}

		for count >= 4 && count < len(deck) {
			if done == true {
				printState(user, house)
				whoWins(user, house)
				fmt.Println("=========================")

				fmt.Println("\nNew game?\nPress 1 to Play again\nPress 2 to Quit")
				fmt.Scanln(&choice)
				if choice == 1 {
					done = false
					user.Score = 0
					user.Hand = nil
					house.Score = 0
					house.Hand = nil

					user.addCard(Deal(deck, &count))
					house.addCard(Deal(deck, &count))
					user.addCard(Deal(deck, &count))
					house.addCard(Deal(deck, &count))

				} else {
					break
				}
			}

			printState(user, house)
			fmt.Println("\nPress 1 to Hit\nPress 2 to Stay")
			fmt.Scanln(&choice)

			if choice == 1 {
				user.addCard(Deal(deck, &count))
				if user.Score > 21 {
					done = true
				}
			}
			if choice == 2 {
				for house.Score < 17 {
					house.addCard(Deal(deck, &count))
					time.Sleep(100 * time.Millisecond)
				}
				if house.Score >= 17 {
					//INNOVATIVE
					done = true
				}
			}
		}
	}
}

func main() {
	//Initialize new deck using a function
	deck := newDeck()
	Shuffle(deck)

	//Initialize a struct Player
	user := Player{Name: "User"}
	house := Player{Name: "House"}

	BLACKJACK(deck, user, house)

}
