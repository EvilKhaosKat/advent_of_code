package main

import (
	"os"
	"log"
	"bufio"
	"strconv"
	"fmt"
	"strings"
)

func main() {
	input := readInput()
	parsedInput := parseInput(input)

	fmt.Println(parsedInput)
}

func parseInput(input []string) [][]int {
	var result [][]int

	for _, inputLine := range input {
		var line []int

		rawNumbers := strings.Split(inputLine, "\t")
		for _, rawNumber := range rawNumbers {
			line = append(line, getNumByString(rawNumber))
		}

		result = append(result, line)
	}

	return result
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
