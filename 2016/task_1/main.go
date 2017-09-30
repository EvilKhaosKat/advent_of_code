package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
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

type Line struct {
	start, end Coor
}

func (line Line) isIntersected(otherLine Line) *Coor {
	var horizontalLine, verticalLine Line

	if line.start.y == line.end.y {
		horizontalLine = line
		verticalLine = otherLine
	} else {
		horizontalLine = otherLine
		verticalLine = line
	}

	point := &Coor{verticalLine.start.x, horizontalLine.start.y}

	if point.x > min(horizontalLine.start.x, horizontalLine.end.x) &&
		point.x < max(horizontalLine.start.x, horizontalLine.end.x) &&
		point.y > min(verticalLine.start.y, verticalLine.end.y) &&
		point.y < max(verticalLine.start.y, verticalLine.end.y) {
		return point
	} else {
		return nil
	}
}

func min(x int, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func max(x int, y int) int {
	return int(math.Max(float64(x), float64(y)))
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
	coor      Coor
	direction Direction

	movesLines            []Line
	firstIntersectionCoor *Coor
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
	start := me.coor

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

	end := me.coor

	if me.firstIntersectionCoor != nil {
		return
	}

	newMoveLine := Line{start, end}

	intersectionCoordinate := checkIntersected(newMoveLine, me.movesLines)
	if intersectionCoordinate != nil {
		me.firstIntersectionCoor = intersectionCoordinate
	}

	me.movesLines = append(me.movesLines, newMoveLine)
}

func checkIntersected(line Line, lines []Line) *Coor {
	for _, testLine := range lines {
		if intersectionCoor := line.isIntersected(testLine); intersectionCoor != nil {
			return intersectionCoor
		}
	}

	return nil
}

func main() {
	me := initMe()

	input := readInput()

	rawCommands := strings.Split(input, ", ")
	for _, rawCommand := range rawCommands {
		executeCommand(rawCommand, me)
	}

	fmt.Println(math.Abs(float64(me.coor.x)) + math.Abs(float64(me.coor.y)))
	fmt.Println(math.Abs(float64(me.firstIntersectionCoor.x)) + math.Abs(float64(me.firstIntersectionCoor.y)))
	fmt.Println(me.firstIntersectionCoor)
	fmt.Println(me.movesLines)
}

func initMe() *Me {
	return &Me{Coor{0, 0}, north, nil, nil}
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
