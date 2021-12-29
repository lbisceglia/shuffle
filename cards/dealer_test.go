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
	handTransferred     Hand
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
	d = NewDealer(1, random, false)
}

func TeardownTest() {
	ctrl.Finish()
}

func TestDealHand(t *testing.T) {
	reshuffled := Shoe{
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
		// {name, method, cardsRequested, cardsTransferred, drawExpectedSize, discardExpectedSize, shuffles, handTransferred}
		"no reshuffle": {
			{"empty request", DRAW, 0, 0, 52, 0, 0, Hand{}},
			{"empty idempotent", DRAW, 0, 0, 52, 0, 0, Hand{}},
			{"single request, full fill", DRAW, 1, 1, 51, 0, 0, Hand(shuffled2021Deck[:1])},
			{"large request, full fill", DRAW, 50, 50, 1, 0, 0, Hand(shuffled2021Deck[1:51])},
			{"large request, partial fill", DRAW, 17, 1, 0, 0, 1, Hand(shuffled2021Deck[51:])},
			{"single request, empty fill", DRAW, 1, 0, 0, 0, 0, Hand{}},
			{"large request, empty fill", DRAW, 17, 0, 0, 0, 0, Hand{}},
		},
		"reshuffle after exhaustion": {
			{"large request, full fill", DRAW, 52, 52, 0, 0, 1, Hand(shuffled2021Deck)},
			{"replenish", DISCARD, 0, 52, 0, 52, 0, Hand(shuffled2021Deck)},
			{"full fill, after reshuffle", DRAW, 20, 20, 32, 0, 1, Hand(reshuffled[:20])},
		},
		"reshuffle during draw, full fill": {
			{"large request, full fill", DRAW, 51, 51, 1, 0, 0, Hand(shuffled2021Deck[:51])},
			{"replenish", DISCARD, 0, 2, 1, 2, 0, Hand{NewCard(Jack, Spades), NewCard(Five, Hearts)}},
			{"full fill with intermittent reshuffle", DRAW, 2, 2, 1, 0, 1, Hand{NewCard(Four, Clubs), NewCard(Jack, Spades)}},
		},
		"reshuffle during draw, partial fill": {
			{"large request, full fill", DRAW, 51, 51, 1, 0, 0, Hand(shuffled2021Deck[:51])},
			{"replenish", DISCARD, 0, 2, 1, 2, 0, Hand{NewCard(Jack, Spades), NewCard(Five, Hearts)}},
			{"partial fill with intermittent reshuffle", DRAW, 4, 3, 0, 0, 2, Hand{NewCard(Four, Clubs), NewCard(Jack, Spades), NewCard(Five, Hearts)}},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			SetupTest(t)
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
						d.HandleDiscard(handResult.handTransferred)
						after := d.drawSize() + d.discardSize()
						utils.Error(t, after-before, handResult.cardsTransferred, "cards transferred to dealer")
						utils.Error(t, d.drawSize(), handResult.drawExpectedSize, "draw cards remaining")
						utils.Error(t, d.discardSize(), handResult.discardExpectedSize, "cards in discard")
					default:
						t.Fatalf("invalid dealer action attempt")
					}
				})
			}
			TeardownTest()
		})
	}
}

func TestReplaceShoe(t *testing.T) {
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
