package main

import (
	"os"
	"log"
	"bufio"
	"fmt"
	"strings"
	"sort"
)

func main() {
	input := readInput()

	firstResult := solveFirst(input)
	secondResult := solveSecond(input)

	fmt.Println(firstResult)
	fmt.Println(secondResult)
}

func solveFirst(input []string) int {
	validPassphrases := 0

	for _, passphrase := range input {
		if isValidPassphraseFirst(passphrase) {
			validPassphrases++
		}
	}

	return validPassphrases
}

func solveSecond(input []string) int {
	validPassphrases := 0

	for _, passphrase := range input {
		if isValidPassphraseSecond(passphrase) {
			validPassphrases++
		}
	}

	return validPassphrases
}

func sortString(mixedString string) string {
	sortedString := strings.Split(mixedString, "")
	sort.Strings(sortedString)
	return strings.Join(sortedString, "")
}

func isValidPassphraseFirst(passphrase string) bool {
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

func isValidPassphraseSecond(passphrase string) bool {
	usedPassphrases := map[string]bool{}

	for _, word := range strings.Split(passphrase, " ") {
		word = sortString(word)

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
