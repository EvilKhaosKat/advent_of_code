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

func TestSecondSplitBy16(t *testing.T) {
	list := createList(256)

	partsBy16 := splitHashBy16(list)

	partsLen := len(partsBy16)
	if partsLen != 16 {
		t.Error(fmt.Printf("Required 16 parts, but got %v", partsLen))
		t.Fail()
	}

	firstFirstValue := partsBy16[0][0]
	if firstFirstValue != 0 {
		t.Error(fmt.Printf("[0][0] should be 0, but %v", firstFirstValue))
		fmt.Printf("\n%v", partsBy16)
		t.Fail()
	}

	secondFirstValue := partsBy16[1][0]
	if secondFirstValue != 16 {
		t.Error(fmt.Printf("[1][0] should be 16, but %v", secondFirstValue))
		fmt.Printf("\n%v", partsBy16)
		t.Fail()
	}
}

func TestSecondXor(t *testing.T) {
	xor := getXoredPart([]int{65, 27, 9, 1, 4, 3, 40, 50, 91, 7, 6, 0, 2, 5, 68, 22})

	checkResultInt(t, xor, 64)
}

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
