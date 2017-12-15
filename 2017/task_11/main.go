package main

import (
	"os"
	"log"
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"math"
)

func main() {
	input := readInputSingleLine()

	firstResult := solveFirst(input)
	fmt.Println(firstResult)

	secondResult := solveSecond(input)
	fmt.Println(secondResult)

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

	return getDistanceFromStart(x, y, z)
}

func solveSecond(input string) int {
	x := 0
	y := 0
	z := 0

	maxDistance := math.MinInt32

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

		currentDistance := getDistanceFromStart(x, y, z)
		if currentDistance > maxDistance {
			maxDistance = currentDistance
		}
	}

	return maxDistance
}

func getDistanceFromStart(x int, y int, z int) int {
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
