package main

import "testing"

func TestPrintAthfan(t *testing.T) {
	got := PrintAthfan()
	want := "Athfan"

	Equal(t, got, want)
}

func TestHello(t *testing.T) {
	// Both t.Run here are called Subtests
	t.Run("saying hello to someone", func(t *testing.T) {
		got := Hello("Athfan", "")
		want := "Hello, Athfan"

		Equal(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		Equal(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		Equal(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Elodie", "French")
		want := "Bonjour, Elodie"
		Equal(t, got, want)
	})
}