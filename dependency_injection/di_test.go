package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestPrint(t *testing.T) {
	// Rather than using os.Stdout, we'd use buffer (that implements io.Writer)
	// Why?
	//   For UT, it's tricky to check what comes out to stdout.
	//   By using a bytes.Buffer, we are "forward" the output data for validation
	// For example:
	//
	//   func print(w io.Writer, message) // abstract the actual write logic
	//
	//   In Main: print(os.Stdout, message)
	//   In UT:   print(buffer, message)
	buffer := bytes.Buffer{}
	name := "world"

	printHello(&buffer, name)
	if buffer.String() != "Hello world" {
		t.Errorf("expected %q but got %q", buffer.String(), name)
	}
}
