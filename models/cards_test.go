package models

import (
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
	utils.SetRandomSeed(2021)
	code := m.Run()
	os.Exit(code)
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

var newShoeTests = []ShoeResult{
	{-1, 0, 0, 0},
	{0, 0, 0, 0},
	{1, cardsPerDeck, suitsPerDeck, ranksPerDeck},
	{2, 2 * cardsPerDeck, suitsPerDeck, ranksPerDeck},
	{3, 3 * cardsPerDeck, suitsPerDeck, ranksPerDeck},
}

func TestNewShoe(t *testing.T) {
	for _, test := range newShoeTests {
		testShoe(t, NewShoe(test.decks), test)
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

	if counter != test.expectedLength {
		t.Errorf("Expected %v cards in the shoe, not %v\n", test.expectedLength, counter)
	}

	if len(suits) != test.expectedSuits {
		t.Errorf("Expected %v suits in the shoe, not %v\n", test.expectedSuits, len(suits))
	}

	expectedCardsPerSuit := test.decks * test.expectedRanks
	for _, v := range suits {
		if v != expectedCardsPerSuit {
			t.Errorf("Expected %v cards per suit, not %v\n", expectedCardsPerSuit, v)
		}
	}

	if len(ranks) != test.expectedRanks {
		t.Errorf("Expected %v ranks in the shoe, not %v\n", ranksPerDeck, len(ranks))
	}

	expectedCardsPerRank := test.decks * test.expectedSuits
	for _, v := range ranks {
		if v != expectedCardsPerRank {
			t.Errorf("Expected %v cards per rank, not %v\n", expectedCardsPerRank, v)
		}
	}
}

var shuffleTest = shoe{
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

func TestShuffle(t *testing.T) {
	deck := NewShoe(1)
	deck.shuffle()

	deckResult := ShoeResult{1, cardsPerDeck, suitsPerDeck, ranksPerDeck}
	testShoe(t, deck, deckResult)

	for i, card := range deck {
		if got, want := card, shuffleTest[i]; got != want {
			t.Errorf("card = %v; want %v", got, want)
		}
	}
}
