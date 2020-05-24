package blackjack

import "fmt"

type IPlayer interface {
	IBasePlayer
	Bet()
	DoubleBet()
	GetBet() int
	Lose()
	Push()
	Status(handsPlayed int)
	Win(factor float64)
}

type Player struct {
	BasePlayer
	chips  float64
	bet    int
	pushes int
	losses int
	wins   int
}

// this is inherited from IBasePlayer, but we override, because we have chips to set
func (p *Player) Init(name string) {
	p.BasePlayer.Init(name)
	p.chips = 0
}

func (p *Player) Bet() {
	p.bet = 1
}

func (p *Player) DoubleBet() {
	p.bet *= 2
}

func (p *Player) GetAction(upCard int) string {
	action := ""
	hardValue, softValue := p.GetValues()

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

func (p *Player) GetBet() int {
	return p.bet
}

// this is inherited from IBasePlayer, but we override, because we want to show our bet size
// func (p *Player) HandStatus() {
// 	fmt.Printf(
// 		"%s: bet=%d, cards=%s, value=%s, action=%s\n",
// 		p.name,
// 		p.bet,
// 		p.GetCardStr(),
// 		p.getValueStr(),
// 		p.GetAction(),
// 	)
// }

func (p *Player) ResetHand() {
	p.BasePlayer.ResetHand()
	p.bet = 0
}

func (p *Player) Status(handsPlayed int) {
	winPct := 0.0
	losePct := 0.0
	pushPct := 0.0
	bjPct := 0.0
	if handsPlayed > 0 {
		winPct = float64(p.wins) / float64(handsPlayed) * 100
		losePct = float64(p.losses) / float64(handsPlayed) * 100
		pushPct = float64(p.pushes) / float64(handsPlayed) * 100
		bjPct = float64(p.bjs) / float64(handsPlayed) * 100
	}

	fmt.Printf(
		"    %s: %d chips. W%%: %0.2f L%%: %0.2f P%%: %0.2f BJ%%: %0.2f\n",
		p.name,
		int(p.chips),
		winPct,
		losePct,
		pushPct,
		bjPct,
	)
}

// func (p *Player) Bj() {
// 	p.chips += float64(p.bet)
// 	p.wins += 1
// }

func (p *Player) Lose() {
	p.chips -= float64(p.bet)
	p.losses += 1
}

func (p *Player) Push() {
	p.pushes += 1
}

func (p *Player) Win(factor float64) {
	p.chips += float64(p.bet) * factor
	p.wins += 1
}
