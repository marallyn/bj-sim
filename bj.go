package main

import "jeff.dog/bj/blackjack"

func main() {
	// make sure we were invoked correctly
	blackjack.VerifyArgs()

	game := blackjack.NewGame(blackjack.GetArgs())
	// game.Test()
	game.Run()
}
