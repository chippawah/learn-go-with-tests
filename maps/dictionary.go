package main

type Dictionary map[string]string

// This is an error wrapper pattern to give more explicit and immutable error types
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	ErrNotFound         = DictionaryErr("word not found in dictionary")
	ErrWordExists       = DictionaryErr("the word already exists in the dictionary")
	ErrWordDoesNotExist = DictionaryErr("the word does not exist in the dictionary")
)

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

func (d Dictionary) Update(word, defintion string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = defintion
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, word)
	default:
		return err
	}
	return nil
}
