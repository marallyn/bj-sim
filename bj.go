package main

import "jeff.dog/bj/blackjack"

func main() {
	// make sure we were invoked correctly
	blackjack.VerifyArgs()

	numDecks, numPlayers, numHands := blackjack.GetArgs()
	quiet := blackjack.GetQuiet()

	game := blackjack.NewGame(numDecks, numPlayers, numHands, quiet)
	// game.Test()
	game.Run()
}
