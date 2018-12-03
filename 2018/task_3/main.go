package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const (
	fabricSize = 1000
)

type claim struct {
	id            int
	topX, topY    int
	width, height int
	overlaps      bool
}

type square []*claim

func main() {
	input := readInput()

	resultFirst := solveFirst(input)
	fmt.Println(resultFirst)

	resultSecond := solveSecond(input)
	fmt.Println(resultSecond)
}

func solveFirst(input []string) int {
	fabric := initFabric(fabricSize)

	for _, inputLine := range input {
		claim := parseClaim(inputLine)
		fabric, _ = applyClaim(fabric, claim)
	}

	return getMultipleClaimsCount(fabric)
}

func solveSecond(input []string) int {
	fabric := initFabric(fabricSize)
	var claims []*claim

	for _, inputLine := range input {
		claim := parseClaim(inputLine)
		fabric, claim = applyClaim(fabric, claim)

		claims = append(claims, claim)
	}

	for _, claim := range claims {
		if !claim.overlaps {
			return claim.id
		}
	}

	panic("not overlapped claim not found")
}

func getMultipleClaimsCount(fabric [][]square) int {
	multipleClaims := 0

	for x := 0; x < len(fabric); x++ {
		for y := 0; y < len(fabric[0]); y++ {
			if overlappedSquare(fabric, x, y) {
				multipleClaims++
			}
		}
	}

	return multipleClaims
}

func applyClaim(fabric [][]square, claim *claim) ([][]square, *claim) {
	for x := claim.topX; x < claim.topX+claim.width; x++ {
		for y := claim.topY; y < claim.topY+claim.height; y++ {
			fabric[x][y] = append(fabric[x][y], claim)

			if overlappedSquare(fabric, x, y) {
				for _, claimToUpdate := range fabric[x][y] {
					claimToUpdate.overlaps = true
				}
			}
		}
	}

	return fabric, claim
}

func overlappedSquare(fabric [][]square, x int, y int) bool {
	return len(fabric[x][y]) > 1
}

//#123 @ 3,2: 5x4
func parseClaim(claimString string) *claim {
	splitFunc := func(c rune) bool {
		return !unicode.IsNumber(c)
	}

	fields := strings.FieldsFunc(claimString, splitFunc)

	//TODO named regexp - better readability?
	return &claim{getNumByString(fields[0]), getNumByString(fields[1]), getNumByString(fields[2]),
		getNumByString(fields[3]), getNumByString(fields[4]),
		false}
}

func initFabric(size int) [][]square {
	fabric := make([][]square, size)

	for i := range fabric {
		fabric[i] = make([]square, size)
	}

	return fabric
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
