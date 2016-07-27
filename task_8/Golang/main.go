package main

import (
	"bufio"
	"io"
	"os"
	"fmt"
	"unicode/utf8"
	"strings"
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
	charsOfCode       int
	charsOfValue      int
	inputString       string
	totalEncodedChars int
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
	totalEncodedChars := 0
	for i := 0; i < linesCount; i++ {
		result := <-resultChan
		fmt.Printf("read from results: %+v\n", result)

		totalCharsOfCode += result.charsOfCode
		totalCharsOfValue += result.charsOfValue
		totalEncodedChars += result.totalEncodedChars
	}

	fmt.Println("totalCharsOfCode:", totalCharsOfCode)
	fmt.Println("totalCharsOfValue:", totalCharsOfValue)

	fmt.Println("encodedResult:", totalCharsOfCode - totalCharsOfValue)
	fmt.Println("escapedResult:", totalEncodedChars - totalCharsOfCode)
}

func handleString(inputString string, resultChan chan AnalyseResult) {
	state := State{id:NormalChar}

	inputStringLen := utf8.RuneCountInString(inputString)
	charsOfValue := 0

	if inputStringLen == 2 {
		//empty string passed
		resultChan <- AnalyseResult{charsOfCode:2, charsOfValue:0, inputString:inputString, totalEncodedChars:6}
		return
	}

	workingString := inputString[1:inputStringLen - 1] //ignore quotes

	for _, char := range workingString {
		updateState(&state, char)

		if state.id == NormalChar || state.id == EscapeSecond {
			charsOfValue++
		}
	}

	resultChan <- AnalyseResult{charsOfCode:inputStringLen,
		charsOfValue:charsOfValue,
		inputString:inputString,
		totalEncodedChars:encodedCharsCount(inputString)}
}

func encodedCharsCount(inputString string) int {
	newString := strings.Replace(inputString, "\\", "\\\\", -1)
	newString = strings.Replace(newString, "\"", "\\\"", -1)

	return utf8.RuneCountInString(newString) + 2 //outer quotes
}

func updateState(oldState *State, char rune) {
	if oldState.id == NormalChar && char != '\\' {
		return
	}

	if oldState.id == Escape {
		if char == 'x' {
			//char code sequence starting
			oldState.id = EscapeX
		} else {
			//just escaped char
			oldState.id = NormalChar
		}
	} else if oldState.id == EscapeX {
		oldState.id = EscapeFirst
	} else if oldState.id == EscapeFirst {
		oldState.id = EscapeSecond
	} else if oldState.id == EscapeSecond {
		if char == '\\' {
			oldState.id = Escape
		} else {
			oldState.id = NormalChar
		}
	} else if oldState.id == NormalChar {
		if char == '\\' {
			oldState.id = Escape
		}
	}
}

func PrintCurrentDir() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("current dir: " + pwd)
}
