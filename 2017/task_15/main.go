package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
)

const (
	GENERATOR_A_FACTOR = 16807
	GENERATOR_B_FACTOR = 48271
	DIVISOR            = 2147483647
	ITERATIONS_FIRST   = 40000000
	ITERATIONS_SECOND  = 5000000

	MULTIPLE_CHECK_A           = 4
	MULTIPLE_CHECK_B           = 8
	NO_MULTIPLE_CHECK_REQUIRED = -1
)

type Generator struct {
	results chan int
}

func (generator *Generator) run(factor, startNum, multipleBy, requiredIterationsCount int) chan int {
	generator.results = make(chan int, 42) //assuming 42 is enough for everything =^__^=
	value := startNum

	go func() {
		currentIterationNum := 0
		for {
			if currentIterationNum >= requiredIterationsCount {
				break
			}

			value = (value * factor) % DIVISOR

			if multipleBy == NO_MULTIPLE_CHECK_REQUIRED || value%multipleBy == 0 {
				generator.results <- value
				currentIterationNum++
			}
		}

		close(generator.results)
	}()

	return generator.results
}

func main() {
	f, err := os.Create("meow.pprof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	input := readInputMultiLine()

	firstResult := solveFirst(parseInput(input))
	fmt.Println(firstResult)

	secondResult := solveSecond(parseInput(input))
	fmt.Println(secondResult)
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

	resultsA := generatorA.run(GENERATOR_A_FACTOR, generatorAStart, NO_MULTIPLE_CHECK_REQUIRED, ITERATIONS_FIRST)
	resultsB := generatorB.run(GENERATOR_B_FACTOR, generatorBStart, NO_MULTIPLE_CHECK_REQUIRED, ITERATIONS_FIRST)

	judgeCount := 0

	for i := 0; i < ITERATIONS_FIRST; i++ {
		//logIfRequired(i)

		valueA := <-resultsA
		valueB := <-resultsB

		if checkEqualGeneratorsValues(valueA, valueB) {
			judgeCount++
		}
	}

	return judgeCount
}

func solveSecond(generatorAStart, generatorBStart int) int {
	generatorA := Generator{}
	generatorB := Generator{}

	resultsA := generatorA.run(GENERATOR_A_FACTOR, generatorAStart, MULTIPLE_CHECK_A, ITERATIONS_SECOND)
	resultsB := generatorB.run(GENERATOR_B_FACTOR, generatorBStart, MULTIPLE_CHECK_B, ITERATIONS_SECOND)

	judgeCount := 0

	for i := 0; i < ITERATIONS_SECOND; i++ {
		//logIfRequired(i)

		valueA := <-resultsA
		valueB := <-resultsB

		if checkEqualGeneratorsValues(valueA, valueB) {
			judgeCount++
		}
	}

	return judgeCount
}

func logIfRequired(iteration int) {
	if iteration%100000 == 0 {
		fmt.Println("iteration:" + strconv.Itoa(iteration))
	}
}

func checkEqualGeneratorsValues(valueA int, valueB int) bool {
	checkValueA := valueA & 0xFFFF
	checkValueB := valueB & 0xFFFF

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
