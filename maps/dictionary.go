package main

type Dictionary map[string]string

const (
	ErrNotFound   = DictionaryErr("word not found in dictionary")
	ErrWordExists = DictionaryErr("the word already exists in the dictionary")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	defintion, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return defintion, nil
}

func (d Dictionary) Add(word, defintion string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = defintion
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}
