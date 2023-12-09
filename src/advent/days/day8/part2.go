package day8

import (
	"strings"
)

const lastCharPos = 2
const startNodes = 'A'
const targetNodes = 'Z'
const leftChar = 'L'
const leftBool = true

type GhostNode struct {
	id       string
	endsWith rune
	left     string
	right    string
}
type GhostNodeMap map[string]GhostNode

type GhostNodeTraversal struct {
	startNodes []GhostNode
	endNotes   []GhostNode
	steps      []bool
}

// CCC = (ZZZ, GGG)
func parseGhostNode(line string) GhostNode {
	split1 := strings.Split(line, "=")
	id := strings.TrimSpace(split1[0])

	split2 := strings.Split(split1[1], ",")

	left := strings.TrimSpace(strings.Replace(split2[0], "(", "", 1))
	right := strings.TrimSpace(strings.Replace(split2[1], ")", "", 1))
	return GhostNode{id: id, left: left, right: right, endsWith: rune(id[lastCharPos])}
}

func parseInput(lines []string) (moves []bool, n GhostNodeMap) {
	for _, char := range lines[0] {
		moves = append(moves, char == leftChar)
	}

	n = make(GhostNodeMap)
	for _, line := range lines[2:] {
		node := parseGhostNode(line)
		n[node.id] = node
	}
	return moves, n
}

func Part2(lines []string) int {
	parseInput(lines)

	// idea1: enumerate all of the possible outcomes (R, L, RR, RL, LR, LL, etc) and find the first that gets to all Z?

	// idea2: find all paths from A^6 to Z^6, and all cycles from A^6 and Z^6 back to themselves

	// idea3: include the history of "seen" nodes in a sequence so they can find cycles

	return 0
}
