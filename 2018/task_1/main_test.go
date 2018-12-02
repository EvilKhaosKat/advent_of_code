package main

import (
	"fmt"
	"testing"
)

func TestFirst(t *testing.T) {
	input := []string {
		"1",
		"2",
		"-3",
		"42",
	}
	result := solveFirst(input)

	checkResult(t, result, 42)
}

func TestSecond(t *testing.T) {
	input := []string {
		"1",
		"2",
		"-3",
		"2",
	}
	result := solveSecond(input)

	checkResult(t, result, 3)
}

func checkResult(t *testing.T, actualResult int, requiredResult int) {
	t.Helper()

	if actualResult != requiredResult {
		t.Error(fmt.Printf("frequency must be %+v, but: %+v", requiredResult, actualResult))
	}
}