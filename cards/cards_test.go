package cards

import (
	"shuffle/utils"
	"testing"
)

const (
	suitsPerDeck = 4
	ranksPerDeck = 13
	cardsPerDeck = 52
)

type CardStringResult struct {
	card         Card
	expectedRank string
	expectedSuit string
	expectedCard string
}

func TestCardString(t *testing.T) {
	var tests = map[string]CardStringResult{
		"value card; club":      {NewCard(Ten, Clubs), "10", "♣", "10♣"},
		"jack; diamond":         {NewCard(Jack, Diamonds), "J", "♦", "J♦"},
		"queen; heart":          {NewCard(Queen, Hearts), "Q", "♥", "Q♥"},
		"king; spade":           {NewCard(King, Spades), "K", "♠", "K♠"},
		"ace; invalid suit":     {NewCard(Ace, "elephants"), "A", "invalid suit", "A, invalid suit"},
		"invalid rank":          {NewCard("14", Spades), "invalid rank", "♠", "invalid rank, ♠"},
		"invalid rank and suit": {NewCard("14", "elephants"), "invalid rank", "invalid suit", "invalid rank, invalid suit"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			utils.Error(t, test.card.rank.String(), test.expectedRank)
			utils.Error(t, test.card.suit.String(), test.expectedSuit)
			utils.Error(t, test.card.String(), test.expectedCard)
		})
	}
}

type ShoeResult struct {
	decks          int
	expectedLength int
	expectedSuits  int
	expectedRanks  int
}

func TestNewShoe(t *testing.T) {
	newShoeTests := map[string]ShoeResult{
		"neg":           {-1, 0, 0, 0},
		"zero":          {0, 0, 0, 0},
		"one deck":      {1, cardsPerDeck, suitsPerDeck, ranksPerDeck},
		"six card shoe": {6, 6 * cardsPerDeck, suitsPerDeck, ranksPerDeck},
	}

	for name, test := range newShoeTests {
		t.Run(name, func(t *testing.T) {
			testShoe(t, NewShoe(test.decks), test)
		})
	}
}

func testShoe(t *testing.T, s Shoe, test ShoeResult) {
	suits := make(map[Suit]int)
	ranks := make(map[Rank]int)

	counter := 0
	for _, c := range s {
		suits[c.suit]++
		ranks[c.rank]++
		counter++
	}

	utils.Error(t, counter, test.expectedLength, "cards in the shoe")
	utils.Error(t, len(suits), test.expectedSuits, "suits in the shoe")
	utils.Error(t, len(ranks), test.expectedRanks, "ranks in the shoe")

	cardsPerSuit := test.decks * test.expectedRanks
	for suit, got := range suits {
		utils.Error(t, got, cardsPerSuit, suit.String())
	}

	cardsPerRank := test.decks * test.expectedSuits
	for rank, got := range ranks {
		utils.Error(t, got, cardsPerRank, rank.String())
	}
}
