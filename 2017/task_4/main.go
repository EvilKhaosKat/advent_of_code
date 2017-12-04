package main

import (
	"os"
	"log"
	"bufio"
	"fmt"
	"strings"
)

func main() {
	input := readInput()

	firstResult := solveFirst(input)

	fmt.Println(firstResult)
}

func solveFirst(input []string) int {
	validPassphrases := 0

	for _, passphrase := range input {
		if isValidPassphrase(passphrase) {
			validPassphrases++
		}
	}

	return validPassphrases
}

func isValidPassphrase(passphrase string) bool {
	usedPassphrases := map[string]bool{}

	for _, word := range strings.Split(passphrase, " ") {
		_, present := usedPassphrases[word]
		if present {
			return false
		} else {
			usedPassphrases[word] = true
		}
	}

	return true
}


func readInput() []string {
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
