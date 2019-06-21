package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"Time": "To Go"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("Time")
		expected := "To Go"

		assertStrings(t, got, expected)
	})

	t.Run("unbeknownst to man", func(t *testing.T) {
		_, got := dictionary.Search("Unbeknownst")

		assertError(t, got, ErrNotFound)
	})
}

func assertStrings(t *testing.T, got, expected string) {
	t.Helper()

	if got != expected {
		t.Errorf("got %s expected %s given %s", got, expected, "Time")
	}
}

func assertError(t *testing.T, got, expected error) {
	t.Helper()

	if got != expected {
		t.Errorf("got error %s want %s", got, expected)
	}
}
