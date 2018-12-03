package main

import (
	"fmt"
	"testing"
)

func TestFirst(t *testing.T) {
	input := []string{
		"#1 @ 1,3: 4x4",
		"#2 @ 3,1: 4x4",
		"#3 @ 5,5: 2x2",
	}
	result := solveFirst(input)

	checkResult(t, result, 4)
}

func TestSecond(t *testing.T) {
	input := []string{
		"#1 @ 1,3: 4x4",
		"#2 @ 3,1: 4x4",
		"#3 @ 5,5: 2x2",
	}
	result := solveSecond(input)

	checkResult(t, result, 3)
}

func TestParseClaim(t *testing.T) {
	claim := parseClaim("#123 @ 3,2: 5x4")

	checkResult(t, claim.id, 123)
	checkResult(t, claim.topX, 3)
	checkResult(t, claim.topY, 2)
	checkResult(t, claim.width, 5)
	checkResult(t, claim.height, 4)
}

func checkResult(t *testing.T, actualResult interface{}, requiredResult interface{}) {
	t.Helper()

	if actualResult != requiredResult {
		t.Error(fmt.Printf("value must be %+v, but: %+v\n", requiredResult, actualResult))
	}
}
