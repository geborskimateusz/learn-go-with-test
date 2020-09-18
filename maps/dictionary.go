package maps

import (
	"errors"
)

// ErrKeyNotFound thrown when key not found in map
var ErrKeyNotFound = errors.New("There is no such key in dictionary")
var ErrDuplicateKey = errors.New("Key already exist")

// Dictionary stores key:value map of strings
type Dictionary map[string]string

// Search for key in Dictoinary
func (d Dictionary) Search(key string) (string, error) {
	foundEl, ok := d[key]

	if !ok {
		return "", ErrKeyNotFound
	}
	return foundEl, nil
}

// Add unique values to dictrionary
func (d Dictionary) Add(key, value string) error {
	_, ok := d[key]

	if ok {
		return ErrDuplicateKey
	}

	d[key] = value
	return nil
}
