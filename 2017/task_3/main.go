package main

import (
	"os"
	"log"
	"bufio"
	"strconv"
	"fmt"
)

func main() {
	input := readInput()

	firstResult := solveFirst(getNumByString(input))

	fmt.Println(firstResult)
}

func solveFirst(input int) int {
	squareSide := getSquareSide(input)
	squareArea := getSquareArea(squareSide)

	delta := squareArea - input

	deltaRemainder := delta % squareSide

	secondCoor := squareSide - deltaRemainder

	middleCoor := (squareSide + 1) / 2

	return getSquareLevel(squareSide) + abs(middleCoor - secondCoor)
}

func abs(num int) int {
	if num < 0 {
		return (-1) * num
	} else {
		return num
	}
}

func getSquareLevel(squareSide int) int {
	return (squareSide - 1)/2
}

func getSquareSide(input int) int {
	for squareSide := 1; ; squareSide = squareSide + 2 {
		squareArea := getSquareArea(squareSide)
		if squareArea > input {
			return squareSide
		}
	}
}
func getSquareArea(squareSide int) int {
	return squareSide * squareSide
}

func getSquareAreaOLD(squareLevel int) int {
	return (squareLevel*2 + 1) ^ 2
}

func getNumByString(numRaw string) int {
	num, err := strconv.Atoi(numRaw)
	if err != nil {
		panic("Cannot get num:" + err.Error())
	}
	return num
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
