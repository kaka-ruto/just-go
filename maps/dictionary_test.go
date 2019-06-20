package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"Time": "To Go"}

	got := Search(dictionary, "Time")
	expected := "To Go"

	assertStrings(t, got, expected)
}

func assertStrings(t *testing.T, got, expected string) {
	t.Helper()

	if got != expected {
		t.Errorf("got %s expected %s given %s", got, expected, "Time")
	}
}
