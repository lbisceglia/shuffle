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

// TODO: refactor to read decks in from file.
var standardDeck = Shoe{
	NewCard(Ace, Clubs),
	NewCard(Two, Clubs),
	NewCard(Three, Clubs),
	NewCard(Four, Clubs),
	NewCard(Five, Clubs),
	NewCard(Six, Clubs),
	NewCard(Seven, Clubs),
	NewCard(Eight, Clubs),
	NewCard(Nine, Clubs),
	NewCard(Ten, Clubs),
	NewCard(Jack, Clubs),
	NewCard(Queen, Clubs),
	NewCard(King, Clubs),
	NewCard(Ace, Diamonds),
	NewCard(Two, Diamonds),
	NewCard(Three, Diamonds),
	NewCard(Four, Diamonds),
	NewCard(Five, Diamonds),
	NewCard(Six, Diamonds),
	NewCard(Seven, Diamonds),
	NewCard(Eight, Diamonds),
	NewCard(Nine, Diamonds),
	NewCard(Ten, Diamonds),
	NewCard(Jack, Diamonds),
	NewCard(Queen, Diamonds),
	NewCard(King, Diamonds),
	NewCard(Ace, Hearts),
	NewCard(Two, Hearts),
	NewCard(Three, Hearts),
	NewCard(Four, Hearts),
	NewCard(Five, Hearts),
	NewCard(Six, Hearts),
	NewCard(Seven, Hearts),
	NewCard(Eight, Hearts),
	NewCard(Nine, Hearts),
	NewCard(Ten, Hearts),
	NewCard(Jack, Hearts),
	NewCard(Queen, Hearts),
	NewCard(King, Hearts),
	NewCard(Ace, Spades),
	NewCard(Two, Spades),
	NewCard(Three, Spades),
	NewCard(Four, Spades),
	NewCard(Five, Spades),
	NewCard(Six, Spades),
	NewCard(Seven, Spades),
	NewCard(Eight, Spades),
	NewCard(Nine, Spades),
	NewCard(Ten, Spades),
	NewCard(Jack, Spades),
	NewCard(Queen, Spades),
	NewCard(King, Spades),
}

var shuffled2021Deck = Shoe{
	NewCard(Ten, Clubs),
	NewCard(Nine, Diamonds),
	NewCard(Four, Diamonds),
	NewCard(Six, Clubs),
	NewCard(King, Clubs),
	NewCard(King, Spades),
	NewCard(Five, Hearts),
	NewCard(Queen, Clubs),
	NewCard(Jack, Diamonds),
	NewCard(Two, Clubs),
	NewCard(Six, Hearts),
	NewCard(Ten, Diamonds),
	NewCard(Ace, Diamonds),
	NewCard(Three, Hearts),
	NewCard(Two, Diamonds),
	NewCard(Jack, Hearts),
	NewCard(Eight, Clubs),
	NewCard(Jack, Spades),
	NewCard(Queen, Spades),
	NewCard(Ten, Spades),
	NewCard(Two, Hearts),
	NewCard(Ace, Hearts),
	NewCard(Three, Clubs),
	NewCard(Four, Hearts),
	NewCard(Seven, Clubs),
	NewCard(Ten, Hearts),
	NewCard(Nine, Clubs),
	NewCard(Four, Spades),
	NewCard(Queen, Hearts),
	NewCard(Seven, Diamonds),
	NewCard(Ace, Clubs),
	NewCard(Eight, Hearts),
	NewCard(Eight, Spades),
	NewCard(Six, Spades),
	NewCard(Five, Spades),
	NewCard(King, Hearts),
	NewCard(Ace, Spades),
	NewCard(Five, Clubs),
	NewCard(Nine, Spades),
	NewCard(Nine, Hearts),
	NewCard(Jack, Clubs),
	NewCard(Seven, Hearts),
	NewCard(Three, Spades),
	NewCard(King, Diamonds),
	NewCard(Two, Spades),
	NewCard(Six, Diamonds),
	NewCard(Eight, Diamonds),
	NewCard(Three, Diamonds),
	NewCard(Five, Diamonds),
	NewCard(Seven, Spades),
	NewCard(Queen, Diamonds),
	NewCard(Four, Clubs),
}

var shuffled2021Shoe = Shoe{
	NewCard(Five, Diamonds),
	NewCard(Two, Spades),
	NewCard(Nine, Clubs),
	NewCard(King, Spades),
	NewCard(Three, Diamonds),
	NewCard(Seven, Clubs),
	NewCard(Six, Hearts),
	NewCard(Six, Clubs),
	NewCard(Three, Clubs),
	NewCard(Four, Clubs),
	NewCard(Ten, Diamonds),
	NewCard(Eight, Spades),
	NewCard(Jack, Clubs),
	NewCard(Three, Hearts),
	NewCard(Ace, Diamonds),
	NewCard(Four, Diamonds),
	NewCard(Four, Hearts),
	NewCard(Two, Diamonds),
	NewCard(Jack, Clubs),
	NewCard(Nine, Spades),
	NewCard(Six, Hearts),
	NewCard(Two, Clubs),
	NewCard(Five, Hearts),
	NewCard(Ace, Spades),
	NewCard(Three, Spades),
	NewCard(Five, Clubs),
	NewCard(Six, Spades),
	NewCard(Eight, Spades),
	NewCard(Eight, Clubs),
	NewCard(Six, Diamonds),
	NewCard(Ten, Spades),
	NewCard(Four, Spades),
	NewCard(Four, Diamonds),
	NewCard(Nine, Clubs),
	NewCard(King, Hearts),
	NewCard(Ten, Spades),
	NewCard(Six, Spades),
	NewCard(King, Spades),
	NewCard(Five, Spades),
	NewCard(Five, Diamonds),
	NewCard(Eight, Clubs),
	NewCard(Jack, Spades),
	NewCard(Eight, Hearts),
	NewCard(Ace, Hearts),
	NewCard(Queen, Spades),
	NewCard(Nine, Spades),
	NewCard(Five, Clubs),
	NewCard(Three, Hearts),
	NewCard(Queen, Diamonds),
	NewCard(Nine, Hearts),
	NewCard(Five, Hearts),
	NewCard(Seven, Diamonds),
	NewCard(Eight, Diamonds),
	NewCard(Eight, Hearts),
	NewCard(Two, Spades),
	NewCard(King, Diamonds),
	NewCard(Jack, Hearts),
	NewCard(Ten, Hearts),
	NewCard(Four, Clubs),
	NewCard(Nine, Hearts),
	NewCard(Five, Spades),
	NewCard(Queen, Clubs),
	NewCard(Jack, Diamonds),
	NewCard(Two, Diamonds),
	NewCard(Nine, Diamonds),
	NewCard(Jack, Spades),
	NewCard(Ace, Diamonds),
	NewCard(Three, Diamonds),
	NewCard(Nine, Diamonds),
	NewCard(Ace, Hearts),
	NewCard(Ten, Diamonds),
	NewCard(Queen, Clubs),
	NewCard(Six, Diamonds),
	NewCard(Four, Hearts),
	NewCard(Queen, Hearts),
	NewCard(Ten, Clubs),
	NewCard(Seven, Diamonds),
	NewCard(Seven, Spades),
	NewCard(Queen, Diamonds),
	NewCard(Six, Clubs),
	NewCard(King, Clubs),
	NewCard(Ace, Clubs),
	NewCard(Ace, Clubs),
	NewCard(Two, Hearts),
	NewCard(Two, Clubs),
	NewCard(King, Diamonds),
	NewCard(King, Clubs),
	NewCard(Seven, Spades),
	NewCard(Seven, Hearts),
	NewCard(Ten, Clubs),
	NewCard(Jack, Hearts),
	NewCard(Two, Hearts),
	NewCard(Jack, Diamonds),
	NewCard(Eight, Diamonds),
	NewCard(Ace, Spades),
	NewCard(Three, Clubs),
	NewCard(Ten, Hearts),
	NewCard(King, Hearts),
	NewCard(Four, Spades),
	NewCard(Seven, Hearts),
	NewCard(Queen, Hearts),
	NewCard(Three, Spades),
	NewCard(Queen, Spades),
	NewCard(Seven, Clubs),
}

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
