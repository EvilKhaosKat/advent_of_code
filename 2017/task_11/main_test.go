package main

import (
	"testing"
	"fmt"
)

func TestFirst1(t *testing.T) {
	result := solveFirst("ne,ne,ne")
	checkResultInt(t, result, 3)
}

func TestFirst2(t *testing.T) {
	result := solveFirst("ne,ne,sw,sw")
	checkResultInt(t, result, 0)
}

func TestFirst3(t *testing.T) {
	result := solveFirst("ne,ne,s,s")
	checkResultInt(t, result, 2)
}

func TestFirst4(t *testing.T) {
	result := solveFirst("se,sw,se,sw,sw")
	checkResultInt(t, result, 3)
}

//
//helper functions
//
func checkResultString(t *testing.T, actualResult string, requiredResult string) {
	t.Helper()

	if actualResult != requiredResult {
		t.Error(fmt.Printf("required value must be %+v, but: %+v", requiredResult, actualResult))
	}
}

func checkResultInt(t *testing.T, actualResult int, requiredResult int) {
	t.Helper()

	if actualResult != requiredResult {
		t.Error(fmt.Printf("required value must be %+v, but: %+v", requiredResult, actualResult))
	}
}
