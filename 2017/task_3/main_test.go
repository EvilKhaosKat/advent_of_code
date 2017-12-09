package main

import (
	"testing"
	"fmt"
)

func TestFirst12(t *testing.T) {
	result := solveFirst(12)

	checkResult(t, result, 3)
}

func TestFirst23(t *testing.T) {
	result := solveFirst(23)

	checkResult(t, result, 2)
}

func TestFirst1024(t *testing.T) {
	result := solveFirst(1024)

	checkResult(t, result, 31)
}

func TestFirst31(t *testing.T) {
	result := solveFirst(31)

	checkResult(t, result, 6)
}

func TestSecond23(t *testing.T) {
	result := solveSecond(23)

	checkResult(t, result, 25)
}

func TestSecond133(t *testing.T) {
	result := solveSecond(133)

	checkResult(t, result, 142)
}


func checkResult(t *testing.T, actualResult int, requiredResult int) {
	t.Helper()

	if actualResult != requiredResult {
		t.Error(fmt.Printf("steps must be %+v, but: %+v", requiredResult, actualResult))
	}
}