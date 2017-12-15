package main

import (
	"os"
	"log"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := readInputSingleLine()

	firstResult := solveFirst(input)
	fmt.Println(firstResult)

}

func solveFirst(input string) int {
	x := 0
	y := 0
	z := 0

	for _, direction := range strings.Split(input, ",") {
		if direction == "n" {
			y++
			z--
		} else if direction == "ne" {
			z--
			x++
		} else if direction == "se" {
			y--
			x++
		} else if direction == "s" {
			y--
			z++
		} else if direction == "sw" {
			z++
			x--
		} else if direction == "nw" {
			y++
			x--
		}
	}

	return (abs(x) + abs(y) + abs(z)) / 2
}

func abs(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
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
