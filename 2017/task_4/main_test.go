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

func TestSecondPartFirstPassphrase(t *testing.T) {
	result := solveSecond([]string{"abcde fghij"})

	checkResult(t, result, 1)
}

func TestSecondPartSecondPassphrase(t *testing.T) {
	result := solveSecond([]string{"abcde xyz ecdab"})

	checkResult(t, result, 0)
}

func TestSecondPartThirdPassphrase(t *testing.T) {
	result := solveSecond([]string{"a ab abc abd abf abj"})

	checkResult(t, result, 1)
}

func TestSecondPartFourthPassphrase(t *testing.T) {
	result := solveSecond([]string{"iiii oiii ooii oooi oooo"})

	checkResult(t, result, 1)
}

func TestSecondPartFifthPassphrase(t *testing.T) {
	result := solveSecond([]string{"oiii ioii iioi iiio"})

	checkResult(t, result, 0)
}

func checkResult(t *testing.T, actualResult int, requiredResult int) {
	t.Helper()

	if actualResult != requiredResult {
		t.Error(fmt.Printf("correct passphrases count must be %+v, but: %+v", requiredResult, actualResult))
	}
}