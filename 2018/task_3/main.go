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
	id            string
	topX, topY    int
	width, height int
}

func main() {
	input := readInput()

	resultFirst := solveFirst(input)
	fmt.Println(resultFirst)

	//resultSecond := solveSecond(input)
	//fmt.Println(resultSecond)
}

func solveFirst(input []string) int {
	fabric := initFabric(fabricSize)

	for _, inputLine := range input {
		claim := parseClaim(inputLine)
		fabric = applyClaim(fabric, claim)
	}

	return getMultipleClaimsCount(fabric)
}

func getMultipleClaimsCount(fabric [][]int) int {
	multipleClaims := 0

	for x := 0; x < len(fabric); x++ {
		for y := 0; y < len(fabric[0]); y++ {
			if fabric[x][y] > 1 {
				multipleClaims++
			}
		}
	}

	return multipleClaims
}

func applyClaim(fabric [][]int, claim claim) [][]int {
	for x := claim.topX; x < claim.topX+claim.width; x++ {
		for y := claim.topY; y < claim.topY+claim.height; y++ {
			fabric[x][y] = fabric[x][y] + 1
		}
	}

	return fabric
}

//#123 @ 3,2: 5x4
func parseClaim(claimString string) claim {
	splitFunc := func(c rune) bool {
		return !unicode.IsNumber(c)
	}

	fields := strings.FieldsFunc(claimString, splitFunc)

	//TODO named regexp - better readability?
	return claim{fields[0], getNumByString(fields[1]), getNumByString(fields[2]),
		getNumByString(fields[3]), getNumByString(fields[4])}
}

func initFabric(size int) [][]int {
	fabric := make([][]int, size)

	for i := range fabric {
		fabric[i] = make([]int, size)
	}

	//TODO keep id of claims?
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
