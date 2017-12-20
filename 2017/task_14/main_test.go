package main

import (
	"testing"
	"fmt"
)

func TestBinaryHash(t *testing.T) {
	binaryRepresenation := getBinaryKnotHash("a0c2017")

	checkResultString(t, binaryRepresenation, "1010000011000010000000010111")
}

func TestFirst(t *testing.T) {
	usedSquaresCount := solveFirst("flqrgnkx")

	checkResultInt(t, usedSquaresCount, 8108)
}

func TestSecond(t *testing.T) {
	usedSquaresCount := solveSecond("flqrgnkx")

	checkResultInt(t, usedSquaresCount, 1242)
}

//
//helper functions
//
func checkResultString(t *testing.T, actualResult string, requiredResult string) {
	t.Helper()

	if actualResult != requiredResult {
		t.Error(fmt.Printf("required value must be %+v, but: %+v\n", requiredResult, actualResult))
	}
}

func checkResultInt(t *testing.T, actualResult int, requiredResult int) {
	t.Helper()

	if actualResult != requiredResult {
		t.Error(fmt.Printf("required value must be %+v, but: %+v\n", requiredResult, actualResult))
	}
}
