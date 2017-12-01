package main

import (
	"testing"
)

func Test1111(t *testing.T) {
	input := "1111"
	result := solve(input)

	if result != 4 {
		t.Error("sum must be 4, but:", result)
	}
}

func Test1122(t *testing.T) {
	input := "1122"
	result := solve(input)

	if result != 3 {
		t.Error("sum must be 3, but:", result)
	}
}

func Test1234(t *testing.T) {
	input := "1234"
	result := solve(input)

	if result != 0 {
		t.Error("sum must be 0, but:", result)
	}
}

func Test91212129(t *testing.T) {
	input := "91212129"
	result := solve(input)

	if result != 9 {
		t.Error("sum must be 9, but:", result)
	}
}