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

type Command int
type Direction int

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

type Coor struct {
	x, y int
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
	coor Coor
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
		me.coor.y = me.coor.y + distance
	case east:
		me.coor.x = me.coor.x + distance
	case south:
		me.coor.y = me.coor.y - distance
	case west:
		me.coor.x = me.coor.x - distance
	}
}

func main() {
	me := &Me{Coor{0, 0}, north}

	input := readInput()

	rawCommands := strings.Split(input, ", ")
	for _, rawCommand := range rawCommands {
		executeCommand(rawCommand, me)
	}

	fmt.Println(math.Abs(float64(me.coor.x)) + math.Abs(float64(me.coor.y)))
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
