package utils

import (
	"testing"
)

type IntUnaryOpResult struct {
	x    int
	want int
}

type IntBinaryOpResult struct {
	x    int
	y    int
	want int
}

func TestAbs(t *testing.T) {
	tests := map[string]IntUnaryOpResult{
		"neg":  {-1, 1},
		"zero": {0, 0},
		"pos":  {1, 1},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			Evaluate(t, Abs(test.x), test.want)
		})
	}
}

func testBinaryOp(t *testing.T, tests map[string]IntBinaryOpResult, f func(int, int) int) {
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			Evaluate(t, f(test.x, test.y), test.want)
		})
	}
}

func TestMin(t *testing.T) {
	tests := map[string]IntBinaryOpResult{
		"less":    {1, 2, 1},
		"equal":   {5, 5, 5},
		"greater": {0, -1, -1},
	}

	testBinaryOp(t, tests, Min)
}

func TestMax(t *testing.T) {
	tests := map[string]IntBinaryOpResult{
		"less":    {1, 2, 2},
		"equal":   {5, 5, 5},
		"greater": {0, -1, 0},
	}

	testBinaryOp(t, tests, Max)
}

func TestMod(t *testing.T) {
	tests := map[string]IntBinaryOpResult{
		"pos by zero":            {9, 0, -1},
		"neg by zero":            {-9, 0, -1},
		"zero by zero":           {0, 0, -1},
		"zero by pos":            {0, 5, 0},
		"zero by neg":            {0, -5, 0},
		"zero by any":            {0, -5, 0},
		"pos small by pos large": {4, 5, 4},
		"neg small by neg large": {-4, -5, 4},
		"neg small by pos large": {-4, 5, 1},
		"pos small by neg large": {4, -5, 1},
		"pos by pos self":        {5, 5, 0},
		"neg by neg self":        {-5, -5, 0},
		"neg by pos self":        {-5, 5, 0},
		"pos by neg self":        {5, -5, 0},
		"pos large by pos small": {6, 5, 1},
		"neg large by neg small": {-6, -5, 1},
		"neg large by pos small": {-6, 5, 4},
		"pos large by neg small": {6, -5, 4},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := Mod(test.x, test.y)
			want := test.want
			if got != want || got < 0 && err == nil {
				t.Fatalf("ans = %d; want %d\n", got, want)
			}
		})
	}
}
