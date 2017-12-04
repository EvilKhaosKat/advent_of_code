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

	middleCoor := (squareSide + 1) / 2

	fourthCorner := squareArea
	thirdCorner := fourthCorner - squareSide + 1
	secondCorner := thirdCorner - squareSide + 1
	firstCorner := secondCorner - squareSide + 1

	squareLevel := getSquareLevel(squareSide)

	return squareLevel + getSecondCoor(input, squareLevel, middleCoor, firstCorner, secondCorner, thirdCorner, fourthCorner)
}

func getSecondCoor(input int, squareLevel int, middleCoor int, firstCorner int, secondCorner int, thirdCorner int, fourthCorner int) int {
	if input == firstCorner || input == secondCorner || input == thirdCorner || input == fourthCorner {
		return squareLevel
	}

	if input < firstCorner {
		return abs(middleCoor - (firstCorner - input)) - 1
	} else if input < secondCorner {
		return abs(middleCoor - (secondCorner - input)) - 1
	} else if input < thirdCorner {
		return abs(middleCoor - (thirdCorner - input)) - 1
	} else if input < fourthCorner {
		return abs(middleCoor - (fourthCorner - input)) - 1
	} else {
		panic("second coor calc failed")
	}
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
