package models

import (
	"fmt"
	"shuffle/utils"
	"testing"
)

type HandResult struct {
	name               string
	cardsRequested     int
	cardsExpected      int
	drawCardsRemaining int
	handExpected       hand
}

func TestDealHand(t *testing.T) {
	utils.SetRandomSeed(2021)
	defer utils.TeardownRandomSubtest()

	tests := []HandResult{
		{"empty request", 0, 0, 52, hand{}},
		{"empty idempotent", 0, 0, 52, hand{}},
		{"single request, full fill", 1, 1, 51, hand{NewCard(Jack, Spades)}},
		{"large request, full fill", 50, 50, 1, hand{
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
		}},
		{"large request, partial fill", 5, 1, 0, hand{NewCard(Four, Clubs)}},
		{"single request, empty fill", 1, 0, 0, hand{}},
		{"large request, empty fill", 6, 0, 0, hand{}},
	}

	d := NewDealer(1, false)
	d.Shuffle()

	for i, handResult := range tests {
		hand := d.DealHand(handResult.cardsRequested)
		t.Run(handResult.name, func(t *testing.T) {
			utils.Evaluate(t, len(hand), handResult.cardsExpected, "cards dealt")
			utils.Evaluate(t, len(d.drawPile()), handResult.drawCardsRemaining, "draw cards remaining")

			for j, card := range hand {
				utils.Evaluate(t, card, handResult.handExpected[j], fmt.Sprintf("(card %v, hand %v)", i, j))
			}
		})
	}
}

func TestReplaceShoe(t *testing.T) {
	utils.SetRandomSeed(2021)
	defer utils.TeardownRandomSubtest()

	d := NewDealer(1, false)
	draw := d.drawPile()

	for i, card := range draw {
		t.Run(fmt.Sprintf("card %d, standard deck", i), func(t *testing.T) {
			utils.Evaluate(t, card, standardDeck[i])
		})
	}

	utils.Evaluate(t, len(draw), 52)

	SIZE := 5
	d.DealHand(SIZE)
	draw = d.drawPile()

	utils.Evaluate(t, len(draw), 52-SIZE)

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

	d.ReplaceShoe(1)
	draw = d.drawPile()
	utils.Evaluate(t, len(draw), 52)

	for i, card := range draw {
		t.Run(fmt.Sprintf("card %v, replaced deck", i), func(t *testing.T) {
			utils.Evaluate(t, card, test[i])
		})
	}
}

// TODO: TestNewDealer
// TODO: TestHandleDiscard
// TODO: TestDealHandWithReshuffle
// TODO: TestReshuffle
