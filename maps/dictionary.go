package maps

var (
	ErrKeyNotFound  = DictionaryError("There is no such key in dictionary")
	ErrDuplicateKey = DictionaryError("Key already exist")
)

type DictionaryError string

func (e DictionaryError) Error() string {
	return string(e)
}

// Dictionary stores key:value map of strings
type Dictionary map[string]string

// Search for key in dictionary
func (d Dictionary) Search(key string) (string, error) {
	foundEl, ok := d[key]

	if !ok {
		return "", ErrKeyNotFound
	}
	return foundEl, nil
}

// Add unique values to dictionary
func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrKeyNotFound:
		d[key] = value
	case nil:
		return ErrDuplicateKey
	default:
		return err
	}

	return nil
}
