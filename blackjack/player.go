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
	Strategy
}

// this is inherited from IBasePlayer, but we override, because we have chips to set
func (p *Player) Init(name string, chips float64, strategyName string) {
	p.name = name
	p.chips = chips
	p.Strategy.Init(strategyName)
}

func (p *Player) Bet() {
	// this is called from game.playersBet before dealing, so the player only has one hand
	p.hands[0].Bet()
}

func (p *Player) Bj() {
	p.bjs += 1
}

func (p *Player) GetHand(index int) *Hand {
	return &p.hands[index]
}

func (p *Player) GetBet(hand Hand) int {
	return hand.GetBet()
}

func (p *Player) GetName() string {
	return p.name
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
		"    %10s: %6d chips %d hands W%%: %5.2f L%%: %5.2f P%%: %5.2f BJ%%: %5.2f\n",
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
