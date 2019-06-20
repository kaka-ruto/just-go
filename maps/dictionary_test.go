package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"Time": "To Go"}

	got := Search(dictionary, "Time")
	expected := "To Go"

	if got != expected {
		t.Errorf("got %s expected %s given %s", got, expected, "Time")
	}
}
