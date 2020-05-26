package blackjack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)


type Strategy struct {
	Name     string  `json:"name"`
	Hard map[int][]string  `json:"hard"`
	Soft map[int][]string  `json:"soft"`
	Pair map[int][]string  `json:"pair"`
}

func (s *Strategy) Init(name string) {
	s.readStrategyFile(name)
}

func (s *Strategy)readStrategyFile(name string) {
	fileName := "./strategies/st-" + name + ".json"
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Printf("Error loading strategy file '%s'.\n", fileName)
		fmt.Println("Switching to hard-coded basic strategy.")
		return
	}

	err = json.Unmarshal(data, s)

	if err != nil {
		fmt.Printf("The strategy file '%s' has a problem in it.\n", fileName)
		fmt.Println("Switching to hard-coded basic strategy.")
		fmt.Println("Just kidding, not implemented yet, so find a way to fix your strategy.")
	}
}

func (s *Strategy) GetAction(hand Hand, upCard int) string {
	action := ""
	hardValue, softValue := hand.GetValues()
	dealerCardIndex := upCard - 2
	if dealerCardIndex < 0 {
		dealerCardIndex = 9
	}

	switch {
	// handle some special cases, so we don't break the strategy tables
	case len(hand.cards) < 2:
		// this happens when we split a hand, and only have one card
		action = "h"
	case softValue >= 21:
		// don't need anymore cards if we are already busted or have 21
		action = "s"
	case len(hand.cards) == 2 && hand.cards[0].GetShortName() == hand.cards[1].GetShortName():
		// there are two cards, and their short names are the same, so it's a pair
		action = s.Pair[hand.cards[0].GetValue()][dealerCardIndex]
	case hardValue == softValue:
		action = s.Hard[hardValue][dealerCardIndex]
	default:
		action = s.Soft[softValue-11][dealerCardIndex]
	}

	// strategy files use abbreviations for actions, so translate to action name
	switch action {
	case "d":
		return "double"
	case "h":
		return "hit"
	case "s":
		return "stand"
	case "v":
		return "split"
	default:
		// just in case someone screws up the stratgey file, stand
		return "stand"
	}
}

func (s *Strategy) GetName() string {
	return s.Name
}

