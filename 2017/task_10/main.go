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
	parsedInput := parseInput(input)

	firstResult := solveFirst(256, parsedInput)
	//secondResult := solveSecond(input)

	fmt.Println(firstResult)
	//fmt.Println(secondResult)
}

func parseInput(input string) []int {
	var parsedInput []int

	rawNumbers := strings.Split(input, ",")

	for _, rawNumber := range rawNumbers {
		parsedInput = append(parsedInput, getNumByString(rawNumber))
	}

	return parsedInput

}

func solveFirst(listSize int, lengths []int) int {
	list := createList(listSize)

	currentPos := 0
	skipSize := 0

	for _, length := range lengths {
		list = reverseOrder(list, currentPos%listSize, length)

		currentPos += length + skipSize
		skipSize++
	}

	return list[0] * list[1]
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
		exchangeValues(list, i, listLength - i - 1)
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
