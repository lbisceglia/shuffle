package utils

import (
	"testing"
)

// TODO: setup and takedown functions

type IntUnaryOpResult struct {
	name     string
	x        int
	expected int
}

type IntBinaryOpResult struct {
	x        int
	y        int
	expected int
}

var absTests = []IntUnaryOpResult{
	{"neg", -1, 1},
	{"zero", 0, 0},
	{"pos", 1, 1},
}

var minTests = []IntBinaryOpResult{
	{1, 2, 1},
	{5, 5, 5},
	{0, -1, -1},
}

var maxTests = []IntBinaryOpResult{
	{1, 2, 2},
	{5, 5, 5},
	{0, -1, 0},
}

var modTests = []IntBinaryOpResult{
	{9, 0, -1},
	{4, 5, 4},
	{-4, -5, 4},
	{-4, 5, 1},
	{4, -5, 1},
	{5, 5, 0},
	{-5, -5, 0},
	{-5, 5, 0},
	{5, -5, 0},
	{6, 5, 1},
	{-6, -5, 1},
	{-6, 5, 4},
	{6, -5, 4},
}

func TestAbs(t *testing.T) {
	for _, test := range absTests {
		result := Abs(test.x)
		if result != test.expected {
			t.Errorf("Expected value %d not given\n", test.expected)
		}
	}
}

func testBinaryOp(t *testing.T, tests []IntBinaryOpResult, f func(int, int) int) {
	for _, test := range tests {
		result := f(test.x, test.y)
		if result != test.expected {
			t.Errorf("Expected value %d not given\n", test.expected)
		}
	}
}

func TestMin(t *testing.T) {
	testBinaryOp(t, minTests, Min)
}

func TestMax(t *testing.T) {
	testBinaryOp(t, maxTests, Max)
}

func TestMod(t *testing.T) {
	for _, test := range modTests {
		result, err := Mod(test.x, test.y)
		if result != test.expected || result < 0 && err == nil {
			t.Errorf("Expected value %d not given\n", test.expected)
		}
	}
}
