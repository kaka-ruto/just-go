package main

const (
	ErrNotFound   = DictionaryErr("Couldn't find the word you were looking for")
	ErrWordExists = DictionaryErr("Cannot add word because it already exists")
)

type Dictionary map[string]string
type DictionaryErr string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (e DictionaryErr) Error() string {
	return string(e)
}
