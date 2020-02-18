package maps

import "testing"

func TestDictionary(t *testing.T) {
	t.Run("Search Dictionary", func(t *testing.T) {
		dict := Dictionary{"test": "this is a test"}

		got, _ := dict.Search("test")
		assertString(t, got, "this is a test")

		_, gotError := dict.Search("unknown")
		assertError(t, gotError, ErrDefinitionNotFound)
	})

	t.Run("Add to Dictionary", func(t *testing.T) {
		dict := Dictionary{}
		ok := dict.Add("test", "this is a test")
		assertError(t, ok, nil)

		err := dict.Add("test", "new test")
		assertError(t, err, ErrWordExists)

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
