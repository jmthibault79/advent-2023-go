package main

import (
	"advent/days/day8"
	"advent/util"
	"fmt"
	"strings"
)

const tooMany = 1_000_000
const startNode = "AAA"
const targetNode = "ZZZ"
const leftChar = 'L'
const leftBool = true

type Node struct {
	id    string
	left  string
	right string
}
type NodeMap map[string]Node

// CCC = (ZZZ, GGG)
func parseNode(line string) Node {
	split1 := strings.Split(line, "=")
	id := strings.TrimSpace(split1[0])

	split2 := strings.Split(split1[1], ",")

	left := strings.TrimSpace(strings.Replace(split2[0], "(", "", 1))
	right := strings.TrimSpace(strings.Replace(split2[1], ")", "", 1))
	return Node{id: id, left: left, right: right}
}

func parseInput(lines []string) (moves []bool, n NodeMap) {

	for _, char := range lines[0] {
		moves = append(moves, char == leftChar)
	}

	n = make(NodeMap)
	for _, line := range lines[2:] {
		node := parseNode(line)
		n[node.id] = node
	}
	return moves, n
}

func day8part1(moves []bool, n NodeMap) (moveCount int) {
	current := n[startNode]
	fmt.Print(current.id)
	for moveCount := 0; moveCount < tooMany; moveCount++ {
		if current.id == targetNode {
			fmt.Println()
			return moveCount
		}

		if moves[moveCount%len(moves)] == leftBool {
			current = n[current.left]
		} else {
			current = n[current.right]
		}
		fmt.Print(" -> ", current.id)
	}
	return tooMany
}

func main() {
	lines1 := util.ReadInput("test", "8")
	moves, nodeMap := parseInput(lines1)
	result := day8part1(moves, nodeMap)
	fmt.Println("Part1", result)

	lines2 := util.ReadInput("test", "8b")
	result2 := day8.Part2(lines2)
	fmt.Println("Part2", result2)
}
