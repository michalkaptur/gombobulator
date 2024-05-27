package main

import (
	"testing"
)

func TestAddition(t *testing.T) {
	data := []int{2, 2}
	got := add(data)
	want := 4
	if got != want {
		t.Errorf("got %v but wanted %v", got, want)
	}
}
