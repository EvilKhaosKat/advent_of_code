package main

import (
	"os"
	"log"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	name  string
	links map[string]*Node
}

type Group map[string]bool

//TODO doesn't look clear enough - simplify
func main() {
	input := readInputMultiLine()

	firstResult := solveFirst(input)
	fmt.Println(firstResult)

	secondResult := solveSecond(input)
	fmt.Println(secondResult)
}

func solveFirst(input []string) int {
	nodes := createNodes(input)

	node0 := nodes["0"]
	return len(calculateGroupNodes(node0))
}

func solveSecond(input []string) int {
	nodes := createNodes(input)
	return calculateGroupsCount(nodes)
}

func calculateGroupsCount(nodes map[string]*Node) int {
	var groups []Group

	for _, node := range nodes {
		if isNodeFromFoundGroup(groups, node) {
			continue
		} else {
			groups = append(groups, calculateGroupNodes(node))
		}
	}

	return len(groups)
}

func isNodeFromFoundGroup(groups []Group, node *Node) bool {
	for _, group := range groups {
		if _, present := group[node.name]; present {
			return true
		}
	}

	return false
}

func calculateGroupNodes(anyGraphNode *Node) map[string]bool {
	visitedNodes := make(map[string]bool)
	var nodesToCheck []*Node
	nodesToCheck = append(nodesToCheck, anyGraphNode)

	for {
		if len(nodesToCheck) == 0 {
			break
		}

		var nodeToCheck *Node
		nodeToCheck, nodesToCheck = nodesToCheck[0], nodesToCheck[1:]

		if _, present := visitedNodes[nodeToCheck.name]; present {
			continue
		}

		visitedNodes[nodeToCheck.name] = true

		for _, nodeToAdd := range nodeToCheck.links {
			nodesToCheck = append(nodesToCheck, nodeToAdd)
		}
	}

	return visitedNodes
}

func createNodes(input []string) map[string]*Node {
	nodes := make(map[string]*Node)

	for _, line := range input {
		parts := strings.Split(line, " <-> ")

		nodeName := parts[0]
		connectedNodesNames := strings.Split(parts[1], ", ")

		node := getNode(nodes, nodeName)
		for _, connectedNodeName := range connectedNodesNames {
			connectedNode := getNode(nodes, connectedNodeName)

			connectNodes(node, connectedNode)
		}
	}

	return nodes
}

func connectNodes(node1 *Node, node2 *Node) {
	node1.links[node2.name] = node2
	node2.links[node1.name] = node1
}

func getNode(nodes map[string]*Node, nodeName string) *Node {
	node, present := nodes[nodeName]
	if present {
		return node
	} else {
		newNode := &Node{name: nodeName, links: make(map[string]*Node)}
		nodes[nodeName] = newNode

		return newNode
	}
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
