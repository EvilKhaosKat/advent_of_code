package main

import (
	"testing"
	"fmt"
)

func Test12(t *testing.T) {
	result := solveFirst(12)

	checkResult(t, result, 3)
}

func Test23(t *testing.T) {
	result := solveFirst(23)

	checkResult(t, result, 2)
}

func Test1024(t *testing.T) {
	result := solveFirst(1024)

	checkResult(t, result, 31)
}


func checkResult(t *testing.T, actualResult int, requiredResult int) {
	t.Helper()

	if actualResult != requiredResult {
		t.Error(fmt.Printf("steps must be %+v, but: %+v", requiredResult, actualResult))
	}
}