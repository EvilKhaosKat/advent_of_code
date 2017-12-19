package main

import (
	"os"
	"log"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Direction int

const (
	DOWN Direction = iota
	UP
)

type Layer struct {
	empty bool
	depth int

	scannerPosition  int
	scannerDirection Direction
}

func main() {
	input := readInputMultiLine()

	firstResult, _ := solveFirst(input, 0)
	fmt.Println(firstResult)

	secondResult := solveSecond(input)
	fmt.Println(secondResult)
}

func solveFirst(input []string, delayTicks int) (severity int, scannersFailed int) {
	layers := initLayers(input, delayTicks)

	packetCoor := -1

	if delayTicks%1000 == 0 {
		fmt.Println("current delayTicks:" + strconv.Itoa(delayTicks))
	}

	for {
		if packetCoor == len(layers)-1 {
			break
		}

		packetCoor++
		packetCaught := isPacketCaught(layers, packetCoor)
		if packetCaught {
			severity += packetCoor * layers[packetCoor].depth
			scannersFailed++
		}

		layers = scannerTickLayers(layers)
	}

	return severity, scannersFailed
}

func getLayersString(layers []*Layer) string {
	var result string

	for _, layer := range layers {
		if layer.empty {
			result += "."
		} else {
			result += strconv.Itoa(layer.scannerPosition)
		}
	}

	return result
}

//instead of simulation try to calculate position of scanner by delay num
func solveSecond(input []string) int {
	var delayTicks int

	for delayTicks = 0; ; delayTicks++ {
		_, scannersFailed := solveFirst(input, delayTicks)
		if scannersFailed == 0 {
			break
		}
	}

	return delayTicks + 1
}

func isPacketCaught(layers []*Layer, packetCoor int) bool {
	layer := layers[packetCoor]

	return !layer.empty && layer.scannerPosition == 0
}

func scannerTickLayers(layers []*Layer) []*Layer {
	for _, layer := range layers {
		if !layer.empty {
			scannerTickLayer(layer)
		}
	}

	return layers
}

func scannerTickLayer(layer *Layer) {
	depth := layer.depth
	scannerPos := layer.scannerPosition
	scannerDirection := layer.scannerDirection

	if scannerPos == depth-1 && scannerDirection == DOWN {
		layer.scannerPosition--
		layer.scannerDirection = UP
	} else if scannerPos == 0 && scannerDirection == UP {
		layer.scannerPosition++
		layer.scannerDirection = DOWN
	} else if scannerPos >= 0 && scannerDirection == DOWN {
		layer.scannerPosition++
	} else if scannerDirection >= 0 && scannerDirection == UP {
		layer.scannerPosition--
	}
}

func initLayers(input []string, delayTicks int) []*Layer {
	var layers []*Layer

	currentLayerNum := 0

	for _, line := range input {
		layerNum, depth := parseInputLine(line)

		if layerNum > currentLayerNum {
			layers = createEmptyLayers(layers, layerNum-currentLayerNum)
			currentLayerNum = layerNum
		}

		layers = createLayer(layers, depth)
		currentLayerNum++
	}

	if delayTicks > 0 {
		for _, layer := range layers {
			updateLayerScannerState(layer, delayTicks)
		}
	}

	return layers
}

func updateLayerScannerState(layer *Layer, delayTicks int) {
	if layer.empty {
		return
	}

	cycleTicks := layer.depth + layer.depth - 2
	//cyclesPassed := (delayTicks + 1) / cycleTicks
	sidesPassed := (delayTicks + 1) / (cycleTicks/2)
	ticksRemained := (delayTicks + 1) % (cycleTicks/2)
	if ticksRemained < 0 {
		ticksRemained = 0
	}

	evenSides := sidesPassed%2 == 0
	if evenSides {
		layer.scannerDirection = DOWN
	} else {
		layer.scannerPosition = layer.depth - 1
		layer.scannerDirection = UP
	}

	for i := ticksRemained; i > 0; i-- {
		scannerTickLayer(layer)
	}
}

func createEmptyLayers(layers []*Layer, emptyLayersCount int) []*Layer {
	for i := 0; i < emptyLayersCount; i++ {
		layers = append(layers, &Layer{empty: true})
	}

	return layers
}

func createLayer(layers []*Layer, depth int) []*Layer {
	layer := &Layer{depth: depth}

	layers = append(layers, layer)

	return layers
}

func parseInputLine(inputLine string) (layerNum int, depth int) {
	parts := strings.Split(inputLine, ": ")

	return getNumByString(parts[0]), getNumByString(parts[1])
}

//
//helper methods starts here
//
func getNumByString(numRaw string) int {
	num, err := strconv.Atoi(numRaw)
	if err != nil {
		panic("Cannot get num:" + err.Error())
	}
	return num
}

func readInputSingleLine() string {
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

func readInputMultiLine() []string {
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
