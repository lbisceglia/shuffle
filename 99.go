package main

import (
	"shuffle/cards"
)

// Main initializes a single round of 99 with 4 players.
// The current implementation is a simple proof of concept.
// Players are currently all controlled by the same person (you) via the command line.
// All cards are temporarily disclosed for debugging purposes.
func main() {
	names := []string{"Alice", "Bob", "Charlie", "Dan"}
	players := initializePlayers(names)
	mgr := new(cards.NNGameManager)
	mgr.NewGame(players)
}

// InitializePlayers is a factory that creates players with the given names.
func initializePlayers(names []string) (players []*cards.NNPlayer) {
	for _, name := range names {
		players = append(players, cards.NewNNPlayer(name))
	}
	return players
}
