package blackjack

import (
	"fmt"
	"os"
	"strconv"
)

func GetArgs() (int, int, int, bool) {
	numDecks, _ := strconv.Atoi(os.Args[1])
	numPlayers, _ := strconv.Atoi(os.Args[2])
	numHands, _ := strconv.Atoi(os.Args[3])

	// the quiet arg ca be anything, it just has to exist
	quiet := len(os.Args) > 4

	return numDecks, numPlayers, numHands, quiet
}

func showHelp() {
	fmt.Println("\nbj, it's not what you think, it's a blackjack simulator")
	fmt.Println("    Usage:")
	fmt.Println("        bj <# of decks> <# of players> <# of hands> [quiet]\n")
}

func VerifyArgs() {
	// if we have less than the three required arguments, get out
	if len(os.Args) < 4 {
		showHelp()
		os.Exit(1)
	}

	d, p, h, _ := GetArgs()

	if d <= 0 || p <= 0 || h <= 0 {
		showHelp()
		os.Exit(1)
	}
}
