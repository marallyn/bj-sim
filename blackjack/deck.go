package blackjack

import (
	"fmt"
	"math/rand"
	"time"
)

type IDeck interface {
	Init(numDecks int)
	DealCard() Card
	NeedToShuffle(numPlayers int) bool
	Shuffle()
	Status()
}

type Deck struct {
	masterDeck []Card
	available  []Card
	used       []Card
}

func (deck *Deck) Init(numDecks int) {
	// fmt.Printf("Initializing deck...\n")
	suits := [4]string{"Clubs", "Diamonds", "Hearts", "Spades"}

	for d := 0; d < numDecks; d++ {
		for _, suit := range suits {
			for id := 1; id <= 13; id++ {
				deck.masterDeck = append(deck.masterDeck, Card{
					id:   id,
					suit: suit,
				})
			}
		}
	}
}

func (deck *Deck) DealCard() Card {
	card := deck.available[0]

	// remove the card from available
	deck.available = deck.available[1:]

	// add the card to used
	deck.used = append(deck.used, card)

	return card
}

func (deck *Deck) NeedToShuffle(numPlayers int) bool {
	// the average number of cards used is 2.9 per player and dealer, so
	// shuffle if we have less than four per player available
	// bullshit. 71,532 hands into a 100,000 simulation with four players
	// and two decks, I ran out of cards during dealer play
	return len(deck.available) < (numPlayers+1)*5
}

func (deck *Deck) ShowAllDecks() {
	deck.ShowCards()
	deck.ShowAvailable()
	deck.ShowUsed()
}

func (deck *Deck) ShowAvailable() {
	fmt.Printf("There are %d cards remaining in the deck.\n", len(deck.available))
	for _, card := range deck.available {
		fmt.Printf("%s,", card.GetShortFullName())
	}
	fmt.Println("\n\n")
}

func (deck *Deck) ShowCards() {
	fmt.Printf("There are %d cards in the master deck.\n", len(deck.masterDeck))
	for _, card := range deck.masterDeck {
		fmt.Printf("%s,", card.GetShortFullName())
	}
	fmt.Println("\n\n")
}

func (deck *Deck) ShowUsed() {
	fmt.Printf("%d cards have been dealt.\n", len(deck.used))
	for _, card := range deck.used {
		fmt.Printf("%s,", card.GetShortFullName())
	}
	fmt.Println("\n\n")
}

func (deck *Deck) Shuffle() {
	numCards := len(deck.masterDeck)

	// allocate remainingCards, and copy the masterDeck into it
	remainingCards := make([]Card, numCards)
	copy(remainingCards, deck.masterDeck)

	// reset the available and used cards
	deck.available = make([]Card, 0)
	deck.used = make([]Card, 0)

	// using UnixNano, because we execute so fast, Unix was returning the same value
	// each time, and we were getting the same deck order each time
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < numCards; i++ {
		index := rand.Intn(numCards - i)

		// add the card to the available cards
		deck.available = append(deck.available, remainingCards[index])

		// remove the card from the remainingCards by move the last card to replace
		// the just used card, and slicing one less card from remainingCards
		remainingCards[index] = remainingCards[len(remainingCards)-1]
		remainingCards = remainingCards[0 : len(remainingCards)-1]
	}
}

func (deck *Deck) Status() {
	fmt.Printf("%d of %d cards remain. ", len(deck.available), len(deck.masterDeck))
	fmt.Printf("%d cards have been dealt.\n", len(deck.used))
}
