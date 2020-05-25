package blackjack

import "fmt"

type IPlayer interface {
	IBasePlayer
	Bet()
	DoubleBet()
	GetBet() int
	Lose()
	Push()
	Status()
	Win(factor float64)
}

type Player struct {
	BasePlayer
	chips  float64
	pushes int
	losses int
	wins   int
	hands  []Hand
}

// this is inherited from IBasePlayer, but we override, because we have chips to set
func (p *Player) Init(name string) {
	p.BasePlayer.Init(name)
}

func (p *Player) Bet() {
	// this is called from game.playersBet before dealing, so the player only has one hand
	p.hands[0].Bet()
}

func (p *Player) GetAction(hand Hand, upCard int) string {
	action := ""
	hardValue, softValue := hand.GetValues()

	if upCard == 1 {
		upCard = 11
	}

	if hardValue == softValue {
		action = p.getHardAction(upCard, hardValue)
	} else {
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
	case hardValue == 9 && (upCard >= 3 && upCard <= 5):
		action = "double"
	case hardValue == 10 && upCard > 9:
		action = "hit"
	case hardValue == 10 && upCard <= 9:
		action = "double"
	case hardValue == 11 && upCard > 10:
		action = "hit"
	case hardValue == 11 && upCard <= 10:
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

func (p *Player) getSoftAction(upCard int, softValue int) string {
	action := ""

	switch {
	case softValue == 12:
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
