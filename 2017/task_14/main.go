package main

import (
	"os"
	"log"
	"bufio"
	"fmt"
	"strconv"
	"bytes"
)

func main() {
	input := readInputSingleLine()

	firstResult := solveFirst(input)
	fmt.Println(firstResult)

	//secondResult := solveSecond(input)
	//fmt.Println(secondResult)
}

func solveFirst(knotInput string) int {
	var grid [][]int

	for i := 0; i < 128; i++ {
		knotHash := getKnotHash(256, parseInputKnotHash(knotInput+"-"+strconv.Itoa(i)))
		binaryRepresenation := getBinaryKnotHash(knotHash)

		var line []int
		for _, symbol := range binaryRepresenation {

			num, err := strconv.Atoi(string(symbol))
			if err != nil {
				panic(err)
			}

			line = append(line, num)
		}

		grid = append(grid, line)
	}

	return getUsedBlockCount(grid)
}

func getUsedBlockCount(grid [][]int) int {
	count := 0

	for _, line := range grid {
		for _, num := range line {
			if num == 1 {
				count++
			}
		}
	}

	return count
}

func getBinaryKnotHash(knotHash string) string {
	var result bytes.Buffer

	//TODO quite ugly transformation from hex symbol to 4-binary form
	for _, symbol := range knotHash {
		intValue, err := strconv.ParseUint(string(symbol), 16, 64)
		if err != nil {
			panic(err)
		}

		result.WriteString(fmt.Sprintf("%04b", intValue))
	}

	return result.String()
}

//
// knot hash stuff
//
func parseInputKnotHash(input string) []int {
	var parsedInput []int

	for _, rawNumber := range input {
		parsedInput = append(parsedInput, int(rawNumber))
	}

	parsedInput = append(parsedInput, []int{17, 31, 73, 47, 23}...)

	return parsedInput
}

func getKnotHash(listSize int, lengths []int) string {
	list := createList(listSize)

	sparseHash := getSparseHash(lengths, list, listSize)
	denseHash := getDenseHash(sparseHash)
	hexString := getHexString(denseHash)

	return hexString
}

func getHexString(denseHash []int) string {
	var hexStringBuffer bytes.Buffer

	for _, num := range denseHash {
		hexStringPart := fmt.Sprintf("%02x", num)
		hexStringBuffer.WriteString(hexStringPart)
	}

	return hexStringBuffer.String()
}

func getDenseHash(sparseHash []int) []int {
	partsBy16 := splitHashBy16(sparseHash)

	var denseHash []int

	for i := 0; i < 16; i++ {
		denseHash = append(denseHash, getXoredPart(partsBy16[i]))
	}

	return denseHash
}

func getXoredPart(partBy16 []int) int {
	xor := 0

	for _, num := range partBy16 {
		xor = xor ^ num
	}
	return xor
}

func splitHashBy16(sparseHash []int) [][]int {
	var partsBy16 [][]int

	for partNum := 0; partNum < 16; partNum++ {
		var part []int

		for i := 1; i <= 16; i++ {
			part = append(part, sparseHash[partNum*16+i-1])
		}

		partsBy16 = append(partsBy16, part)
	}

	return partsBy16
}

func getSparseHash(lengths []int, list []int, listSize int) []int {
	currentPos := 0
	skipSize := 0
	for i := 0; i < 64; i++ {
		for _, length := range lengths {
			list = reverseOrder(list, currentPos%listSize, length)

			currentPos += length + skipSize
			skipSize++
		}
	}
	return list
}

func reverseOrder(list []int, currentPos int, length int) []int {
	sublist := getSublist(list, currentPos, length)
	sublist = reverseList(sublist)

	for i := 0; i < length; i++ {
		listIndex := (currentPos + i) % len(list)
		list[listIndex] = sublist[i]
	}

	return list
}

func reverseList(list []int) []int {
	listLength := len(list)

	for i := 0; i < listLength/2; i++ {
		exchangeValues(list, i, listLength-i-1)
	}

	return list
}

func getSublist(list []int, currentPos int, length int) []int {
	var sublist []int

	for i := currentPos; i < currentPos+length; i++ {
		actualIndex := i % len(list)

		sublist = append(sublist, list[actualIndex])
	}

	return sublist
}

func exchangeValues(list []int, fromIndex int, toIndex int) {
	fromValue := list[fromIndex]
	list[fromIndex] = list[toIndex]
	list[toIndex] = fromValue
}

func createList(listSize int) []int {
	var list []int

	for i := 0; i < listSize; i++ {
		list = append(list, i)
	}

	return list
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
