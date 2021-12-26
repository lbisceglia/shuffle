package models

import (
	"fmt"
	"shuffle/utils"
	"testing"
)

const (
	DRAW    = "draw"
	DISCARD = "discard"
)

type HandResult struct {
	name                string
	method              string
	cardsRequested      int
	cardsTransferred    int
	drawExpectedSize    int
	discardExpectedSize int
	handTransferred     hand
}

type DealerResult struct {
	numCards int
	shuffle  bool
	draw     shoe
}

func TestDealHand(t *testing.T) {
	reshuffled := shoe{
		NewCard(Seven, Diamonds),
		NewCard(Nine, Spades),
		NewCard(Eight, Clubs),
		NewCard(Jack, Diamonds),
		NewCard(Three, Spades),
		NewCard(Ten, Spades),
		NewCard(Six, Clubs),
		NewCard(Two, Diamonds),
		NewCard(Seven, Spades),
		NewCard(King, Hearts),
		NewCard(Ten, Clubs),
		NewCard(Nine, Hearts),
		NewCard(Eight, Hearts),
		NewCard(Eight, Diamonds),
		NewCard(Queen, Clubs),
		NewCard(Ten, Hearts),
		NewCard(Ace, Diamonds),
		NewCard(Five, Diamonds),
		NewCard(Six, Spades),
		NewCard(Ten, Diamonds),
		NewCard(Jack, Clubs),
		NewCard(Five, Spades),
		NewCard(Two, Hearts),
		NewCard(Four, Spades),
		NewCard(Ace, Spades),
		NewCard(Three, Clubs),
		NewCard(Nine, Diamonds),
		NewCard(Jack, Spades),
		NewCard(Ace, Hearts),
		NewCard(Jack, Hearts),
		NewCard(Nine, Clubs),
		NewCard(Six, Hearts),
		NewCard(Six, Diamonds),
		NewCard(Four, Clubs),
		NewCard(King, Diamonds),
		NewCard(Queen, Hearts),
		NewCard(King, Clubs),
		NewCard(Five, Hearts),
		NewCard(Queen, Spades),
		NewCard(Seven, Hearts),
		NewCard(Ace, Clubs),
		NewCard(Three, Diamonds),
		NewCard(Four, Diamonds),
		NewCard(Queen, Diamonds),
		NewCard(King, Spades),
		NewCard(Five, Clubs),
		NewCard(Eight, Spades),
		NewCard(Three, Hearts),
		NewCard(Two, Spades),
		NewCard(Four, Hearts),
		NewCard(Seven, Clubs),
		NewCard(Two, Clubs),
	}

	tests := map[string][]HandResult{
		"no reshuffle": {
			{"empty request", DRAW, 0, 0, 52, 0, hand{}},
			{"empty idempotent", DRAW, 0, 0, 52, 0, hand{}},
			{"single request, full fill", DRAW, 1, 1, 51, 0, hand(shuffled2021Deck[:1])},
			{"large request, full fill", DRAW, 50, 50, 1, 0, hand(shuffled2021Deck[1:51])},
			{"large request, partial fill", DRAW, 17, 1, 0, 0, hand(shuffled2021Deck[51:])},
			{"single request, empty fill", DRAW, 1, 0, 0, 0, hand{}},
			{"large request, empty fill", DRAW, 17, 0, 0, 0, hand{}},
		},
		"reshuffle after exhaustion": {
			{"large request, full fill", DRAW, 52, 52, 0, 0, hand(shuffled2021Deck)},
			{"replenish", DISCARD, 0, 52, 0, 52, hand(shuffled2021Deck)},
			{"full fill, after reshuffle", DRAW, 20, 20, 32, 0, hand(reshuffled[:20])},
		},
		"reshuffle during draw, full fill": {
			{"large request, full fill", DRAW, 51, 51, 1, 0, hand(shuffled2021Deck[:51])},
			{"replenish", DISCARD, 0, 2, 1, 2, hand{NewCard(Jack, Spades), NewCard(Five, Hearts)}},
			{"full fill with intermittent reshuffle", DRAW, 2, 2, 1, 0, hand{NewCard(Four, Clubs), NewCard(Jack, Spades)}},
		},
		"reshuffle during draw, partial fill": {
			{"large request, full fill", DRAW, 51, 51, 1, 0, hand(shuffled2021Deck[:51])},
			{"replenish", DISCARD, 0, 2, 1, 2, hand{NewCard(Jack, Spades), NewCard(Five, Hearts)}},
			{"partial fill with intermittent reshuffle", DRAW, 4, 3, 0, 0, hand{NewCard(Four, Clubs), NewCard(Jack, Spades), NewCard(Five, Hearts)}},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			utils.SetRandomSeed(2021)
			defer utils.TeardownRandomSubtest()

			d := NewDealer(1)

			for _, handResult := range test {
				switch handResult.method {
				case DRAW:
					t.Run(handResult.name, func(t *testing.T) {
						hand := d.DealHand(handResult.cardsRequested)
						utils.Evaluate(t, len(hand), handResult.cardsTransferred, "cards dealt")
						utils.Evaluate(t, d.drawSize(), handResult.drawExpectedSize, "draw cards remaining")
						utils.Evaluate(t, d.discardSize(), handResult.discardExpectedSize, "cards in discard")
					})
				case DISCARD:
					t.Run(handResult.name, func(t *testing.T) {
						before := d.drawSize() + d.discardSize()
						d.HandleDiscard(handResult.handTransferred)
						after := d.drawSize() + d.discardSize()
						utils.Evaluate(t, after-before, handResult.cardsTransferred, "cards transferred to dealer")
						utils.Evaluate(t, d.drawSize(), handResult.drawExpectedSize, "draw cards remaining")
						utils.Evaluate(t, d.discardSize(), handResult.discardExpectedSize, "cards in discard")
					})
				default:
					t.Fatalf("invalid dealer action attempt")
				}
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

	test := shuffled2021Deck

	d.ReplaceShoe(1)
	draw = d.drawPile()
	utils.Evaluate(t, len(draw), 52)

	for i, card := range draw {
		t.Run(fmt.Sprintf("card %v, replaced deck", i), func(t *testing.T) {
			utils.Evaluate(t, card, test[i])
		})
	}
}

func TestNewDealer(t *testing.T) {
	tests := map[string]DealerResult{
		"unshuffled, single":   {1, false, standardDeck},
		"unshuffled, multiple": {2, false, append(standardDeck, standardDeck...)},
		"shuffled, single":     {1, true, shuffled2021Deck},
		"shuffled, multiple":   {2, true, shuffled2021Shoe},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			utils.SetRandomSeed(2021)
			defer utils.TeardownRandomSubtest()
			d := NewDealer(test.numCards, test.shuffle)
			for i, card := range d.drawPile() {
				utils.Evaluate(t, card, test.draw[i])
			}
		})
	}
}

func TestHandleDiscard(t *testing.T) {
	// TODO after mocking refactor
}
