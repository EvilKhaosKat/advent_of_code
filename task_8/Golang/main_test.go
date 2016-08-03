package main

import (
	"testing"
	"fmt"
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

func TestTaskInputResults(t *testing.T) {
	result := calculateResult()

	if result.encodedResult != CORRECT_ENCODED_RESULT {
		t.Fatal(fmt.Sprintf("result.encodedResult equals %d, but must be %d",
			result.encodedResult, CORRECT_ENCODED_RESULT))
	}

	if result.escapedResult != CORRECT_ESCAPED_RESULT {
		t.Fatal(fmt.Sprintf("result.escapedResult equals %d, but must be %d",
			result.escapedResult, CORRECT_ESCAPED_RESULT))
	}
}

func checkString(inputString string, charsOfCode int, charsOfValue int, totalEncodedChars int, t *testing.T) {
	resultChan := make(chan AnalyseResult, 1)

	handleString(inputString, resultChan)
	result := <-resultChan

	if result.charsOfCode != charsOfCode {
		t.Fatal(fmt.Sprintf("Fail for %s, charsOfCode equals %d, but must be %d",
			result.inputString, result.charsOfCode, charsOfCode))
	} else if result.charsOfValue != charsOfValue {
		t.Fatal(fmt.Sprintf("Fail for %s, charsOfValue equals %d, but must be %d",
			result.inputString, result.charsOfValue, charsOfValue))
	} else if result.totalEncodedChars != totalEncodedChars {
		t.Fatal(fmt.Sprintf("Fail for %s, totalEncodedChars equals %d, but must be %d",
			result.inputString, result.totalEncodedChars, totalEncodedChars))
	}
}