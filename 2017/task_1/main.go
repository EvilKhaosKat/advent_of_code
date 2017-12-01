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
	result := solve(input)

	fmt.Println(result)
}

func solve(input string) int {
	sum := 0
	lastNum := getLastNum(input)
	for _, c := range input {
		currNum := getNumByChar(string(c))

		if lastNum == currNum {
			sum += currNum
		}

		lastNum = currNum
	}

	return sum
}

func getLastNum(input string) int {
	lastChar := string(input[utf8.RuneCountInString(input) - 1])

	return getNumByChar(lastChar)
}
func getNumByChar(lastChar string) int {
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
