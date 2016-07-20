package main

import (
	"bufio"
	"io"
	"os"
	"fmt"
)

type analyseResult struct {
	charsOfCode  int
	charsOfValue int
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
	resultChan := make(chan analyseResult)

	linesCount := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		go handleString(scanner.Text(), resultChan)
		linesCount++
	}
	fmt.Println("linesCount: ", linesCount)

	for i := 0; i < linesCount; i++ {
		fmt.Printf("read from results: %+v\n", <-resultChan)
	}
}

func handleString(inputString string, resultChan chan analyseResult) {
	//fmt.Println("handleString:" + inputString)
	resultChan <- analyseResult{charsOfCode: 1, charsOfValue: 1}
}

func PrintCurrentDir() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("current dir: " + pwd)
}
