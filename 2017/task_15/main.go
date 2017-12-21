package main

import (
	"os"
	"log"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

const (
	GENERATOR_A_FACTOR = 16807
	GENERATOR_B_FACTOR = 48271
	DIVISOR            = 2147483647
	ITERATIONS         = 40000000
)

type Generator struct {
	results chan int
}

func (generator *Generator) run(factor, startNum, iterations int) chan int {
	generator.results = make(chan int, 42) //assuming 42 is enough for everything =^__^=
	value := startNum

	go func() {
		for currentIterationNum := 0; currentIterationNum < iterations; currentIterationNum++ {
			value = (value * factor) % DIVISOR
			generator.results <- value
		}

		close(generator.results)
	}()

	return generator.results
}

func main() {
	input := readInputMultiLine()

	firstResult := solveFirst(parseInput(input))
	fmt.Println(firstResult)

	//secondResult := solveSecond(input)
	//fmt.Println(secondResult)
}

func parseInput(input []string) (int, int) {
	return getGeneratorNumFromString(input[0]), getGeneratorNumFromString(input[1])
}

func getGeneratorNumFromString(generatorInfoLine string) int {
	parts := strings.Split(generatorInfoLine, " ")

	return getNumByString(parts[len(parts)-1])
}

func solveFirst(generatorAStart, generatorBStart int) int {
	generatorA := Generator{}
	generatorB := Generator{}

	resultsA := generatorA.run(GENERATOR_A_FACTOR, generatorAStart, ITERATIONS)
	resultsB := generatorB.run(GENERATOR_B_FACTOR, generatorBStart, ITERATIONS)

	judgeCount := 0

	for i := 0; i<ITERATIONS; i++ {
		logIfRequired(i)

		valueA := <- resultsA
		valueB := <- resultsB

		if checkEqualGeneratorsValues(valueA, valueB) {
			judgeCount++
		}
	}

	return judgeCount
}

func logIfRequired(iteration int) {
	if iteration % 100000 == 0 {
		fmt.Println("iteration:" + strconv.Itoa(iteration))
	}
}

func checkEqualGeneratorsValues(valueA int, valueB int) bool {
	checkValueA := valueA << (64 - 16)
	checkValueB := valueB << (64 - 16)

	return checkValueA == checkValueB
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
