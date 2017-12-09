package main

import (
	"os"
	"log"
	"bufio"
	"fmt"
	"strconv"
)

func main() {
	input := readInput()

	firstResult := solveFirst(parseInput(input))
	secondResult := solveSecond(parseInput(input))

	fmt.Println(firstResult)
	fmt.Println(secondResult)
}

func parseInput(input []string) []int {
	parsedInput := make([]int, len(input))

	for index, inputLine := range input {
		parsedInput[index] = getNumByString(inputLine)
	}

	return parsedInput
}

func solveFirst(parsedInput []int) int {
	stepsCount := 0
	currentOffset := 0

	offsetsListLen := len(parsedInput)

	for {
		if currentOffset >= offsetsListLen || currentOffset < 0 {
			break
		}

		jumpValue := parsedInput[currentOffset]

		parsedInput[currentOffset] = jumpValue + 1
		currentOffset = currentOffset + jumpValue

		stepsCount++
	}

	return stepsCount
}

func solveSecond(parsedInput []int) int {
	stepsCount := 0
	currentOffset := 0

	offsetsListLen := len(parsedInput)

	for {
		if currentOffset >= offsetsListLen || currentOffset < 0 {
			break
		}

		jumpValue := parsedInput[currentOffset]

		if jumpValue >= 3 {
			parsedInput[currentOffset] = jumpValue - 1
		} else {
			parsedInput[currentOffset] = jumpValue + 1
		}

		currentOffset = currentOffset + jumpValue

		stepsCount++
	}

	return stepsCount
}

func getNumByString(numRaw string) int {
	num, err := strconv.Atoi(numRaw)
	if err != nil {
		panic("Cannot get num:" + err.Error())
	}
	return num
}

func readInput() []string {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var result []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
