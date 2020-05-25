package blackjack

type IBasePlayer interface {
	Init(name string)
	Bj()
	GetName() string
}

type BasePlayer struct {
	name        string
	bjs         int
	handsPlayed int
	// strategy Strategy
}

func (bp *BasePlayer) Init(name string) {
	bp.name = name
}

func (bp *BasePlayer) Bj() {
	bp.bjs += 1
}

func (bp *BasePlayer) GetName() string {
	return bp.name
}
