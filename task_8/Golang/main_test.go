package main

import (
	"testing"
	"fmt"
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
