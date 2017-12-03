package main

import (
	"testing"
	"fmt"
)

func TestFirst(t *testing.T) {
	input := [][]int {
		{5, 1, 9, 5},
		{7, 5, 3},
		{2, 4, 6, 8},
	}
	result := solveFirst(input)

	checkResult(t, result, 18)
}

func TestSecond(t *testing.T) {
	input := [][]int {
		{5, 9, 2, 8},
		{9, 4, 7, 3},
		{3, 8, 6, 5},
		}
	result := solveSecond(input)

	checkResult(t, result, 9)
}

func checkResult(t *testing.T, actualResult int, requiredResult int) {
	t.Helper()

	if actualResult != requiredResult {
		t.Error(fmt.Printf("checksum must be %+v, but: %+v", requiredResult, actualResult))
	}
}