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
		_, err := dictionary.Search("Unbeknownst")
		expected := "Couldn't find the word you were looking for"

		if err == nil {
			t.Fatal("Couldn't find the word")
		}

		assertStrings(t, err.Error(), expected)
	})
}

func assertStrings(t *testing.T, got, expected string) {
	t.Helper()

	if got != expected {
		t.Errorf("got %s expected %s given %s", got, expected, "Time")
	}
}
