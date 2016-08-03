package main

import (
	"testing"
	"fmt"
	. "github.com/franela/goblin"
)

const (
	CORRECT_ENCODED_RESULT = 1342
	CORRECT_ESCAPED_RESULT = 2074
)

func TestStringDecoding(t *testing.T) {
	checkString(`""`, 2, 0, 6, t)
	checkString(`"abc"`, 5, 3, 9, t)
	checkString(`"aaa\"aaa"`, 10, 7, 16, t)
	checkString(`"\x27"`, 6, 1, 11, t)
}

func checkString(inputString string, charsOfCode int, charsOfValue int, totalEncodedChars int, t *testing.T) {
	resultChan := make(chan AnalyseResult, 1)

	handleString(inputString, resultChan)
	result := <-resultChan

	g := Goblin(t)
	g.Describe(fmt.Sprintf("Sanity checks with %s", inputString), func() {
		g.It(fmt.Sprintf("charsOfCode should be %d", charsOfCode), func() {
			g.Assert(result.charsOfCode).Equal(charsOfCode)
		})

		g.It(fmt.Sprintf("charsOfValue should be %d", charsOfValue), func() {
			g.Assert(result.charsOfValue).Equal(charsOfValue)
		})

		g.It(fmt.Sprintf("totalEncodedChars should be %d", totalEncodedChars), func() {
			g.Assert(result.totalEncodedChars).Equal(totalEncodedChars)
		})
	})
}

func TestWithGoblin(t *testing.T) {
	result := calculateResult()

	g := Goblin(t)
	g.Describe("Check with input ", func() {
		g.It(fmt.Sprintf("First part result should be %d", CORRECT_ENCODED_RESULT), func() {
			g.Assert(result.encodedResult).Equal(CORRECT_ENCODED_RESULT)
		})
		g.It(fmt.Sprintf("Second part result should be %d", CORRECT_ESCAPED_RESULT), func() {
			g.Assert(result.escapedResult).Equal(CORRECT_ESCAPED_RESULT)
		})
	})
}