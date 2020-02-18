package maps

import "testing"

func TestDictionary(t *testing.T) {
	t.Run("Search", func(t *testing.T) {
		dict := Dictionary{"test": "this is a test"}

		got, _ := dict.Search("test")
		assertString(t, got, "this is a test")

		_, gotError := dict.Search("unknown")
		assertError(t, gotError, ErrDefinitionNotFound)
	})

	t.Run("Add", func(t *testing.T) {
		dict := Dictionary{}
		ok := dict.Add("test", "this is a test")
		assertError(t, ok, nil)

		err := dict.Add("test", "new test")
		assertError(t, err, ErrWordExists)

	})

	t.Run("Update", func(t *testing.T) {
		dict := Dictionary{"test": "this is a test"}
		ok := dict.Update("test", "new definition")
		assertError(t, ok, nil)

		err := dict.Update("aaa", "aaa")
		assertError(t, err, ErrWordDoesNotExist)

	})

	t.Run("Delete", func(t *testing.T) {
		dict := Dictionary{"test": "this is a test"}
		ok := dict.Delete("test")
		assertError(t, ok, nil)

		err := dict.Delete("test")
		assertError(t, err, ErrWordDoesNotExist)

	})

}

func assertString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q but want %q", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("Expected error %q, but got %q", got, want)
	}
}
