package utils

import (
	"math/rand"
	"os"
	"testing"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

const (
	SEED = "seed"
	READ = "read"
)

type RandomSeedResult struct {
	method string
	n      int64
}

var p = message.NewPrinter(language.English)

func teardownSubtest() {
	reset()
}

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestSetRandomSeed(t *testing.T) {
	tests := map[string][]RandomSeedResult{
		"positive seed": {
			{SEED, 2},
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
			defer func() {
				teardownSubtest()
			}()
			for _, action := range test {
				switch action.method {
				case SEED:
					SetRandomSeed(action.n)
				case READ:
					if got, want := rand.Int63(), action.n; got != want {
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
	ITERS := 1000000

	for i := 0; i < ITERS; i++ {
		if seed, err := generateRandomSeed(); err != nil {
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
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("the code panicked")
		}
		teardownSubtest()
	}()
	Randomize()
}
