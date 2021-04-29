package main

import "errors"

var ErrNotFound = errors.New("could not find the word you were looking for")

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	if val, ok := d[word]; ok {
		return val, nil
	} 

	return "", ErrNotFound
}

func (d Dictionary) Add(key, value string) {
	d[key] = value
}