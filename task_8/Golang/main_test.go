package main

import (
	"testing"
	"fmt"
)

func TestHandleString(t *testing.T) {
	checkString("\"\"", 2, 0, t)
}

func checkString(inputString string, charsOfCode int, charsOfValue int, t *testing.T) {
	resultChan := make(chan AnalyseResult, 1)

	handleString("\"\"", resultChan)
	emptyResult := <- resultChan

	if emptyResult.charsOfCode != charsOfCode {
		t.Error(fmt.Sprintf("charsOfCode equals %s, but must be %s", emptyResult.charsOfCode, charsOfCode))
	} else if emptyResult.charsOfValue != charsOfValue {
		t.Error(fmt.Sprintf("charsOfValue equals %s, but must be %s", emptyResult.charsOfValue, charsOfValue))
	}
}
