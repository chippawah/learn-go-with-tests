package main

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("word not found in dictionary")

func (d Dictionary) Search(word string) (string, error) {
	defintion, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return defintion, nil
}
