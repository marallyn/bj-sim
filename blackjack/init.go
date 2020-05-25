package blackjack

import (
	"fmt"
	"os"
	"strconv"
)

func GetQuiet() bool {
	return len(os.Args) > 4
}

func GetArgs() (int, int, int) {
	numDecks, _ := strconv.Atoi(os.Args[1])
	numPlayers, _ := strconv.Atoi(os.Args[2])
	numHands, _ := strconv.Atoi(os.Args[3])

	return numDecks, numPlayers, numHands
}

func showHelp() {
	fmt.Println("\nbj, it's not what you think, it's a blackjack simulator")
	fmt.Println("    Usage:")
	fmt.Println("        bj <# of decks> <# of players> <# of hands> [quiet]\n")
}

func VerifyArgs() {
	if len(os.Args) < 4 {
		showHelp()
		os.Exit(1)
	}

	d, p, h := GetArgs()

	if d <= 0 || p <= 0 || h <= 0 {
		showHelp()
		os.Exit(1)
	}
}
