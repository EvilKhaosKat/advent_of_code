package main

import (
	"testing"
	"fmt"
)

func TestFirst1(t *testing.T) {
	result := solveFirst("{}")

	checkResultInt(t, result, 1)
}

func TestFirst2(t *testing.T) {
	result := solveFirst("{{{}}}")

	checkResultInt(t, result, 6)
}

func TestFirst3(t *testing.T) {
	result := solveFirst("{{},{}}")

	checkResultInt(t, result, 5)
}

func TestFirst4(t *testing.T) {
	result := solveFirst("{{{},{},{{}}}}")

	checkResultInt(t, result, 16)
}

func TestFirst5(t *testing.T) {
	result := solveFirst("{<a>,<a>,<a>,<a>}")

	checkResultInt(t, result, 1)
}

func TestFirst6(t *testing.T) {
	result := solveFirst("{{<ab>},{<ab>},{<ab>},{<ab>}}")

	checkResultInt(t, result, 9)
}

func TestFirst7(t *testing.T) {
	result := solveFirst("{{<!!>},{<!!>},{<!!>},{<!!>}}")

	checkResultInt(t, result, 9)
}

func TestFirst8(t *testing.T) {
	result := solveFirst("{{<a!>},{<a!>},{<a!>},{<ab>}}")

	checkResultInt(t, result, 3)
}

func TestSecond1(t *testing.T) {
	result := solveSecond("<>")

	checkResultInt(t, result, 0)
}

func TestSecond2(t *testing.T) {
	result := solveSecond("<random characters>")

	checkResultInt(t, result, 17)
}

func TestSecond3(t *testing.T) {
	result := solveSecond("<<<<>")

	checkResultInt(t, result, 3)
}

func TestSecond4(t *testing.T) {
	result := solveSecond("<{!>}>")

	checkResultInt(t, result, 2)
}

func TestSecond5(t *testing.T) {
	result := solveSecond("<!!>")

	checkResultInt(t, result, 0)
}

func TestSecond6(t *testing.T) {
	result := solveSecond("<!!!>>")

	checkResultInt(t, result, 0)
}

func TestSecond7(t *testing.T) {
	result := solveSecond("<{o\"i!a,<{i<a>")

	checkResultInt(t, result, 10)
}

func checkResultString(t *testing.T, actualResult string, requiredResult string) {
	t.Helper()

	if actualResult != requiredResult {
		t.Error(fmt.Printf("top node name must be %+v, but: %+v", requiredResult, actualResult))
	}
}

func checkResultInt(t *testing.T, actualResult int, requiredResult int) {
	t.Helper()

	if actualResult != requiredResult {
		t.Error(fmt.Printf("required weight must be %+v, but: %+v", requiredResult, actualResult))
	}
}
