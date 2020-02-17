package arrays

import (
	"reflect"
	"testing"
)

func TestSumOnArray(t *testing.T) {

	numberArray := [5]int{1, 2}

	got := Sum(numberArray[:])
	want := 3

	if got != want {
		t.Errorf("got %d want %d, give %v", got, want, numberArray)
	}

}

func TestSumOnSlice(t *testing.T) {
	numberSlice := []int{1, 2}
	got := Sum(numberSlice)
	want := 3

	if got != want {
		t.Errorf("got %d want %d, give %v", got, want, numberSlice)
	}
}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	got := SumAll([]int{}, []int{0, 9})
	want := []int{0, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}

}
