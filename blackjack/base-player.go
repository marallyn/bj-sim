package blackjack

import (
	"strconv"
	"strings"
)

type IBasePlayer interface {
	Init(name string)
	Bj()
	CheckForNatural() bool
	DealCard(card Card)
	// GetAction() string
	GetCardStr() string
	GetName() string
	GetValues() (int, int)
	getValueStr() string
	// HandStatus()
	ResetHand()
}

type BasePlayer struct {
	name  string
	bjs   int
	cards []Card
	// strategy Strategy
}

func (bp *BasePlayer) Init(name string) {
	bp.name = name
}

func (bp *BasePlayer) Bj() {
	bp.bjs += 1
}

func (bp *BasePlayer) CheckForNatural() bool {
	// we only check for naturals when the bp has two cards
	if len(bp.cards) != 2 {
		return false
	}

	// since we have two cards, a natural exists when the softValue is 21
	_, softValue := bp.GetValues()

	return softValue == 21
}

func (bp *BasePlayer) DealCard(card Card) {
	bp.cards = append(bp.cards, card)
	// fmt.Printf("%s recieves %s\n", bp.name, card.GetShortFullName())
}

// func (bp *BasePlayer) GetAction() string {
// 	return "stand"
// }

func (bp *BasePlayer) GetCardStr() string {
	cardStr := ""

	for _, card := range bp.cards {
		cardStr += card.GetShortFullName() + " "
	}

	return strings.TrimRight(cardStr, " ")
}

func (bp *BasePlayer) GetName() string {
	return bp.name
}

func (bp *BasePlayer) GetValues() (int, int) {
	hard, soft := 0, 0
	hasAce := false

	for _, card := range bp.cards {
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

func (bp *BasePlayer) getValueStr() string {
	hardValue, softValue := bp.GetValues()

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

// func (bp *BasePlayer) HandStatus() {
// 	fmt.Printf(
// 		"%s: cards=%s, value=%s, action=%s\n",
// 		bp.name,
// 		bp.GetCardStr(),
// 		bp.getValueStr(),
// 		// bp.GetAction(),
// 	)
// }

func (bp *BasePlayer) ResetHand() {
	bp.cards = []Card{}
}
