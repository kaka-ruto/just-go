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
	t.Run("New Word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "New"
		definition := "This is a new value"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("Existing Word", func(t *testing.T) {
		word := "New"
		definition := "This is a new value"
		dictionary := Dictionary{word: definition}
		newDefinition := "Just another value"

		err := dictionary.Add(word, newDefinition)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
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

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {

	t.Helper()

	expected := "This is a new value"
	got, err := dictionary.Search("New")

	if err != nil {
		t.Fatal("Should find added word:", err)
	}

	if expected != got {
		t.Errorf("Got '%s' expected '%s'", got, expected)
	}
}
