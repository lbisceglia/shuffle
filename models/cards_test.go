package models

import (
	"shuffle/utils"
	"testing"
)

const (
	suitsPerDeck = 4
	ranksPerDeck = 13
	cardsPerDeck = 52
)

var standardDeck = shoe{
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

var shuffled2021Deck = shoe{
	NewCard(Jack, Spades),
	NewCard(Five, Hearts),
	NewCard(Eight, Diamonds),
	NewCard(Ten, Spades),
	NewCard(Nine, Spades),
	NewCard(Five, Spades),
	NewCard(King, Hearts),
	NewCard(Ace, Diamonds),
	NewCard(Ten, Diamonds),
	NewCard(Three, Diamonds),
	NewCard(Queen, Clubs),
	NewCard(Two, Spades),
	NewCard(Jack, Diamonds),
	NewCard(Queen, Hearts),
	NewCard(Seven, Clubs),
	NewCard(Seven, Spades),
	NewCard(Queen, Diamonds),
	NewCard(Five, Diamonds),
	NewCard(Three, Spades),
	NewCard(Nine, Clubs),
	NewCard(King, Clubs),
	NewCard(Three, Clubs),
	NewCard(Eight, Hearts),
	NewCard(Four, Hearts),
	NewCard(Ace, Hearts),
	NewCard(Six, Clubs),
	NewCard(Ace, Clubs),
	NewCard(Six, Diamonds),
	NewCard(Eight, Spades),
	NewCard(Seven, Diamonds),
	NewCard(Two, Clubs),
	NewCard(Queen, Spades),
	NewCard(Four, Spades),
	NewCard(King, Spades),
	NewCard(Five, Clubs),
	NewCard(Six, Spades),
	NewCard(Three, Hearts),
	NewCard(Eight, Clubs),
	NewCard(King, Diamonds),
	NewCard(Ten, Hearts),
	NewCard(Two, Diamonds),
	NewCard(Jack, Clubs),
	NewCard(Nine, Hearts),
	NewCard(Ten, Clubs),
	NewCard(Two, Hearts),
	NewCard(Nine, Diamonds),
	NewCard(Ace, Spades),
	NewCard(Seven, Hearts),
	NewCard(Six, Hearts),
	NewCard(Jack, Hearts),
	NewCard(Four, Diamonds),
	NewCard(Four, Clubs),
}

var shuffled2021Shoe = shoe{
	NewCard(Jack, Spades),
	NewCard(Five, Hearts),
	NewCard(Eight, Diamonds),
	NewCard(Five, Diamonds),
	NewCard(Nine, Spades),
	NewCard(Five, Spades),
	NewCard(King, Hearts),
	NewCard(Ace, Diamonds),
	NewCard(Ten, Diamonds),
	NewCard(Three, Diamonds),
	NewCard(Eight, Hearts),
	NewCard(Nine, Hearts),
	NewCard(Ace, Hearts),
	NewCard(Three, Clubs),
	NewCard(Queen, Spades),
	NewCard(Six, Spades),
	NewCard(Queen, Diamonds),
	NewCard(Five, Diamonds),
	NewCard(Ten, Hearts),
	NewCard(Nine, Clubs),
	NewCard(King, Clubs),
	NewCard(Three, Clubs),
	NewCard(Eight, Hearts),
	NewCard(Four, Hearts),
	NewCard(Nine, Clubs),
	NewCard(Six, Clubs),
	NewCard(Queen, Diamonds),
	NewCard(Six, Diamonds),
	NewCard(Eight, Spades),
	NewCard(King, Clubs),
	NewCard(Seven, Spades),
	NewCard(Queen, Clubs),
	NewCard(Queen, Hearts),
	NewCard(King, Spades),
	NewCard(King, Hearts),
	NewCard(Seven, Clubs),
	NewCard(Three, Hearts),
	NewCard(Eight, Clubs),
	NewCard(King, Diamonds),
	NewCard(Ace, Clubs),
	NewCard(Two, Diamonds),
	NewCard(Two, Clubs),
	NewCard(Nine, Hearts),
	NewCard(Ten, Clubs),
	NewCard(Ten, Spades),
	NewCard(Nine, Diamonds),
	NewCard(Nine, Diamonds),
	NewCard(Seven, Hearts),
	NewCard(Four, Spades),
	NewCard(Three, Hearts),
	NewCard(Four, Diamonds),
	NewCard(Four, Clubs),
	NewCard(Ten, Hearts),
	NewCard(Jack, Clubs),
	NewCard(Queen, Hearts),
	NewCard(Eight, Clubs),
	NewCard(Eight, Diamonds),
	NewCard(Two, Spades),
	NewCard(Five, Spades),
	NewCard(Two, Diamonds),
	NewCard(Three, Diamonds),
	NewCard(Two, Clubs),
	NewCard(Seven, Diamonds),
	NewCard(Queen, Spades),
	NewCard(Jack, Clubs),
	NewCard(Jack, Spades),
	NewCard(Six, Hearts),
	NewCard(Ace, Hearts),
	NewCard(Four, Clubs),
	NewCard(Ten, Spades),
	NewCard(Seven, Clubs),
	NewCard(Ace, Diamonds),
	NewCard(Three, Spades),
	NewCard(Eight, Spades),
	NewCard(Seven, Hearts),
	NewCard(Five, Clubs),
	NewCard(Four, Hearts),
	NewCard(Jack, Diamonds),
	NewCard(Two, Spades),
	NewCard(Ace, Spades),
	NewCard(Jack, Hearts),
	NewCard(Ace, Clubs),
	NewCard(Ace, Spades),
	NewCard(Four, Diamonds),
	NewCard(Six, Spades),
	NewCard(Queen, Clubs),
	NewCard(Two, Hearts),
	NewCard(Nine, Spades),
	NewCard(King, Spades),
	NewCard(King, Diamonds),
	NewCard(Five, Clubs),
	NewCard(Six, Clubs),
	NewCard(Jack, Diamonds),
	NewCard(Four, Spades),
	NewCard(Six, Hearts),
	NewCard(Ten, Diamonds),
	NewCard(Seven, Spades),
	NewCard(Ten, Clubs),
	NewCard(Five, Hearts),
	NewCard(Three, Spades),
	NewCard(Seven, Diamonds),
	NewCard(Two, Hearts),
	NewCard(Six, Diamonds),
	NewCard(Jack, Hearts),
}

type CardStringResult struct {
	card         card
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
			utils.Evaluate(t, test.card.rank.String(), test.expectedRank)
			utils.Evaluate(t, test.card.suit.String(), test.expectedSuit)
			utils.Evaluate(t, test.card.String(), test.expectedCard)
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

func testShoe(t *testing.T, s shoe, test ShoeResult) {
	suits := make(map[suit]int)
	ranks := make(map[rank]int)

	counter := 0
	for _, c := range s {
		suits[c.suit]++
		ranks[c.rank]++
		counter++
	}

	utils.Evaluate(t, counter, test.expectedLength, "cards in the shoe")
	utils.Evaluate(t, len(suits), test.expectedSuits, "suits in the shoe")
	utils.Evaluate(t, len(ranks), test.expectedRanks, "ranks in the shoe")

	cardsPerSuit := test.decks * test.expectedRanks
	for suit, got := range suits {
		utils.Evaluate(t, got, cardsPerSuit, suit.String())
	}

	cardsPerRank := test.decks * test.expectedSuits
	for rank, got := range ranks {
		utils.Evaluate(t, got, cardsPerRank, rank.String())
	}
}

// TODO: remove
// func TestShuffle(t *testing.T) {
// 	utils.SetRandomSeed(2021)
// 	defer utils.TeardownRandomSubtest()

// 	test := shuffled2021Deck

// 	deck := NewShoe(1)
// 	deck.shuffle()

// 	deckResult := ShoeResult{1, cardsPerDeck, suitsPerDeck, ranksPerDeck}
// 	testShoe(t, deck, deckResult)

// 	for i, card := range deck {
// 		t.Run(fmt.Sprintf("card %v", i), func(t *testing.T) {
// 			utils.Evaluate(t, card, test[i])
// 		})
// 	}
// }
