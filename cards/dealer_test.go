package cards

import (
	"shuffle/utils"
	"testing"

	gomock "github.com/golang/mock/gomock"
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
	shuffles            int
	handDiscarded       Hand
}

type DealerResult struct {
	numCards int
	shuffle  bool
	shuffles int
}

var ctrl *gomock.Controller
var random *MockRandomizer
var d *dealer

func SetupTest(t *testing.T) {
	ctrl = gomock.NewController(t)
	random = NewMockRandomizer(ctrl)
}

func TestDealAndDiscard(t *testing.T) {

	tests := map[string][]HandResult{
		// {name, method, cardsRequested, cardsTransferred, drawExpectedSize, discardExpectedSize, shuffles, handDiscarded}
		"no reshuffle": {
			{"empty request", DRAW, 0, 0, 52, 0, 0, nil},
			{"empty idempotent", DRAW, 0, 0, 52, 0, 0, nil},
			{"single request, full fill", DRAW, 1, 1, 51, 0, 0, nil},
			{"large request, full fill", DRAW, 50, 50, 1, 0, 0, nil},
			{"large request, partial fill", DRAW, 17, 1, 0, 0, 1, nil},
			{"single request, empty fill", DRAW, 1, 0, 0, 0, 0, nil},
			{"large request, empty fill", DRAW, 17, 0, 0, 0, 0, nil},
		},
		"reshuffle after exhaustion": {
			{"large request, full fill", DRAW, 52, 52, 0, 0, 1, nil},
			{"replenish", DISCARD, 0, 52, 0, 52, 0, Hand(NewShoe(1))},
			{"full fill, after reshuffle", DRAW, 20, 20, 32, 0, 1, nil},
		},
		"reshuffle during draw, full fill": {
			{"large request, full fill", DRAW, 51, 51, 1, 0, 0, nil},
			{"replenish", DISCARD, 0, 2, 1, 2, 0, Hand{NewCard(Jack, Spades), NewCard(Five, Hearts)}},
			{"full fill with intermittent reshuffle", DRAW, 2, 2, 1, 0, 1, nil},
		},
		"reshuffle during draw, partial fill": {
			{"large request, full fill", DRAW, 51, 51, 1, 0, 0, nil},
			{"replenish", DISCARD, 0, 2, 1, 2, 0, Hand{NewCard(Jack, Spades), NewCard(Five, Hearts)}},
			{"partial fill with intermittent reshuffle", DRAW, 4, 3, 0, 0, 2, nil},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			SetupTest(t)
			d = NewDealer(1, random, false)
			for _, handResult := range test {
				t.Run(handResult.name, func(t *testing.T) {

					random.EXPECT().Shuffle(gomock.Any()).Times(handResult.shuffles)

					switch handResult.method {
					case DRAW:
						hand := d.DealHand(handResult.cardsRequested)
						utils.Error(t, len(hand), handResult.cardsTransferred, "cards dealt")
						utils.Error(t, d.drawSize(), handResult.drawExpectedSize, "draw cards remaining")
						utils.Error(t, d.discardSize(), handResult.discardExpectedSize, "cards in discard")
					case DISCARD:
						before := d.drawSize() + d.discardSize()
						d.HandleDiscard(handResult.handDiscarded)
						after := d.drawSize() + d.discardSize()
						utils.Error(t, after-before, handResult.cardsTransferred, "cards transferred to dealer")
						utils.Error(t, d.drawSize(), handResult.drawExpectedSize, "draw cards remaining")
						utils.Error(t, d.discardSize(), handResult.discardExpectedSize, "cards in discard")
					default:
						t.Fatalf("invalid dealer action attempt")
					}
				})
			}
		})
	}
}

func TestReplaceShoe(t *testing.T) {
	SetupTest(t)
	random.EXPECT().Shuffle(gomock.Any()).Times(2)

	d := NewDealer(1, random)
	draw := d.drawPile()
	utils.Error(t, len(draw), 52)

	SIZE := 5
	d.DealHand(SIZE)
	draw = d.drawPile()
	utils.Error(t, len(draw), 52-SIZE)

	d.ReplaceShoe(1)
	draw = d.drawPile()
	utils.Fatal(t, len(draw), 52)
}

func TestNewDealer(t *testing.T) {
	tests := map[string]DealerResult{
		"unshuffled, single":   {1, false, 0},
		"unshuffled, multiple": {2, false, 0},
		"shuffled, single":     {1, true, 1},
		"shuffled, multiple":   {2, true, 1},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			random.EXPECT().Shuffle(gomock.Any()).Times(test.shuffles)

			NewDealer(test.numCards, random, test.shuffle)
		})
	}
}
