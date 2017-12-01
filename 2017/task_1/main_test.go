package main

import (
	"testing"
	"fmt"
)

func TestFirst1122(t *testing.T) {
	input := "1122"
	result := solveFirst(input)

	checkResult(t, result, 3)
}

func TestFirst1234(t *testing.T) {
	input := "1234"
	result := solveFirst(input)

	checkResult(t, result, 0)
}

func TestFirst91212129(t *testing.T) {
	input := "91212129"
	result := solveFirst(input)

	checkResult(t, result, 9)
}

func TestFirst1111(t *testing.T) {
	input := "1111"
	result := solveFirst(input)

	checkResult(t, result, 4)
}

func TestSecond1212(t *testing.T) {
	input := "1212"
	result := solveSecond(input)

	checkResult(t, result, 6)
}

func TestSecond1221(t *testing.T) {
	input := "1221"
	result := solveSecond(input)

	checkResult(t, result, 0)
}

func TestSecond123425(t *testing.T) {
	input := "123425"
	result := solveSecond(input)

	checkResult(t, result, 4)
}

func TestSecond123123(t *testing.T) {
	input := "123123"
	result := solveSecond(input)

	checkResult(t, result, 12)
}

func TestSecond12131415(t *testing.T) {
	input := "12131415"
	result := solveSecond(input)

	checkResult(t, result, 4)
}

func checkResult(t *testing.T, actualResult int, requiredResult int) {
	t.Helper()

	if actualResult != requiredResult {
		t.Error(fmt.Printf("sum must be %+v, but: %+v", requiredResult, actualResult))
	}
}