package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input := readInput()

	resultFirst := solveFirst(input)
	fmt.Println(resultFirst)

	resultSecond := solveSecond(input)
	fmt.Println(resultSecond)
}

func solveFirst(input[] string) int {
	freq := 0

	for _, inputLine := range input {
		freq += getNumByString(inputLine)
	}

	return freq
}

func solveSecond(input[] string) int {
	freq := 0
	freqMet := make(map[int]bool)

	for {
		for _, inputLine := range input {
			freq += getNumByString(inputLine)
			if freqMet[freq] {
				return freq
			} else {
				freqMet[freq] = true
			}
		}
	}
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
