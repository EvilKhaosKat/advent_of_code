package main

import (
	"os"
	"log"
	"bufio"
	"strconv"
	"fmt"
)

type Direction int

const (
	up    Direction = iota
	left
	down
	right
)

func main() {
	input := readInput()

	firstResult := solveFirst(getNumByString(input))
	secondResult := solveSecond(getNumByString(input))

	fmt.Println(firstResult)
	fmt.Println(secondResult)
}

func solveFirst(input int) int {
	squareSide := getSquareSideFirst(input)
	squareArea := getSquareArea(squareSide)

	middleCoor := (squareSide + 1) / 2

	fourthCorner := squareArea
	thirdCorner := fourthCorner - squareSide + 1
	secondCorner := thirdCorner - squareSide + 1
	firstCorner := secondCorner - squareSide + 1

	squareLevel := getSquareLevel(squareSide)

	return squareLevel + getSecondCoor(input, squareLevel, middleCoor, firstCorner, secondCorner, thirdCorner, fourthCorner)
}

func solveSecond(input int) int {
	squareSide := getSquareSideSecond(input)
	square := initSquare(squareSide)

	middleCoor := squareSide/2

	square[middleCoor][middleCoor] = 1

	direction := right

	x := middleCoor
	y := middleCoor

	for {
		direction, x, y = updatePosition(square, direction, x, y)
		newValue := fillCurrentValue(square, x, y)
		square[x][y] = newValue

		//printSquare(square)

		if newValue > input {
			return newValue
		}
	}
}

func printSquare(square [][]int) {
	squareSide := len(square)

	for i := 0; i < squareSide; i++ {
		for j := 0; j < squareSide; j++ {
			fmt.Print(strconv.Itoa(square[i][j]) + " ")
		}

		fmt.Println()
	}

	fmt.Println()
}

func fillCurrentValue(square [][]int, x int, y int) int {
	value := 0

	value += square[x-1][y-1]
	value += square[x-1][y]
	value += square[x-1][y+1]
	value += square[x][y-1]
	value += square[x][y]
	value += square[x][y+1]
	value += square[x+1][y-1]
	value += square[x+1][y]
	value += square[x+1][y+1]

	return value
}

func updatePosition(square [][]int, direction Direction, x int, y int) (Direction, int, int) {
	x, y = getNewCoor(direction, x, y)

	switch direction {
	case up:
		if square[x-1][y] == 0 {
			direction = left
		}
	case left:
		if square[x][y+1] == 0 {
			direction = down
		}
	case down:
		if square[x+1][y] == 0 {
			direction = right
		}
	case right:
		if square[x][y-1] == 0 {
			direction = up
		}
	}

	return direction, x, y
}


func getNewCoor(direction Direction, x int, y int) (int, int) {
	switch direction {
	case up:
		y = y - 1
	case left:
		x = x - 1
	case down:
		y = y + 1
	case right:
		x = x + 1
	}

	return x, y
}

func initSquare(squareSide int) [][]int {
	var square [][]int

	for i := 0; i < squareSide; i++ {
		var line []int

		for j := 0; j < squareSide; j++ {
			line = append(line, 0)
		}

		square = append(square, line)
	}

	return square
}

func getSecondCoor(input int, squareLevel int, middleCoor int, firstCorner int, secondCorner int, thirdCorner int, fourthCorner int) int {
	if input == firstCorner || input == secondCorner || input == thirdCorner || input == fourthCorner {
		return squareLevel
	}

	if input < firstCorner {
		return abs(middleCoor-(firstCorner-input)) - 1
	} else if input < secondCorner {
		return abs(middleCoor-(secondCorner-input)) - 1
	} else if input < thirdCorner {
		return abs(middleCoor-(thirdCorner-input)) - 1
	} else if input < fourthCorner {
		return abs(middleCoor-(fourthCorner-input)) - 1
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
	return (squareSide - 1) / 2
}

//TODO replace with sqrt upper ceiling
func getSquareSideFirst(input int) int {
	for squareSide := 1; ; squareSide = squareSide + 2 {
		squareArea := getSquareArea(squareSide)
		if squareArea > input {
			return squareSide
		}
	}
}

func getSquareSideSecond(input int) int {
	return getSquareSideFirst(input) + 2 //TODO second requires significantly less size - calculate in better way
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
