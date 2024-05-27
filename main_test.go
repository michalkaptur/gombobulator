package main

import (
	"testing"
)

var addition_tests = []struct {
	in  []int
	out int
}{
	{[]int{2, 2}, 4},
	{[]int{1}, 1},
	{[]int{1, 1, 1, 1, 1, 1}, 6},
	{[]int{}, 0},
	{[]int{0}, 0},
	{[]int{10000000, 1}, 10000001},
}

func TestAddition(t *testing.T) {
	for _, pair := range addition_tests {
		got := add(pair.in)
		want := pair.out
		if got != want {
			t.Errorf("got %v but wanted %v", got, want)
		}

	}
}
