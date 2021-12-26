package utils

import (
	"testing"
)

func TeardownRandomSubtest() {
	Reset()
}

func Evaluate(t *testing.T, got interface{}, want interface{}, context ...string) {
	if got != want {
		if len(context) == 0 {
			t.Fatalf("got %v; want %v\n", got, want)
		} else {
			t.Fatalf("got %v %s; want %v\n", got, context[0], want)
		}
	}
}
