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
	parsedInput := parseInput(input)

	firstResult := solveFirst(parsedInput)

	fmt.Println(firstResult)
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
	offset := 0

	offsetsListLen := len(parsedInput)

	for {
		if offset >= offsetsListLen {
			break
		}

		num := parsedInput[offset]

		parsedInput[offset] = num + 1
		offset = offset + num

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
