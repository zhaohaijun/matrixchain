package main

import "testing"

func TestAdd(t *testing.T) {
	sum := Add(1, 2)
	if sum != 3 {
		t.Error("Add 1 and 2 result isn't 3")
	}
}
func TestMultiAdd(t *testing.T) {
	var tests = []struct {
		Arg1 int
		Arg2 int
		Sum  int
	}{
		{1, 1, 2},
		{-1, -1, -2},
		{1, -1, 0},
		{0, 1, 1},
	}
	for _, test := range tests {
		sum := Add(test.Arg1, test.Arg2)
		if sum != test.Sum {
			t.Errorf("Add %v adn %v result isn't %v", test.Arg1, test.Arg2, test.Sum)
		}
	}
}
