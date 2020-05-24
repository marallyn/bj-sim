package blackjack

import "strconv"

type ICard interface {
	GetFullName() string
	GetName() string
	GetShortFullName() string
	GetShortName() string
	GetShortSuit() string
	GetValue() int
}

type Card struct {
	id   int
	suit string
}

func (card Card) GetFullName() string {
	return card.GetName() + " of " + card.suit
}

func (card Card) GetShortName() string {
	switch card.id {
	case 1:
		return "A"
	case 11:
		return "J"
	case 12:
		return "Q"
	case 13:
		return "K"
	default:
		return strconv.Itoa(card.id)
	}
}

func (card Card) GetShortFullName() string {
	return card.GetShortName() + card.GetShortSuit()
}

func (card Card) GetShortSuit() string {
	return string(card.suit[0])
}

func (card Card) GetName() string {
	switch card.id {
	case 1:
		return "Ace"
	case 11:
		return "Jack"
	case 12:
		return "Queen"
	case 13:
		return "King"
	default:
		return strconv.Itoa(card.id)
	}
}

func (card Card) GetValue() int {
	switch {
	case card.id >= 10:
		return 10
	default:
		return card.id
	}
}
