package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"Time": "To Go"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("Time")
		expected := "To Go"

		assertStrings(t, got, expected)
	})

	t.Run("unbeknownst word", func(t *testing.T) {
		_, got := dictionary.Search("Unbeknownst")

		assertError(t, got, ErrNotFound)
	})
}

func TesAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("New", "This is a new value")

	expected := "This is a new value"
	got, err := dictionary.Search("New")

	if err != nil {
		t.Fatal("Should find added word:", err)
	}

	if expected != got {
		t.Errorf("Got '%s' expected '%s'", got, expected)
	}
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
