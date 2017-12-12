package main

import (
	"os"
	"log"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type node struct {
	name string
	weight int

	rawChildren string

	parent *node
	children []*node
}

type nodes []*node

func main() {
	input := readInputMultiLine()

	firstResult := solveFirst(input)

	fmt.Println(firstResult)
}

func solveFirst(input []string) string {
	topNode := buildGraph(input)

	return topNode.name
}

func buildGraph(input []string) *node {
	foundNodes := initNodes(input)
	foundNodes = updateParentChildrenInfo(foundNodes)

	return getTopParent(getFirstElement(foundNodes))
}

func getFirstElement(nodes map[string]*node) *node {
	for _, node := range nodes {
		return node
	}

	return nil
}

func getTopParent(node *node) *node {
	currentNode := node

	for currentNode.parent != nil {
		currentNode = currentNode.parent
	}

	return currentNode
}

func updateParentChildrenInfo(foundNodes map[string]*node) map[string]*node {
	for _, node := range foundNodes {
		rawChildren := node.rawChildren
		if rawChildren != "" {
			var children nodes

			childrenRaw := strings.Split(rawChildren, ", ")
			for _, childrenName := range childrenRaw {
				childrenNode := foundNodes[childrenName]

				children = append(children, childrenNode)
				childrenNode.parent = node
			}

			node.children = children
		}
	}

	return foundNodes
}

func initNodes(input []string) map[string]*node {
	foundNodes := make(map[string]*node)

	for _, line := range input {
		nodeAndChildParts := strings.Split(line, " -> ")
		nodePart := nodeAndChildParts[0]

		node := parseNodeInfo(nodePart)

		if len(nodeAndChildParts) == 2 {
			node.rawChildren = nodeAndChildParts[1]
		}

		foundNodes[node.name] = node
	}

	return foundNodes
}

//uwzmqi (57)
func parseNodeInfo(nodePart string) *node {
	node := &node{}

	nameWeightParts := strings.Split(nodePart, " ")
	node.name = nameWeightParts[0]
	node.weight = parseWeight(nameWeightParts[1])

	return node
}

//(123)
func parseWeight(weightPart string) int {
	rawWeight := weightPart[1:len(weightPart) - 1]

	return getNumByString(rawWeight)
}

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


