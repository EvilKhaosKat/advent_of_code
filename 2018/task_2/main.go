package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	input := readInput()

	resultFirst := solveFirst(input)
	fmt.Println(resultFirst)

	resultSecond := solveSecond(input)
	fmt.Println(resultSecond)
}

func solveFirst(input[] string) int {
	var totalTwoCount, totalThreeCount int

	for _, inputLine := range input {
		twoCount, threeCount := parseId(inputLine)
		totalTwoCount += twoCount
		totalThreeCount += threeCount
	}

	return totalTwoCount * totalThreeCount
}

func solveSecond(input[] string) string {
	for lineNum, inputLine := range input {
		for i := lineNum + 1; i < len(input); i++ {
			if suitable, diff := suitableDifference(inputLine, input[i]); suitable {
				return diff
			}
		}
	}

	panic("not found")
}


func suitableDifference(first string, second string) (bool, string) {
	diffPosition := -1

	for charNum, char := range first {
		if char != []rune(second)[charNum] {
			if diffPosition != -1 {
				return false, ""
			}

			diffPosition = charNum
		}
	}

	var result strings.Builder
	for charNum, char := range first {
		if charNum == diffPosition {
			continue
		}

		result.WriteRune(char)
	}

	return true, result.String()
}

func parseId(id string) (int, int) {
	letterCount := make(map[rune]int)

	for _, c := range id {
		letterCount[c] = letterCount[c] + 1
	}

	var twoCount, threeCount int
	for _, count := range letterCount {
		if count == 2 {
			twoCount = 1
		} else if count == 3 {
				threeCount = 1
		}
	}

	return twoCount, threeCount
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
