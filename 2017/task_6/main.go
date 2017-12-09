package main

import (
	"os"
	"log"
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"bytes"
)

func main() {
	input := readInputSingleLine()
	banks := parseInput(input)

	firstResult := solveFirst(banks)
	secondResult := solveSecond(banks)

	fmt.Println(firstResult)
	fmt.Println(secondResult)
}

func solveFirst(banks []int) int {
	stepsCount := 0

	states := map[string]bool{}

	for {
		maxIndex, maxBanks := getMaxNum(banks)

		blocksToSpend := maxBanks
		banks[maxIndex] = 0

		redistributeBanks(maxIndex, blocksToSpend, banks)

		stepsCount++

		currentState := getCurrentState(banks)
		if _, present := states[currentState]; present {
			break
		} else {
			states[currentState] = true
		}
	}

	return stepsCount
}

func solveSecond(banks []int) int {
	stepsCount := 0
	sameStateStepsCount := 0
	present := false

	states := map[string]int{}

	for {
		maxIndex, maxBanks := getMaxNum(banks)

		blocksToSpend := maxBanks
		banks[maxIndex] = 0

		redistributeBanks(maxIndex, blocksToSpend, banks)

		stepsCount++

		currentState := getCurrentState(banks)
		sameStateStepsCount, present = states[currentState];
		if present {
			break
		} else {
			states[currentState] = stepsCount
		}
	}

	return stepsCount - sameStateStepsCount
}

func redistributeBanks(maxIndex int, blocksToSpend int, banks []int) {
	banksCount := len(banks)

	index := maxIndex + 1
	for {
		if blocksToSpend == 0 {
			break
		}

		index = index % banksCount

		banks[index]++
		blocksToSpend--

		index++
	}
}

func getCurrentState(banks []int) string {
	var state bytes.Buffer

	for i := 0; i < len(banks); i++ {
		state.WriteString(strconv.Itoa(banks[i]))
		state.WriteString(" ")
	}

	return state.String()
}

func getMaxNum(numbers []int) (index int, maxNum int) {
	for i := 0; i < len(numbers); i++ {
		num := numbers[i]
		if num > maxNum {
			maxNum = num
			index = i
		}
	}

	return index, maxNum
}

func parseInput(input string) []int {
	var parsedInput []int

	for _, numString := range strings.Split(input, "\t") {
		parsedInput = append(parsedInput, getNumByString(numString))
	}

	return parsedInput
}

func getNumByString(numRaw string) int {
	num, err := strconv.Atoi(numRaw)
	if err != nil {
		panic("Cannot get num:" + err.Error())
	}
	return num
}

func readInputSingleLine() string {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		return scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	panic("Nothing found/read from input")
}

