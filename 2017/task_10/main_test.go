package main

import (
	"testing"
	"fmt"
	"reflect"
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

func TestSecondParseInput(t *testing.T) {
	result := parseInputSecond("1,2,3")

	requiredResult := []int{49, 44, 50, 44, 51, 17, 31, 73, 47, 23}
	if !reflect.DeepEqual(result, requiredResult) {
		t.Error(fmt.Printf("actual %v differs from required result %v\n", result, requiredResult))
		t.Fail()
	}
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
