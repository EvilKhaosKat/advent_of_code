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

	totalWeight int

	rawChildren string

	parent *node
	children []*node
}

func (node *node) recalculateTotalWeight() int {
	node.totalWeight = node.weight

	if len(node.children) > 0 {
		for _, child := range node.children {
			node.totalWeight += child.recalculateTotalWeight()

		}
	}

	return node.totalWeight
}

type refToNode *node
type nodes []*node


func main() {
	input := readInputMultiLine()

	firstResult := solveFirst(input)
	secondResult := solveSecond(input)

	fmt.Println(firstResult)
	fmt.Println(secondResult)
}

func solveFirst(input []string) string {
	topNode := buildGraph(input)

	return topNode.name
}

func solveSecond(input []string) int {
	topNode := buildGraph(input)
	topNode.recalculateTotalWeight()

	return getRequiredForBalanceWeight(topNode)
}

func getRequiredForBalanceWeight(node *node) int {
	nodesToCheck := make(chan refToNode, 42) //good enough nodes buffer :3
	nodesToCheck <- node

	for checkNode := range nodesToCheck {
		children := checkNode.children

		if len(children) > 0 {
			childrenCorrect, _, errorNode := isChildrenCorrect(children)
			if childrenCorrect {
				addChildrenToCheck(nodesToCheck, children)
			} else {
				lowestCorrectNode, lowestErrorNode := getLowestErrorNodes(errorNode)
				return getCorrectWeight(lowestCorrectNode, lowestErrorNode)
			}
		}
	}

	return -1
}

func getLowestErrorNodes(errorNode *node) (lowestCorrectNode *node, lowestErrorNode *node) {
	nodeToCheck := errorNode

	for {
		children := nodeToCheck.children

		if len(children) > 0 {
			childrenCorrect, _, errorNode := isChildrenCorrect(children)
			if childrenCorrect {
				return getLowestCorrectNode(nodeToCheck), nodeToCheck
			} else {
				nodeToCheck = errorNode
			}
		}

	}
}

func getLowestCorrectNode(node *node) *node {
	for _, child := range node.parent.children {
		if child != node {
			return child
		}
	}

	return nil
}

func getCorrectWeight(correctNode *node, errorNode *node) int {
	deltaWeight := correctNode.totalWeight - errorNode.totalWeight

	return errorNode.weight + deltaWeight
}

func addChildrenToCheck(nodesToCheck chan refToNode, nodes []*node) {
	for _, node := range nodes {
		nodesToCheck <- node
	}
}

func isChildrenCorrect(nodes []*node) (correct bool, correctNode *node, errorNode *node) {
	for _, childNode := range nodes {
		for _, childNodeToCheck := range nodes {
			if childNode == childNodeToCheck {
				continue
			}

			if childNode.totalWeight != childNodeToCheck.totalWeight {
				if otherNodesSameWeightExist(nodes, childNode) {
					return false, childNode, childNodeToCheck
				} else {
					return false, childNodeToCheck, childNode
				}
			}
		}
	}

	return true, nil, nil
}

func otherNodesSameWeightExist(nodes []*node, node *node) bool {
	for _, nodeToCheck := range nodes {
		if node == nodeToCheck {
			continue
		}

		if node.totalWeight == nodeToCheck.totalWeight {
			return true
		}
	}

	return false
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


