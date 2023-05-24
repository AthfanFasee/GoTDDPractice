package arrayslice

import "testing"

func Equal[T comparable](t testing.TB, actual, expected T, given []int) {
	t.Helper()

	if actual != expected {
		t.Errorf("got: %v; want: %v; given %v", actual, expected, given)
	}
}

func TestSum(t *testing.T) {
	slice := []int{2, 2, 2}
	got := Sum(slice)
	want := 6

	Equal(t, got, want, slice)
}