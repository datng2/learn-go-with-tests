package main

import (
	"bytes"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	buf := bytes.Buffer{}

	Countdown(3, time.Duration(1*time.Microsecond), &buf)

	got := buf.String()
	want := "3\n2\n1\nGo!"
	if got != want {
		t.Errorf(`
		Expected:
		%q
		Got:
		%q`, want, got)
	}
}
