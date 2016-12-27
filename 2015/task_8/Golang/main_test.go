package main

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
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

func TestOnInput(t *testing.T) {
	result := calculateResult()

	assert.Equal(t, result.encodedResult, CORRECT_ENCODED_RESULT, fmt.Sprintf("First part result should be %d", CORRECT_ENCODED_RESULT))
	assert.Equal(t, result.escapedResult, CORRECT_ESCAPED_RESULT, fmt.Sprintf("Second part result should be %d", CORRECT_ESCAPED_RESULT))
}

func checkString(inputString string, charsOfCode int, charsOfValue int, totalEncodedChars int, t *testing.T) {
	result := getAnalyseResult(inputString)

	assert.Equal(t, result.charsOfCode, charsOfCode, fmt.Sprintf("charsOfCode should be %d", charsOfCode))
	assert.Equal(t, result.charsOfValue, charsOfValue, fmt.Sprintf("charsOfValue should be %d", charsOfValue))
	assert.Equal(t, result.totalEncodedChars, totalEncodedChars, fmt.Sprintf("totalEncodedChars should be %d", totalEncodedChars))
}

func getAnalyseResult(inputString string) AnalyseResult {
	resultChan := make(chan AnalyseResult, 1)

	handleString(inputString, resultChan)
	return <-resultChan
}
