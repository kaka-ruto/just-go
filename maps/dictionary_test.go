package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"Time": "To Go"}

	got := dictionary.Search("Time")
	expected := "To Go"

	assertStrings(t, got, expected)
}

func assertStrings(t *testing.T, got, expected string) {
	t.Helper()

	if got != expected {
		t.Errorf("got %s expected %s given %s", got, expected, "Time")
	}
}
