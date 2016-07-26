package main

import (
	"testing"
	"fmt"
)

func TestHandleString(t *testing.T) {
	checkString("\"\"", 2, 0, t)
	checkString("\"abc\"", 5, 3, t)
	checkString("\"aaa\\\"aaa\"", 10, 7, t)
	checkString("\"\\x27\"", 6, 1, t)
}

func checkString(inputString string, charsOfCode int, charsOfValue int, t *testing.T) {
	resultChan := make(chan AnalyseResult, 1)

	handleString(inputString, resultChan)
	emptyResult := <- resultChan

	if emptyResult.charsOfCode != charsOfCode {
		t.Error(fmt.Sprintf("charsOfCode equals %d, but must be %d", emptyResult.charsOfCode, charsOfCode))
	} else if emptyResult.charsOfValue != charsOfValue {
		t.Error(fmt.Sprintf("charsOfValue equals %d, but must be %d", emptyResult.charsOfValue, charsOfValue))
	}
}
