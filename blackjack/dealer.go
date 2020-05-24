package blackjack

type IDealer interface {
	GetUpCard() string
	GetUpCardValue() int
}

type Dealer struct {
	BasePlayer
}

func (dealer *Dealer) GetUpCard() string {
	return dealer.cards[0].GetShortFullName()
}

// func (dealer *Dealer) GetUpCardValue() int {
// 	return dealer.cards[0].GetValue()
// }

func (dealer *Dealer) GetUpCardValue() int {
	return dealer.cards[0].GetValue()
}

func (dealer *Dealer) GetAction() string {
	action := ""

	hardValue, softValue := dealer.GetValues()

	switch {
	case softValue <= 16:
		action = "hit"
	case softValue == 17 && hardValue < softValue:
		// dealer hits a soft 17
		action = "hit"
	default:
		action = "stand"
	}

	return action
}
