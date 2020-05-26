package blackjack

import "fmt"

type IStats interface {
	Init(chips float64)
	Bj()
	Lose()
	Push()
	ShowStats()
	Win(factor float64)
}

type Stats struct {
	chips       float64
	handsPlayed int
	wins        int
	losses      int
	pushes      int
	bjs         int
}

func (s *Stats) Init(chips float64) {
	s.chips = chips
}

func (s *Stats) Bj() {
	s.bjs += 1
}

func (s *Stats) Lose(hand Hand) {
	s.chips -= float64(hand.GetBet())
	s.losses += 1
	s.handsPlayed += 1
}

func (s *Stats) Push() {
	s.pushes += 1
	s.handsPlayed += 1
}

func (s *Stats) ShowStats(name string, strategyName string) {
	winPct := 0.0
	losePct := 0.0
	pushPct := 0.0
	bjPct := 0.0
	if s.handsPlayed > 0 {
		winPct = float64(s.wins) / float64(s.handsPlayed) * 100
		losePct = float64(s.losses) / float64(s.handsPlayed) * 100
		pushPct = float64(s.pushes) / float64(s.handsPlayed) * 100
		bjPct = float64(s.bjs) / float64(s.handsPlayed) * 100
	}

	playerAndStrategyNames := name + " (" + strategyName + ")"
	fmt.Printf(
		"    %20s: %6d chips %d hands W%%: %5.2f L%%: %5.2f P%%: %5.2f BJ%%: %5.2f\n",
		playerAndStrategyNames,
		int(s.chips),
		s.handsPlayed,
		winPct,
		losePct,
		pushPct,
		bjPct,
	)
}

func (s *Stats) Win(hand Hand, factor float64) {
	s.chips += float64(hand.GetBet()) * factor
	s.wins += 1
	s.handsPlayed += 1
}
