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
	input := readInputMultiLine()

	firstResult := solveFirst(input)
	secondResult := solveSecond(input)

	fmt.Println(firstResult)
	fmt.Println(secondResult)
}

func solveFirst(input []string) int {
	registers := make(map[string]int)

	for _, line := range input {
		operationCondition := strings.Split(line, " if ")

		operation := operationCondition[0]
		condition := operationCondition[1]

		if isConditionTrue(condition, registers) {
			registers = performOperation(operation, registers)
		}
	}

	return getMaxValue(registers)
}

func solveSecond(input []string) int {
	registers := make(map[string]int)
	maxValue := math.MinInt32

	for _, line := range input {
		operationCondition := strings.Split(line, " if ")

		operation := operationCondition[0]
		condition := operationCondition[1]

		if isConditionTrue(condition, registers) {
			registers = performOperation(operation, registers)

			currentMaxValue := getMaxValue(registers)
			if currentMaxValue > maxValue {
				maxValue = currentMaxValue
			}
		}
	}

	return maxValue
}

func getMaxValue(registers map[string]int) int {
	maxValue := math.MinInt64

	for _, value := range registers {
		if value > maxValue {
			maxValue = value
		}
	}

	return maxValue
}

func performOperation(operationString string, registers map[string]int) map[string]int {
	register, operation, operand := getParts(operationString)

	value := getRegisterValue(register, registers)

	if operation == "inc" {
		value = value + operand
	} else if operation == "dec" {
		value = value - operand
	}

	registers[register] = value

	return registers
}

func getParts(operationString string) (register string, operation string, operand int) {
	parts := strings.Split(operationString, " ")

	return parts[0],  parts[1], getNumByString(parts[2])
}

func isConditionTrue(condition string, registers map[string]int) bool {
	register, operation, operand := getParts(condition)

	value := getRegisterValue(register, registers)

	if operation == ">" {
		return value > operand
	} else if operation == "<" {
		return value < operand
	} else if operation == "==" {
		return value == operand
	} else if operation == "<=" {
		return value <= operand
	} else if operation == ">=" {
		return value >= operand
	} else if operation == "!=" {
		return value != operand
	} else {
		panic("unknown operation found:" + operation)
	}
}

func getRegisterValue(register string, registers map[string]int) int {
	registerValue, present := registers[register]

	if present {
		return registerValue
	} else {
		return 0
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


