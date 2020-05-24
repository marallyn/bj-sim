package blackjack

type Strategy struct {
	name     string
	strategy map[int]map[int]string
}

func (s *Strategy) Init(name string) {
	s.name = name

	switch name {
	case "basic":
		s.strategy = loadBasicStrategy()
	}
}

func loadBasicStrategy() map[int]map[int]string {
	s := map[int]map[int]string{}

	s[1] = map[int]string{
		2:  "hit",
		3:  "hit",
		4:  "hit",
		5:  "hit",
		6:  "hit",
		7:  "hit",
		8:  "hit",
		9:  "hit",
		10: "hit",
		11: "hit",
		12: "hit",
		13: "hit",
		14: "hit",
		15: "hit",
		16: "hit",
		17: "stand",
		18: "stand",
		19: "stand",
		20: "stand",
	}

	return s
}
