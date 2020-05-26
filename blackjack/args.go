package blackjack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// well, this took too long to figure out
// the fields in the struct need to be exported to be filled in by Unmarshal
// rather than change Player to export all fields, I created ArgPlayer to read
// Name, Chips (and soon to be Strategy) from a setup file
type Args struct {
	NumDecks   int  `json:"numDecks"`
	NumPlayers int  `json:"numPlayers"`
	NumHands   int  `json:"numHands"`
	Quiet      bool `json:"quiet"`
	Players    []ArgPlayer
}

type ArgPlayer struct {
	Name  string  `json:"name"`
	Chips float64 `json:"chips"`
	Strategy string `json:"strategy"`
}

var setupObj Args

func GetArgs() Args {
	return setupObj
}

func readCommandLineArgs() {
	setupObj.NumDecks, _ = strconv.Atoi(os.Args[1])
	setupObj.NumPlayers, _ = strconv.Atoi(os.Args[2])
	setupObj.NumHands, _ = strconv.Atoi(os.Args[3])

	// the quiet arg ca be anything, it just has to exist
	setupObj.Quiet = len(os.Args) > 4
}

func readSetupFile() {
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Either '%s' does not exist, or I can't read it.\n", os.Args[1])
		os.Exit(1)
	}

	err = json.Unmarshal(data, &setupObj)

	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}

func showHelp() {
	fmt.Println("\nbj, it's not what you think, it's a blackjack simulator")
	fmt.Println("    Usage:")
	fmt.Println("        bj <# of decks> <# of players> <# of hands> [quiet]\n")
	fmt.Println("    or:")
	fmt.Println("        bj <setup file>\n")
}

func VerifyArgs() {
	// if we have less than the three required arguments, get out

	switch {
	case len(os.Args) < 2:
		showHelp()
		os.Exit(1)
	case len(os.Args) == 2:
		readSetupFile()
	case len(os.Args) > 3:
		readCommandLineArgs()
	}

	if setupObj.NumDecks <= 0 || setupObj.NumPlayers <= 0 || setupObj.NumHands <= 0 {
		showHelp()
		os.Exit(1)
	}
}
