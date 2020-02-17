package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("Adds two number", func(t *testing.T) {
		got := Add(1, 2)
		want := 3
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
