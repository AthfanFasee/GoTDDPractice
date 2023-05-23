package integer

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

func TestAdd(t *testing.T) {
	t.Run("Adding two numbers", func(t *testing.T) {
		got := Add(2, 4)
		want := 6

		Equal(t, got, want)
	})
}

// Example function is just another test
// Output comment here matters
// By adding this code the example will appear in the documentation inside godoc
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}