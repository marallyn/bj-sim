package blackjack

type IPlayer interface {
	Init(name string, chips float64, strategyName string)
	Bet()
	GetBet(hand Hand) int
	GetHand(index int) *Hand
	GetName() string
	ResetHands()
	Split(hand *Hand, handIndex int) *Hand
}

type Player struct {
	name        string
	hands       []Hand
	Stats
	Strategy
}

// this is inherited from IBasePlayer, but we override, because we have chips to set
func (p *Player) Init(name string, chips float64, strategyName string) {
	p.name = name
	p.Stats.Init(chips)
	p.Strategy.Init(strategyName)
}

func (p *Player) Bet() {
	// this is called from game.playersBet before dealing, so the player only has one hand
	p.hands[0].Bet()
}

func (p *Player) GetBet(hand Hand) int {
	return hand.GetBet()
}

func (p *Player) GetHand(index int) *Hand {
	return &p.hands[index]
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) ResetHands() {
	// start fresh with one hand
	p.hands = make([]Hand, 1)
}

func (p *Player) ShowStats() Stats {
	p.Stats.ShowStats(p.name, p.Strategy.GetName())

	return p.Stats
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
