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

var cardStringTests = []CardStringResult{
	{NewCard(Ten, Clubs), "10", "♣", "10♣"},
	{NewCard(Jack, Diamonds), "J", "♦", "J♦"},
	{NewCard(Queen, Hearts), "Q", "♥", "Q♥"},
	{NewCard(King, Spades), "K", "♠", "K♠"},
	{NewCard(Ace, "elephants"), "A", "invalid suit", "A, invalid suit"},
	{NewCard("14", Spades), "invalid rank", "♠", "invalid rank, ♠"},
	{NewCard("14", "elephants"), "invalid rank", "invalid suit", "invalid rank, invalid suit"},
}

func TestCardString(t *testing.T) {
	for _, test := range cardStringTests {
		if got := test.card.rank.String(); got != test.expectedRank {
			t.Errorf("rank = %v; want %v", got, test.expectedRank)
		}

		if got := test.card.suit.String(); got != test.expectedSuit {
			t.Errorf("suit = %v; want %v", got, test.expectedSuit)
		}

		if got := test.card.String(); got != test.expectedCard {
			t.Errorf("card = %v; want %v", got, test.expectedCard)
		}
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

	if got, want := counter, test.expectedLength; got != want {
		t.Errorf("got %v cards in the shoe; want %v\n", got, want)
	}

	if got, want := len(suits), test.expectedSuits; got != want {
		t.Errorf("got %v suits in the shoe; want %v\n", got, want)
	}

	if got, want := len(ranks), test.expectedRanks; got != want {
		t.Errorf("got %v ranks in the shoe; want %v\n", got, want)
	}

	cardsPerSuit := test.decks * test.expectedRanks
	for suit, got := range suits {
		if want := cardsPerSuit; got != want {
			t.Errorf("got %v %v; want %v\n", got, suit, want)
		}
	}

	cardsPerRank := test.decks * test.expectedSuits
	for rank, got := range ranks {
		if want := cardsPerRank; got != want {
			t.Errorf("got %v %v; want %v\n", got, rank, want)
		}
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
			if got, want := card, test[i]; got != want {
				t.Fatalf("got %v; want %v", got, want)
			}
		})
	}
}
