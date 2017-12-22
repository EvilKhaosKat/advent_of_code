package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	firstResult := solveFirst(readInputSingleLine())
	fmt.Println(firstResult)

	//secondResult := solveSecond(parseInput(input))
	//fmt.Println(secondResult)
}

func solveFirst(input string) string {
	programs := getProgramsByLine("abcdefghijklmnop")

	commands := strings.Split(input, ",")
	for _, command := range commands {
		commandType := command[0]
		if commandType == 's' {
			programs = spin(programs, getNumByString(command[1:]))
		} else if commandType == 'x' {
			programs = exchange(programs, command[1:])
		} else if commandType == 'p' {
			programs = partner(programs, command[1:])
		}
	}

	return getStringByPrograms(programs)
}

//params=1/2
func exchange(programs []string, rawParams string) []string {
	parts := strings.Split(rawParams, "/")

	firstNum, secondNum := getNumByString(parts[0]), getNumByString(parts[1])

	return swapProgramsByCoor(programs, firstNum, secondNum)
}

func swapProgramsByCoor(programs []string, firstNum int, secondNum int) []string {
	firstValue := programs[firstNum]
	secondValue := programs[secondNum]

	programs[firstNum] = secondValue
	programs[secondNum] = firstValue

	return programs
}

//params=a/b
func partner(programs []string, rawParams string) []string {
	parts := strings.Split(rawParams, "/")

	firstProgram, secondProgram := parts[0], parts[1]

	firstNum := getProgramCoor(programs, firstProgram)
	secondNum := getProgramCoor(programs, secondProgram)

	return swapProgramsByCoor(programs, firstNum, secondNum)
}

func getProgramCoor(programs []string, programToFind string) int {
	for index, program := range programs {
		if program == programToFind {
			return index
		}
	}

	panic(fmt.Sprintf("not found program '%s' in programs '%s'", programToFind, programs))
}

func getStringByPrograms(programs []string) string {
	return strings.Join(programs, "")
}

func getProgramsByLine(line string) []string {
	return strings.Split(line, "")
}

func spin(programs []string, spinAmount int) []string {
	spinCoor := len(programs) - spinAmount

	headPart := programs[:spinCoor]
	tailPart := programs[spinCoor:]

	var result []string

	result = appendSlice(result, tailPart)
	result = appendSlice(result, headPart)

	return result
}
func appendSlice(result []string, toBeAppended []string) []string {
	for _, part := range toBeAppended {
		result = append(result, part)
	}

	return result
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
