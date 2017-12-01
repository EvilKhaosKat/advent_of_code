package main

import (
	"os"
	"log"
	"bufio"
	"unicode/utf8"
	"strconv"
	"fmt"
)

func main() {
	input := readInput()

	resultFirst := solveFirst(input)
	fmt.Println(resultFirst)

	resultSecond := solveSecond(input)
	fmt.Println(resultSecond)
}

func solveFirst(input string) int {
	sum := 0
	lastNum := getLastNum(input)
	for _, c := range input {
		currNum := getNumByString(string(c))

		if lastNum == currNum {
			sum += currNum
		}

		lastNum = currNum
	}

	return sum
}

func solveSecond(input string) int {
	sum := 0
	inputSize := utf8.RuneCountInString(input)
	jumpSize := inputSize/2

	for index, c := range input {
		currNum := getNumByString(string(c))
		nextNum := getNextNum(input, index, inputSize, jumpSize)

		if currNum == nextNum {
			sum += currNum
		}
	}

	return sum
}

func getNextNum(input string, index int, inputSize int, jumpSize int) int {
	resultIndex := index + jumpSize
	if resultIndex >= inputSize {
		resultIndex = resultIndex % inputSize
	}

	return getNumByString(string(input[resultIndex]))
}

func getLastNum(input string) int {
	lastChar := string(input[utf8.RuneCountInString(input) - 1])

	return getNumByString(lastChar)
}

func getNumByString(lastChar string) int {
	lastNum, err := strconv.Atoi(lastChar)
	if err != nil {
		panic("Cannot get last num:" + err.Error())
	}
	return lastNum
}

func readInput() string {
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
