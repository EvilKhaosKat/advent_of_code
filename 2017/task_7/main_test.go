package main

import (
	"testing"
	"fmt"
)

func TestFirst(t *testing.T) {
	result := solveFirst([]string{
		"pbga (66)",
		"xhth (57)",
		"ebii (61)",
		"havc (66)",
		"ktlj (57)",
		"fwft (72) -> ktlj, cntj, xhth",
		"qoyq (66)",
		"padx (45) -> pbga, havc, qoyq",
		"tknk (41) -> ugml, padx, fwft",
		"jptl (61)",
		"ugml (68) -> gyxo, ebii, jptl",
		"gyxo (61)",
		"cntj (57)",
	})

	checkResult(t, result, "tknk")
}

func checkResult(t *testing.T, actualResult string, requiredResult string) {
	t.Helper()

	if actualResult != requiredResult {
		t.Error(fmt.Printf("top node name must be %+v, but: %+v", requiredResult, actualResult))
	}
}
