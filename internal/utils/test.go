package utils

import (
"runtime"
"testing"
)

func ShouldBeEqual(t *testing.T, expected, actual interface{}) {
	if expected != actual {
		_, fn, line, _ := runtime.Caller(1)
		t.Fatalf("\nExp: %v\nGot: %v\nLocation: %s:%d", expected, actual, fn, line)
	}
}

