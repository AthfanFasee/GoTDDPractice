package arrayslice

import (
	"reflect"
	"testing"
)

func Equal[T comparable](t testing.TB, actual, expected T, given []int) {
	t.Helper()

	if actual != expected {
		t.Errorf("got: %v; want: %v; given %v", actual, expected, given)
	}
}

func TestSum(t *testing.T) {
	got := Sum([]int{2, 2, 2})
	want := 6

	Equal(t, got, want, []int{2, 2, 2})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 3}, []int{1, 5})
	want := []int{4, 6}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v; want %v; given %v, %v", got, want, []int{1, 3}, []int{1, 5})
	}
}

func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll([]int{1, 3}, []int{1, 5})
	}
}

func TestSumAllVersinTwo(t *testing.T) {
	got := SumAll([]int{1, 3}, []int{1, 5})
	want := []int{4, 6}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v; want %v; given %v, %v", got, want, []int{1, 3}, []int{1, 5})
	}
}

func BenchmarkSumAllVersionTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll([]int{1, 3}, []int{1, 5})
	}
}

func TestSumAllTails(t *testing.T) {
	// New technique, assigning a function to a variable
	// By defining this function inside the test, it cannot be used by other functions in this package. 
	// Hiding variables and functions that don't need to be exported is an important design consideration.
	checkSums := func(t *testing.T, got, want, given1, given2 []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v; want %v; given %v, %v", got, want, given1, given2)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 3}, []int{1, 5})
		want := []int{3, 5}

		checkSums(t, got, want, []int{1, 3}, []int{1, 5})

	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 5})
		want := []int{0, 5}

		checkSums(t, got, want, []int{}, []int{1, 5})
	})
}