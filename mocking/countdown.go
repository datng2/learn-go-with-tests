package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// =========
// Task
// ==========
//     You have been asked to write a program which counts down from 3, printing
// each number on a new line (with a 1 second pause) and when it reaches zero
// it will print "Go!" and exit.

func Countdown(start int, interval time.Duration, w io.Writer) {
	for i := start; i > 0; i-- {
		// fmt.Println(i)
		fmt.Fprintln(w, i) // for mocking, replace writer from stdout to buffer
		time.Sleep(interval)
	}
	fmt.Fprint(w, "Go!")
}

func main() {
	Countdown(3, time.Duration(1*time.Second), os.Stdout)
}
