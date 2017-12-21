package main

import (
	"testing"
	"fmt"
)

func TestGeneratorGenerationA(t *testing.T) {
	generatorA := &Generator{}
	results := generatorA.run(GENERATOR_A_FACTOR, 65, 5)

	requiredResults := map[int]int{
		0: 1092455,
		1: 1181022009,
		2: 245556042,
		3: 1744312007,
		4: 1352636452,
	}

	index := 0
	for value := range results {
		requiredValue := requiredResults[index]
		checkResultInt(t, value, requiredValue)

		index++
	}
}

func TestGeneratorGenerationB(t *testing.T) {
	generatorA := &Generator{}
	results := generatorA.run(GENERATOR_B_FACTOR, 8921, 5)

	requiredResults := map[int]int{
		0: 430625591,
		1: 1233683848,
		2: 1431495498,
		3: 137874439,
		4: 285222916,
	}

	index := 0
	for value := range results {
		requiredValue := requiredResults[index]
		checkResultInt(t, value, requiredValue)

		index++
	}
}

//
//helper functions
//
func checkResultString(t *testing.T, actualResult string, requiredResult string) {
	t.Helper()

	if actualResult != requiredResult {
		t.Error(fmt.Printf("required value must be %+v, but: %+v\n", requiredResult, actualResult))
	}
}

func checkResultInt(t *testing.T, actualResult int, requiredResult int) {
	t.Helper()

	if actualResult != requiredResult {
		t.Error(fmt.Printf("required value must be %+v, but: %+v\n", requiredResult, actualResult))
	}
}
