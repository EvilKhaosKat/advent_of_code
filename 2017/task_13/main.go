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

	firstResult := solveFirst(input)
	fmt.Println(firstResult)

	//secondResult := solveSecond(input)
	//fmt.Println(secondResult)
}

func solveFirst(input []string) int {
	layers := initLayers(input)

	layers = scannerTick(layers)

	packetCoor := 0
	severity := 0

	for {
		if packetCoor == len(layers)-1 {
			break
		}

		packetCoor++
		packetCaught := isPacketCaught(layers, packetCoor)
		if packetCaught {
			severity += packetCoor * layers[packetCoor].depth
		}
		layers = scannerTick(layers)
	}

	return severity
}

func isPacketCaught(layers []*Layer, packetCoor int) bool {
	layer := layers[packetCoor]

	return layer.scannerPosition == 0
}

func scannerTick(layers []*Layer) []*Layer {
	for _, layer := range layers {
		if !layer.empty {
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
	}

	return layers
}

func initLayers(input []string) []*Layer {
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

	return layers
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
