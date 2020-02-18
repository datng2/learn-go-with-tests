package maps

type DictionaryService interface {
	Add(word string, definition string) error
	Update(word string, newDefinition string) error
	Search(word string) (string, error)
	Delete(word string) error
}
