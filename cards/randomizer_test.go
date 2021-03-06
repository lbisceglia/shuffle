package cards

import (
	"shuffle/utils"
	"testing"

	// empty import is a temporary workaround for this issue: https://github.com/golang/mock/issues/415
	_ "github.com/golang/mock/mockgen/model"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

const (
	SEED = "seed"
	READ = "read"
)

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

type RandomSeedResult struct {
	method string
	n      int64
}

type ShuffleResult struct {
	seed     int64
	numDecks int
	want     Shoe
}

var p = message.NewPrinter(language.English)

func TestSetRandomSeed(t *testing.T) {
	tests := map[string][]RandomSeedResult{
		"positive seed": {
			{SEED, 1},
			{READ, 5577006791947779410},
			{READ, 8674665223082153551},
			{READ, 6129484611666145821},
			{READ, 4037200794235010051},
			{READ, 3916589616287113937},
			{READ, 6334824724549167320},
			{READ, 605394647632969758},
			{READ, 1443635317331776148},
			{READ, 894385949183117216},
			{READ, 2775422040480279449},
		},
		"zero seed": {
			{SEED, 0},
			{READ, 8717895732742165505},
			{READ, 2259404117704393152},
			{READ, 6050128673802995827},
			{READ, 501233450539197794},
			{READ, 3390393562759376202},
			{READ, 2669985732393126063},
			{READ, 1774932891286980153},
			{READ, 6044372234677422456},
			{READ, 8274930044578894929},
			{READ, 1543572285742637646},
		},
		"negative seed": {
			{SEED, -1},
			{READ, 3644962268338389676},
			{READ, 550171362161912239},
			{READ, 3094056749125766625},
			{READ, 5185619287625335803},
			{READ, 5303699392440114477},
			{READ, 2885068457955727142},
			{READ, 8617297895100121056},
			{READ, 7780200714598034794},
			{READ, 8329828311202461790},
			{READ, 6445856527061488741},
		},
		"idempotent seed": {
			{SEED, 1},
			{SEED, 0},
			{READ, 5577006791947779410},
			{READ, 8674665223082153551},
			{READ, 6129484611666145821},
			{SEED, -1},
			{READ, 4037200794235010051},
			{READ, 3916589616287113937},
			{READ, 6334824724549167320},
			{SEED, 1},
			{READ, 605394647632969758},
			{READ, 1443635317331776148},
			{READ, 894385949183117216},
			{READ, 2775422040480279449},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			rng := NewRng()

			for _, action := range test {
				switch action.method {
				case SEED:
					rng.Seed(action.n)
				case READ:
					if got, want := rng.r.Int63(), action.n; got != want {
						t.Fatalf("next rand = %d; want %d\n", got, want)
					}
				default:
					t.Fatalf("invalid seed action %v attempted\n", action.method)
				}
			}
		})
	}
}

func TestGenerateRandomSeed(t *testing.T) {
	set := make(map[int64]bool)
	ITERS := 100000

	for i := 0; i < ITERS; i++ {
		if seed, err := GenerateRandomSeed(); err != nil {
			t.Errorf("system's secure random number generator failed")
		} else if set[seed] {
			t.Errorf("random seed %v either not unique or hashing collision occured", seed)
		} else {
			set[seed] = true
		}
	}

	success := float64(len(set)) / float64(ITERS)
	pct := p.Sprint(number.Percent(success, number.Scale(2)))

	t.Logf("%v of randomly generated seeds unique (%v  of %v)\n", pct, p.Sprint(len(set)), p.Sprint(ITERS))
}

func TestRandomize(t *testing.T) {
	rng := NewRng()
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("the code panicked")
		}
	}()
	rng.Randomize()
}

func TestShuffle(t *testing.T) {
	tests := map[string]ShuffleResult{
		"no deck":         {2021, 0, Shoe{}},
		"single deck":     {2021, 1, shuffled2021Deck},
		"multi-deck shoe": {2021, 2, shuffled2021Shoe},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			rng := NewRngAt(test.seed)
			shoe := NewShoe(test.numDecks)
			rng.Shuffle(shoe)
			utils.Error(t, shoe, test.want)
		})
	}
}
