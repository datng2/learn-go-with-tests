package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	Name   string
	Input  interface{}
	Output []string
}

func TestWalk(t *testing.T) {
	testCases := []TestCase{
		{
			Name:   "Struct with no string field",
			Input:  struct{ Age int }{1},
			Output: []string{},
		},
		{
			Name:   "Struct with no string field",
			Input:  "dat",
			Output: []string{"dat"},
		},
		{
			Name:   "Struct with one string field",
			Input:  Simple{1, "A"},
			Output: []string{"A"},
		},
		{
			Name:   "Pointer to simple struct",
			Input:  &Simple{1, "A"},
			Output: []string{"A"},
		},
		{
			Name:   "Array of struct",
			Input:  [2]Simple{{1, "A"}, {1, "B"}},
			Output: []string{"A", "B"},
		},
		{
			Name:   "Struct with multiple string fields and array",
			Input:  Tricky{1, "A", "B", []string{"C", "D"}, map[int]string{4: "E"}},
			Output: []string{"A", "B", "C", "D", "E"},
		},
	}

	for _, test := range testCases {
		got := []string{}
		walk(test.Input, func(x string) {
			got = append(got, x)
		})

		if !reflect.DeepEqual(got, test.Output) {
			t.Errorf("Testcase %q:\ngot %v, want %v", test.Name, got, test.Output)
		}
	}

}
func assertContains(t *testing.T, haystack []string, needle string) {
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}

type Simple struct {
	Age  int
	Name string
}

type Tricky struct {
	Age             int
	Name            string
	CurrAddress     string
	PrevAddresses   []string
	AcademicRecords map[int]string // schoolID - school
}
