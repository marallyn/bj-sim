package blackjack

import (
	"strconv"
	"strings"
)

type IHand interface {
	Bet()
	CheckForNatural() bool
	DoubleBet()
	GetBet() int
	GetCardStr() string
	GetValues() (int, int)
	GetValueStr() string
	ReceiveCard(card Card)
	ResetHand()
}

type Hand struct {
	cards []Card
	bet   int
}

func (hand *Hand) Bet() {
	hand.bet = 1
}

func (hand *Hand) CheckForNatural() bool {
	// we only check for naturals when the hand has two cards
	if len(hand.cards) != 2 {
		return false
	}

	// since we have two cards, a natural exists when the softValue is 21
	_, softValue := hand.GetValues()

	return softValue == 21
}

func (hand *Hand) DoubleBet() {
	hand.bet *= 2
}

func (hand *Hand) GetBet() int {
	return hand.bet
}

func (hand *Hand) GetCardStr() string {
	cardStr := ""

	for _, card := range hand.cards {
		cardStr += card.GetShortFullName() + " "
	}

	return strings.TrimRight(cardStr, " ")
}

func (hand *Hand) GetValues() (int, int) {
	hard, soft := 0, 0
	hasAce := false

	for _, card := range hand.cards {
		val := card.GetValue()
		if val == 1 {
			hasAce = true
		}
		hard += val

	}

	if hasAce && hard <= 11 {
		// adding 10 to a soft value is optional, so only do it if we don't bust
		soft = hard + 10
	} else {
		soft = hard
	}

	return hard, soft
}

func (hand *Hand) GetValueStr() string {
	hardValue, softValue := hand.GetValues()

	switch {
	case hardValue > 21:
		return "Busted!"
	case hardValue == 21 || softValue == 21:
		return "21!"
	case hardValue == softValue || softValue > 21:
		return strconv.Itoa(hardValue)
	default:
		return strconv.Itoa(hardValue) + " or " + strconv.Itoa(softValue)
	}

}

func (hand *Hand) ReceiveCard(card Card) {
	hand.cards = append(hand.cards, card)
}

func (hand *Hand) ResetHand() {
	hand.cards = make([]Card, 0)
}
