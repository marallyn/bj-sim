package main

import "jeff.dog/bj/blackjack"

func main() {
	// make sure we were invoked correctly
	blackjack.VerifyArgs()

	numDecks, numPlayers, numHands, quiet := blackjack.GetArgs()

	game := blackjack.NewGame(numDecks, numPlayers, numHands, quiet)
	// game.Test()
	game.Run()
}
