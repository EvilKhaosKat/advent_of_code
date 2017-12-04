package main

import (
	"testing"
	"fmt"
)

func TestFirstPartFirstPassphrase(t *testing.T) {
	result := solveFirst([]string{"aa bb cc dd ee"})

	checkResult(t, result, 1)
}

func TestFirstPartSecondPassphrase(t *testing.T) {
	result := solveFirst([]string{"aa bb cc dd aa"})

	checkResult(t, result, 0)
}

func TestFirstPartThirdPassphrase(t *testing.T) {
	result := solveFirst([]string{"aa bb cc dd aaa"})

	checkResult(t, result, 1)
}

func checkResult(t *testing.T, actualResult int, requiredResult int) {
	t.Helper()

	if actualResult != requiredResult {
		t.Error(fmt.Printf("correct passphrases count must be %+v, but: %+v", requiredResult, actualResult))
	}
}