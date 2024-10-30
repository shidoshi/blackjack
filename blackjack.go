package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Card represents a playing card
type Card struct {
	Rank string
	Suit string
}

// Deck represents a deck of cards
type Deck []Card

// NewDeck creates a new deck of cards
func NewDeck() Deck {
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	deck := make(Deck, 0, len(suits)*len(ranks))
	for _, suit := range suits {
		for _, rank := range ranks {
			deck = append(deck, Card{Rank: rank, Suit: suit})
		}
	}
	return deck
}

// Shuffle shuffles the deck
func (d Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d), func(i, j int) { d[i], d[j] = d[j], d[i] })
}

// Deal deals a card from the deck
func (d *Deck) Deal() Card {
	card := (*d)[0]
	*d = (*d)[1:]
	return card
}

// Score calculates the score of a hand
func Score(hand []Card) int {
	score := 0
	hasAce := false
	for _, card := range hand {
		switch card.Rank {
		case "A":
			hasAce = true
			score += 11
		case "J", "Q", "K":
			score += 10
		default:
			score += int([]rune(card.Rank)[0] - '0')
		}
	}
	if hasAce && score > 21 {
		score -= 10
	}
	return score
}

func main() {
	deck := NewDeck()
	deck.Shuffle()

	// Deal cards to player and dealer
	playerHand := []Card{deck.Deal(), deck.Deal()}
	dealerHand := []Card{deck.Deal(), deck.Deal()}

	// Print initial hands
	fmt.Println("Your hand:", playerHand)
	fmt.Println("Dealer's hand:", dealerHand[0], "and an unknown card")

	// Player's turn
	for Score(playerHand) < 21 {
		fmt.Println("Your score:", Score(playerHand))
		fmt.Print("Hit or stand? ")
		var action string
		fmt.Scanln(&action)
		if action == "stand" {
			break
		}
		playerHand = append(playerHand, deck.Deal())
		fmt.Println("Your hand:", playerHand)
	}

	playerScore := Score(playerHand)
	fmt.Println("Your final score:", playerScore)

	// Dealer's turn
	for Score(dealerHand) < 17 {
		dealerHand = append(dealerHand, deck.Deal())
	}
	dealerScore := Score(dealerHand)
	fmt.Println("Dealer's hand:", dealerHand)
	fmt.Println("Dealer's final score:", dealerScore)

	// Determine the winner
	if playerScore > 21 {
		fmt.Println("You bust! Dealer wins.")
	} else if dealerScore > 21 {
		fmt.Println("Dealer busts! You win.")
	} else if playerScore > dealerScore {
		fmt.Println("You win!")
	} else if playerScore < dealerScore {
		fmt.Println("Dealer wins.")
	} else {
		fmt.Println("It's a push!")
	}
}
