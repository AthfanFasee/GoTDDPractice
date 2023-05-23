package main

import "testing"

func Equal[T comparable](t testing.TB, actual, expected T) {
	t.Helper()

	if actual != expected {
		t.Errorf("got: %v; want: %v", actual, expected)
	}
}