package models

import (
	"fmt"
	"os"
	"shuffle/utils"
	"testing"
)

const (
	suitsPerDeck = 4
	ranksPerDeck = 13
	cardsPerDeck = 52
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func teardownRandomSubtest() {
	utils.Reset()
}

type CardStringResult struct {
	card         card
	expectedRank string
	expectedSuit string
	expectedCard string
}

func evaluate(t *testing.T, got interface{}, want interface{}, context ...string) {
	if got != want {
		if len(context) == 0 {
			t.Fatalf("got %v; want %v\n", got, want)
		} else {
			t.Fatalf("got %v %s; want %v\n", got, context[0], want)
		}
	}
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
			evaluate(t, test.card.rank.String(), test.expectedRank)
			evaluate(t, test.card.suit.String(), test.expectedSuit)
			evaluate(t, test.card.String(), test.expectedCard)
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

	evaluate(t, counter, test.expectedLength, "cards in the shoe")
	evaluate(t, len(suits), test.expectedSuits, "suits in the shoe")
	evaluate(t, len(ranks), test.expectedRanks, "ranks in the shoe")

	cardsPerSuit := test.decks * test.expectedRanks
	for suit, got := range suits {
		evaluate(t, got, cardsPerSuit, suit.String())
	}

	cardsPerRank := test.decks * test.expectedSuits
	for rank, got := range ranks {
		evaluate(t, got, cardsPerRank, rank.String())
	}
}

func TestShuffle(t *testing.T) {
	utils.SetRandomSeed(2021)
	defer teardownRandomSubtest()

	test := shoe{
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

	deck := NewShoe(1)
	deck.shuffle()

	deckResult := ShoeResult{1, cardsPerDeck, suitsPerDeck, ranksPerDeck}
	testShoe(t, deck, deckResult)

	for i, card := range deck {
		t.Run(fmt.Sprintf("card %v", i), func(t *testing.T) {
			evaluate(t, card, test[i])
		})
	}
}
