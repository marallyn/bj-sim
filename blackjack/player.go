package blackjack

import "fmt"

type IPlayer interface {
	Bet()
	DoubleBet()
	GetBet() int
	Lose()
	Push()
	Status()
	Win(factor float64)
}

type Player struct {
	name        string
	chips       float64
	handsPlayed int
	wins        int
	losses      int
	pushes      int
	bjs         int
	hands       []Hand
	// strategy Strategy
}

// this is inherited from IBasePlayer, but we override, because we have chips to set
func (p *Player) Init(name string, chips float64) {
	p.name = name
	p.chips = chips
}

func (p *Player) Bet() {
	// this is called from game.playersBet before dealing, so the player only has one hand
	p.hands[0].Bet()
}

func (p *Player) Bj() {
	p.bjs += 1
}

func (p *Player) GetAction(hand Hand, upCard int) string {
	action := ""
	hardValue, softValue := hand.GetValues()

	if upCard == 1 {
		upCard = 11
	}

	switch {
	case len(hand.cards) == 2 && hand.cards[0].GetShortName() == hand.cards[1].GetShortName():
		// there are two cards, and their short names are the same, so it's a pair
		action = p.getPairAction(upCard, hand.cards[0].GetShortName())
	case hardValue == softValue:
		action = p.getHardAction(upCard, hardValue)
	default:
		action = p.getSoftAction(upCard, softValue)
	}

	return action
}

func (p *Player) GetHand(index int) *Hand {
	return &p.hands[index]
}

func (p *Player) getHardAction(upCard int, hardValue int) string {
	action := ""

	switch {
	case hardValue < 9:
		action = "hit"
	case hardValue == 9 && (upCard < 3 || upCard > 5):
		action = "hit"
	case hardValue == 9:
		action = "double"
	case hardValue == 10 && upCard > 9:
		action = "hit"
	case hardValue == 10:
		action = "double"
	case hardValue == 11 && upCard > 10:
		action = "hit"
	case hardValue == 11:
		action = "double"
	case hardValue == 12 && (upCard < 4 || upCard > 6):
		action = "hit"
	case (hardValue >= 13 && hardValue <= 16) && upCard > 6:
		action = "hit"
	default:
		action = "stand"
	}

	return action
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) getPairAction(upCard int, cardName string) string {
	action := ""

	switch {
	case cardName == "A":
		action = "split"
	case (cardName == "2" || cardName == "3") && upCard < 8:
		action = "split"
	case cardName == "2" || cardName == "3":
		action = "hit"
	case cardName == "3" && (upCard >= 3 && upCard <= 5):
		action = "double"
	case cardName == "4" && (upCard == 5 || upCard == 6):
		action = "split"
	case cardName == "4":
		action = "hit"
	case cardName == "5" && upCard <= 9:
		action = "double"
	case cardName == "5":
		action = "hit"
	case cardName == "6" && upCard < 7:
		action = "split"
	case cardName == "6":
		action = "hit"
	case cardName == "7" && upCard < 8:
		action = "split"
	case cardName == "7":
		action = "hit"
	case cardName == "8":
		action = "split"
	case cardName == "9" && (upCard < 7 || upCard == 8 || upCard == 9):
		action = "split"
	default:
		action = "stand"
	}

	return action
}

func (p *Player) getSoftAction(upCard int, softValue int) string {
	action := ""

	switch {
	case softValue <= 12:
		action = "hit"
	case (softValue == 13 || softValue == 14) && (upCard == 5 || upCard == 6):
		action = "double"
	case (softValue == 13 || softValue == 14):
		action = "hit"
	case (softValue == 15 || softValue == 16) && (upCard > 3 && upCard < 7):
		action = "double"
	case (softValue == 15 || softValue == 16):
		action = "hit"
	case (softValue == 17 || softValue == 18) && (upCard > 2 && upCard < 7):
		action = "double"
	case softValue == 17:
		action = "hit"
	case softValue == 18 && upCard > 8:
		action = "hit"
	default:
		action = "stand"
	}

	return action
}

func (p *Player) GetBet(hand Hand) int {
	return hand.GetBet()
}

func (p *Player) ResetHands() {
	// start fresh with one hand
	p.hands = make([]Hand, 1)
}

func (p *Player) Split(hand *Hand, handIndex int) *Hand {
	newHand := Hand{
		cards: make([]Card, 0),
		bet:   hand.GetBet(),
	}

	// move the second card from the first hand to the newHand
	newHand.cards = append(newHand.cards, hand.cards[1])

	// remove the second card from the first hand
	hand.cards = hand.cards[0:1]

	// add the newHand to the player's hand
	p.hands = append(p.hands, newHand)

	return &p.hands[handIndex]
}

func (p *Player) Status() {
	winPct := 0.0
	losePct := 0.0
	pushPct := 0.0
	bjPct := 0.0
	if p.handsPlayed > 0 {
		winPct = float64(p.wins) / float64(p.handsPlayed) * 100
		losePct = float64(p.losses) / float64(p.handsPlayed) * 100
		pushPct = float64(p.pushes) / float64(p.handsPlayed) * 100
		bjPct = float64(p.bjs) / float64(p.handsPlayed) * 100
	}

	fmt.Printf(
		"    %s: %d chips %d hands W%%: %0.2f L%%: %0.2f P%%: %0.2f BJ%%: %0.2f\n",
		p.name,
		int(p.chips),
		p.handsPlayed,
		winPct,
		losePct,
		pushPct,
		bjPct,
	)
}

func (p *Player) Lose(hand Hand) {
	p.chips -= float64(hand.GetBet())
	p.losses += 1
	p.handsPlayed += 1
}

func (p *Player) Push() {
	p.pushes += 1
	p.handsPlayed += 1
}

func (p *Player) Win(hand Hand, factor float64) {
	p.chips += float64(hand.GetBet()) * factor
	p.wins += 1
	p.handsPlayed += 1
}
