package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input := readInput()

	resultFirst := solveFirst(input)
	fmt.Println(resultFirst)

	//resultSecond := solveSecond(input)
	//fmt.Println(resultSecond)
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
