package blackjack

import (
	"fmt"
	"strconv"
)

type Game struct {
	numDecks    int
	numPlayers  int
	numHands    int
	currentHand int
	quiet       bool
	deck        Deck
	dealer      Dealer
	players     []Player
}

func NewGame(decks int, players int, hands int, quiet bool) Game {
	if !quiet {
		fmt.Printf("Starting a %d deck game with %d players for %d hands\n", decks, players, hands)
	}
	game := Game{
		numDecks:   decks,
		numPlayers: players,
		numHands:   hands,
		quiet:      quiet,
	}
	game.deck.Init(decks)

	// Init the dealer
	game.dealer.Init("Dealer")

	// Init the players
	for i := 0; i < game.numPlayers; i++ {
		newPlayer := Player{}
		newPlayer.Init("Player " + strconv.Itoa(i+1))
		game.players = append(game.players, newPlayer)
	}

	return game
}

func (game *Game) Test() {
	// stub for testing new functionality
}

func (game *Game) Run() {
	for hand := 0; hand < game.numHands; hand++ {
		game.currentHand = hand + 1
		game.shuffleIfNecessary()
		game.gameStatus()
		game.playersBet()
		game.deal()
		if game.dealer.CheckForNatural() {
			if !game.quiet {
				fmt.Printf("Dealer has blackjack! %s You suck!\n", game.dealer.GetCardStr())
			}
		} else {
			game.playersAct()
			game.dealerActs()
		}
		game.resolveBets()
		game.resetHands()
		game.playerStatus(false)
	}
	fmt.Println("Final simulation stats:")
	fmt.Println("  Simulation parameters:")
	fmt.Printf("    Number of decks: %d\n", game.numDecks)
	fmt.Printf("    Number of players: %d\n", game.numPlayers)
	fmt.Printf("    Number of hands: %d\n", game.numHands)
	fmt.Println("  Player stats:")
	game.playerStatus(true)
}

func (game *Game) deal() {
	if !game.quiet {
		fmt.Println("Dealing...")
	}
	for i := 0; i < 2; i++ {
		for p := range game.players {
			game.players[p].DealCard(game.deck.DealCard())
		}
		game.dealer.DealCard(game.deck.DealCard())
	}
	if !game.quiet {
		fmt.Printf("Dealer shows %s\n", game.dealer.GetUpCard())
	}
}

func (game *Game) dealerActs() {
	action := "start"
	for action != "stand" {
		if !game.quiet {
			fmt.Printf("Dealer has %s\n", game.dealer.GetCardStr())
		}
		action = game.dealer.GetAction()
		switch {
		case action == "hit":
			if !game.quiet {
				fmt.Println("Dealer takes another card...")
			}
			game.dealer.DealCard(game.deck.DealCard())
		case action == "stand":
			if !game.quiet {
				fmt.Println("Dealer says I'm good here.")
			}
		}
	}
}

func (game *Game) playersAct() {
	dealerUpCardValue := game.dealer.GetUpCardValue()

	for p := range game.players {
		player := &game.players[p]
		action := "start"
		for action != "stand" {
			if !game.quiet {
				fmt.Printf("%s has %s\n", player.GetName(), player.GetCardStr())
			}
			action = player.GetAction(dealerUpCardValue)
			switch {
			case action == "double":
				if !game.quiet {
					fmt.Printf("%s doubles down! What a move!\n", player.GetName())
				}
				player.DealCard(game.deck.DealCard())
				player.DoubleBet()
				action = "stand"
			case action == "hit":
				if !game.quiet {
					fmt.Printf("%s takes another card...\n", player.GetName())
				}
				player.DealCard(game.deck.DealCard())
			case action == "stand":
				if !game.quiet {
					fmt.Printf("%s says I'm good here.\n", player.GetName())
				}
			}
		}
	}
}

func (game *Game) playersBet() {
	for p := range game.players {
		player := &game.players[p]
		player.Bet()
	}
}

func (game *Game) resetHands() {
	for p := range game.players {
		player := &game.players[p]
		player.ResetHand()
	}

	game.dealer.ResetHand()
}

func (game *Game) resolveBets() {
	// softValue is always the best
	_, dealerValue := game.dealer.GetValues()
	dealerHasNatural := game.dealer.CheckForNatural()

	for p := range game.players {
		player := &game.players[p]
		_, playerValue := player.GetValues()
		playerHasNatural := player.CheckForNatural()
		switch {
		case dealerHasNatural && playerHasNatural:
			// player has blackjack, dealer doesn't
			player.Bj()
			player.Push()
			if !game.quiet {
				fmt.Printf(
					"%s gets a bj, and still manages to 'lose' with %s, %d = %d\n",
					player.GetName(),
					player.GetCardStr(),
					playerValue,
					dealerValue,
				)
			}
		case !dealerHasNatural && playerHasNatural:
			// player has blackjack, dealer doesn't
			player.Bj()
			player.Win(1.5)
			if !game.quiet {
				fmt.Printf(
					"%s gets a bj and wins %0.1f with %s, %d > %d\n",
					player.GetName(),
					float64(player.GetBet())*1.5,
					player.GetCardStr(),
					playerValue,
					dealerValue,
				)
			}
		case dealerValue > 21 && playerValue <= 21:
			// dealer busts, player doesn't
			player.Win(1)
			if !game.quiet {
				fmt.Printf(
					"%s wins %d with %s, %d > %d\n",
					player.GetName(),
					player.GetBet(),
					player.GetCardStr(),
					playerValue,
					dealerValue,
				)
			}
		case dealerValue <= 21 && playerValue <= 21 && playerValue > dealerValue:
			// neither bust, playerValue is greater
			player.Win(1)
			if !game.quiet {
				fmt.Printf(
					"%s wins %d with %s, %d > %d\n",
					player.GetName(),
					player.GetBet(),
					player.GetCardStr(),
					playerValue,
					dealerValue,
				)
			}
		case dealerValue <= 21 && playerValue <= 21 && playerValue == dealerValue:
			// player didn't bust, but has same value as dealer
			// push
			player.Push()
			if !game.quiet {
				fmt.Printf(
					"%s, sucks for you %d returned to your chips. %s, %d = %d\n",
					player.GetName(),
					player.GetBet(),
					player.GetCardStr(),
					playerValue,
					dealerValue,
				)
			}
		default:
			// why is the default case a loss? power of positive thinking anyone?
			player.Lose()
			if !game.quiet {
				fmt.Printf(
					"Sonofabitch! %s loses %d with %s, %d < %d\n",
					player.GetName(),
					player.GetBet(),
					player.GetCardStr(),
					playerValue,
					dealerValue,
				)
			}
		}
	}
}

func (game *Game) shuffleIfNecessary() {
	if game.deck.NeedToShuffle(game.numPlayers) {
		if !game.quiet {
			fmt.Println("Shuffling ...\n")
		}
		game.deck.Shuffle()
	}
}

func (game *Game) playerStatus(loud bool) {
	if loud || !game.quiet {
		for _, player := range game.players {
			player.Status(game.currentHand)
		}
	}
}

func (game *Game) gameStatus() {
	if !game.quiet {
		fmt.Printf("Hand %d of %d\n", game.currentHand, game.numHands)
		game.deck.Status()
	}
}
