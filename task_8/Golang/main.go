package main

import (
	"bufio"
	"io"
	"os"
	"fmt"
)

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
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
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
