package main

import (
	"testing"
	"fmt"
)

func TestLayersCreation(t *testing.T) {
	layers := initLayers([]string{
		"0: 3",
		"1: 2",
		"4: 4",
		"6: 4",
	})

	checkResultInt(t, len(layers), 7)
}

func TestFirst(t *testing.T) {
	severity, _ := solveFirst([]string{
		"0: 3",
		"1: 2",
		"4: 4",
		"6: 4",
	}, 0)

	checkResultInt(t, severity, 24)
}

func TestSecond(t *testing.T) {
	delayTick := solveSecond([]string{
		"0: 3",
		"1: 2",
		"4: 4",
		"6: 4",
	})

	checkResultInt(t, delayTick, 10)
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