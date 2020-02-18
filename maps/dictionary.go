package maps

// Dictionary maps a word to its definition
type Dictionary map[string]string

// Why to use constant errors
// https://dave.cheney.net/2016/04/07/constant-errors
const (
	ErrDefinitionNotFound = DictErr("Undefined word")
	ErrWordExists         = DictErr("Word is existed, use Update instead.")
	ErrWordDoesNotExist   = DictErr("World does not exist, cannot update.")
)

type DictErr string

func (e DictErr) Error() string { return string(e) }

// Search returns a word's definition
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrDefinitionNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, ok := d.Search(word)
	if ok == nil { // found existing def for this word
		return ErrWordExists
	}

	d[word] = definition
	return nil
}

func (d Dictionary) Update(word, newDefinition string) error {
	_, ok := d.Search(word)
	if ok == ErrDefinitionNotFound {
		return ErrWordDoesNotExist
	}

	d[word] = newDefinition
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, ok := d.Search(word)
	if ok == ErrDefinitionNotFound {
		return ErrWordDoesNotExist
	}

	// built-in feature of Go to delete an entry from map
	delete(d, word)
	return nil
}
