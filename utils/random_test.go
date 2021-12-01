package utils

import (
	"math/rand"
	"os"
	"testing"
)

const (
	SEED = "seed"
	READ = "read"
)

type RandomSeedResult struct {
	method string
	n      int64
}

var randomSeedTests = [][]RandomSeedResult{
	{
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
	{
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
	{
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
	{
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

var set map[int64]bool

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestSetRandomSeed(t *testing.T) {
	for _, test := range randomSeedTests {
		for _, action := range test {
			switch action.method {
			case SEED:
				SetRandomSeed(action.n)
			case READ:
				if got := rand.Int63(); got != action.n {
					t.Errorf("next random = %d; want %d", got, action.n)
				}
			default:
				t.Errorf("invalid seed action attempted")
			}
		}
		reset()
	}
}

func TestGenerateRandomSeed(t *testing.T) {
	set = make(map[int64]bool)
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
}

func TestRandomize(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("the code panicked")
		}
		reset()
	}()
	Randomize()
}
