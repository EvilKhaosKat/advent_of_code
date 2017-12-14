package main

import (
	"os"
	"log"
	"bufio"
	"fmt"
	"strconv"
)

type context struct {
	groupLevel int

	garbageActive bool
	ignoreActive bool

	score int
}

func main() {
	input := readInputSingleLine()

	firstResult := solveFirst(input)
	//secondResult := solveSecond(input)

	fmt.Println(firstResult)
	//fmt.Println(secondResult)
}

func solveFirst(input string) int {
	context := context{}

	for _, char := range input {
		if context.ignoreActive {
			context.ignoreActive = false
			continue
		}

		if context.garbageActive && char == '!' {
			context.ignoreActive = true
			continue
		}

		if context.garbageActive && char != '>' {
			continue
		}

		if context.garbageActive && char == '>' {
			context.garbageActive = false
			continue
		}

		if char == '<' {
			context.garbageActive = true
			continue
		}

		if char == '{' {
			context.groupLevel++
			continue
		}

		if char == '}' {
			context.score += context.groupLevel
			context.groupLevel--
			continue
		}
	}

	return context.score
}

//
//helper methods starts here
//
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

func readInputMultiLine() []string {
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


