package iteration

import (
	"fmt"
	"testing"
)

func Equal[T comparable](t testing.TB, actual, expected T) {
	t.Helper()

	if actual != expected {
		t.Errorf("got: %v; want: %v", actual, expected)
	}
}

func TestRepeat(t *testing.T) {
	got := Repeat("a", 6)
	want := "aaaaaa"

	Equal(t, got, want)
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	result := Repeat("a", 6)
	fmt.Println(result)
	// Output: aaaaaa
}