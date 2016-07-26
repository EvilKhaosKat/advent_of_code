package main

import (
	"bufio"
	"io"
	"os"
	"fmt"
	"unicode/utf8"
)

//TODO fsm transitions to be declared here somehow
const (
	NormalChar = iota
	Escape
	EscapeX
	EscapeFirst
	EscapeSecond
)

type AnalyseResult struct {
	charsOfCode  int
	charsOfValue int
	inputString string
}

type State struct {
	id int
}

func main() {
	fmt.Println("Advent of code: Golang\nTask 8")

	PrintCurrentDir()

	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		os.Exit(1)
	}
	defer file.Close()

	handleFile(file)
}

func handleFile(file io.Reader) {
	resultChan := make(chan AnalyseResult)

	linesCount := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		go handleString(scanner.Text(), resultChan)
		linesCount++
	}
	fmt.Println("linesCount:", linesCount)

	totalCharsOfCode := 0
	totalCharsOfValue := 0
	for i := 0; i < linesCount; i++ {
		result := <-resultChan
		fmt.Printf("read from results: %+v\n", result)

		totalCharsOfCode += result.charsOfCode
		totalCharsOfValue += result.charsOfValue
	}

	fmt.Println("totalCharsOfCode:", totalCharsOfCode)
	fmt.Println("totalCharsOfValue:", totalCharsOfValue)

	fmt.Println("result:", totalCharsOfCode - totalCharsOfValue) //1391 - too high
}

func handleString(inputString string, resultChan chan AnalyseResult) {
	previousState := State{id:NormalChar}

	inputStringLen := utf8.RuneCountInString(inputString)
	charsOfValue := 0

	if inputStringLen == 2 {
		//empty string passed
		resultChan <- AnalyseResult{charsOfCode:2, charsOfValue:0, inputString:inputString}
		return
	}

	workingString := inputString[1:inputStringLen - 1] //ignore quotes

	for _, char := range workingString {
		if previousState.id == Escape {
			if char == 'x' {
				//char code sequence starting
				previousState.id = EscapeX
			} else {
				//just escaped char
				charsOfValue++
			}
		} else if previousState.id == EscapeX {
			previousState.id = EscapeFirst
		} else if previousState.id == EscapeFirst {
			previousState.id = EscapeSecond
			charsOfValue++
		} else if previousState.id == EscapeSecond {
			previousState.id = NormalChar
		} else if previousState.id == NormalChar {
			if char == '\\' {
				previousState.id = Escape
			} else {
				//just normal char
				charsOfValue++
			}
		}
	}

	resultChan <- AnalyseResult{charsOfCode:inputStringLen, charsOfValue:charsOfValue,  inputString:inputString}
}

func PrintCurrentDir() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("current dir: " + pwd)
}
