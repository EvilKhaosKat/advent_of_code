package main

import "fmt"
import "io/ioutil"
import "os"

func main() {
	fmt.Println("Advent of code: Golang\nTask 8")

	PrintCurrentDir()

	file, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	content := string(file)
	fmt.Println(content)
}

func PrintCurrentDir() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("current dir: " + pwd)
}

