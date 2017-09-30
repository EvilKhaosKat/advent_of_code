package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
	"math"
)

type Direction int
type Command int

func (d Direction) left() Direction {
	result := d - 1
	if result < north {
		return west
	}

	return result
}

func (d Direction) right() Direction {
	result := d + 1
	if result > west {
		return north
	}

	return result
}

const (
	north Direction = iota
	east
	south
	west
)

const (
	left Command = iota
	right
)

type Me struct {
	x, y      int
	direction Direction
}

func (me *Me) left() {
	newDirection := me.direction.left()
	me.direction = newDirection
}

func (me *Me) right() {
	newDirection := me.direction.right()
	me.direction = newDirection
}

func (me *Me) move(distance int) {
	switch me.direction {
	case north:
		me.y = me.y + distance
	case east:
		me.x = me.x + distance
	case south:
		me.y = me.y - distance
	case west:
		me.x = me.x - distance
	}
}

func main() {
	me := &Me{0, 0, north}

	input := readInput()

	rawCommands := strings.Split(input, ", ")
	for _, rawCommand := range rawCommands {
		executeCommand(rawCommand, me)
	}

	fmt.Println(math.Abs(float64(me.x)) + math.Abs(float64(me.y)))
}

func executeCommand(rawCommand string, me *Me) {
	command, distance := parseCommand(rawCommand)

	if command == left {
		me.left()
	} else {
		me.right()
	}

	me.move(distance)
}

func parseCommand(rawCommand string) (Command, int) {
	command := left

	commandLetter := rawCommand[0:1]
	if commandLetter == "L" {
		command = left
	} else {
		command = right
	}

	rawDistance := rawCommand[1:]
	distance, err := strconv.Atoi(rawDistance)
	if err != nil {
		log.Fatal(err)
	}

	return command, distance
}


func readInput() string {
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
