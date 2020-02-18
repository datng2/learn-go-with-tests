package dependencyinjection

import (
	"fmt"
	"io"
	"os"
)

// What exactly is an os.Stdout? What does Fprintf expect to get passed to it for the 1st argument?
//
// Answer: io.Writer - a great general purpose interface for "put this data somewhere".
//
// Bad:
//     func update(writer *bytes.Buffer, data string) // Buffer implements io.Writer interface
// Better
//     func update(writer io.Writer, data string)     // can pass different types of data

func printHello(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello %s", name)
}

func main() {
	printHello(os.Stdout, "world")
}
