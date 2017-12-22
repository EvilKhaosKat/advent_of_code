package main

import (
	"testing"
	"fmt"
)

func TestSpin(t *testing.T) {
	programs := getProgramsByLine("abcde")

	result := spin(programs, 3)

	checkResultString(t, getStringByPrograms(result), "cdeab")
}

func TestExchange(t *testing.T) {
	programs := getProgramsByLine("eabcd")

	result := exchange(programs, "3/4")

	checkResultString(t, getStringByPrograms(result), "eabdc")
}

func TestPartner(t *testing.T) {
	programs := getProgramsByLine("eabdc")

	result := partner(programs, "e/b")

	checkResultString(t, getStringByPrograms(result), "baedc")
}

//
//helper functions
//
func checkResultString(t *testing.T, actualResult, requiredResult string) {
	t.Helper()

	if actualResult != requiredResult {
		t.Error(fmt.Printf("required value must be %+v, but: %+v\n", requiredResult, actualResult))
	}
}

func checkResultInt(t *testing.T, actualResult, requiredResult int) {
	t.Helper()

	if actualResult != requiredResult {
		t.Error(fmt.Printf("required value must be %+v, but: %+v\n", requiredResult, actualResult))
	}
}
