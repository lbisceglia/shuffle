package utils

import (
	"reflect"
	"testing"
)

// Fatal checks that got equals want, and immediately fails the test if they are not.
// It logs the details of the failed test, which are displayed if the verbose flag is enabled.
func Fatal(t *testing.T, got interface{}, want interface{}, context ...string) {
	Evaluate(t, got, want, true, context...)
}

// Error checks that got equals want, and fails the test if they are not.
// It logs the details of the failed test, which are displayed if the verbose flag is enabled.
func Error(t *testing.T, got interface{}, want interface{}, context ...string) {
	Evaluate(t, got, want, false, context...)
}

// Evaluate checks that got equals want, and fails the test either immediately (fatal = true)
// or eventually (fatal = false) if they are not.
// It logs the details of the failed test, which are displayed if the verbose flag is enabled.
func Evaluate(t *testing.T, got interface{}, want interface{}, fatal bool, context ...string) {
	if !reflect.DeepEqual(got, want) {

		if len(context) == 0 {
			t.Logf("got %v; want %v\n", got, want)
		} else {
			t.Logf("got %v %s; want %v\n", got, context[0], want)
		}

		if fatal {
			t.FailNow()
		} else {
			t.Fail()
		}
	}
}
