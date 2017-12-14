package main

import (
	"testing"
	"fmt"
)

func TestListCreation(t *testing.T) {
	list := createList(5)

	for index, value := range list {
		if index != value {
			t.Error(fmt.Printf("value %v != index %v, list:%v\n", value, index, list))
			t.Fail()
		}
	}
}

func TestFirst1(t *testing.T) {
	result := solveFirst(5, []int{3, 4, 1, 5})

	checkResultInt(t, result, 12)
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
