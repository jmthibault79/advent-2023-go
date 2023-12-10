package day8

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
)

const lastCharPos = 2
const startRune = 'A'
const targetRune = 'Z'
const leftChar = 'L'
const leftBool = true
const rightBool = false

// empirical values from input
const numGhosts = 2

type Node struct {
	id       string
	endsWith rune
	left     string
	right    string
}
type GhostNode [numGhosts]Node
type NodeMap map[string]Node

func (gn GhostNode) String() string {
	var b strings.Builder
	b.WriteString(gn[0].id)
	for _, n := range gn[1:] {
		b.WriteString("," + n.id)
	}
	return b.String()
}

func (gn GhostNode) left(m NodeMap) (out GhostNode) {
	lefts := make([]Node, 0, numGhosts)
	for _, n := range gn {
		lefts = append(lefts, m[n.left])
	}
	slices.SortFunc(lefts, nodeSliceComparator)
	copy(out[:], lefts)
	return
}

func (gn GhostNode) right(m NodeMap) (out GhostNode) {
	rights := make([]Node, 0, numGhosts)
	for _, n := range gn {
		rights = append(rights, m[n.right])
	}
	slices.SortFunc(rights, nodeSliceComparator)
	copy(out[:], rights)
	return
}

type GhostPath struct {
	start GhostNode
	end   GhostNode
	steps []bool
}

func (gp GhostPath) String() string {
	var sb strings.Builder
	for _, s := range gp.steps {
		if s == leftBool {
			sb.WriteString("L")
		} else {
			sb.WriteString("R")
		}
	}

	return fmt.Sprintf("%s -> %s in %d steps:%s",
		gp.start.String(), gp.end.String(), len(gp.steps), sb.String())
}

// CCC = (ZZZ, GGG)
func parseNode(line string) Node {
	split1 := strings.Split(line, "=")
	id := strings.TrimSpace(split1[0])

	split2 := strings.Split(split1[1], ",")

	left := strings.TrimSpace(strings.Replace(split2[0], "(", "", 1))
	right := strings.TrimSpace(strings.Replace(split2[1], ")", "", 1))
	return Node{id: id, left: left, right: right, endsWith: rune(id[lastCharPos])}
}

func nodeSliceComparator(a, b Node) int {
	return cmp.Compare(a.id, b.id)
}

func addGhostNode(gp GhostPath, gn GhostNode, direction bool) (out GhostPath) {
	return GhostPath{start: gp.start, end: gn, steps: append(gp.steps, direction)}
}

func enumeratePaths(n NodeMap, start GhostNode) (out []GhostPath) {
	seenNodes := make(map[GhostNode]int)

	// init path queue with a path from start -> start with 0 steps
	pathConsiderationQueue := []GhostPath{{start: start, end: start, steps: nil}}

	// I don't want `range pathConsiderationQueue` here because I'm going to be expanding it
	// so I want to check len() each time
	for queueIdx := 0; queueIdx < len(pathConsiderationQueue); queueIdx++ {
		currentPath := pathConsiderationQueue[queueIdx]
		seenNodes[currentPath.end] = 1

		leftNode := currentPath.end.left(n)
		leftPath := addGhostNode(currentPath, leftNode, leftBool)

		rightNode := currentPath.end.right(n)
		rightPath := addGhostNode(currentPath, rightNode, rightBool)

		// if we've seen this GhostNode before, then we've completed a path
		// add this traversal to output
		if seenNodes[leftNode] > 0 {
			out = append(out, leftPath)
		} else {
			// add to queue
			pathConsiderationQueue = append(pathConsiderationQueue, leftPath)
		}

		if seenNodes[rightNode] > 0 {
			out = append(out, rightPath)
		} else {
			pathConsiderationQueue = append(pathConsiderationQueue, rightPath)
		}
	}

	return
}

func parseInput(lines []string) (moves []bool, n NodeMap, start GhostNode) {
	for _, char := range lines[0] {
		moves = append(moves, char == leftChar)
	}

	n = make(NodeMap)
	startSlice := make([]Node, 0, numGhosts)
	for _, line := range lines[2:] {
		node := parseNode(line)
		if node.endsWith == startRune {
			startSlice = append(startSlice, node)
		}
		n[node.id] = node
	}
	slices.SortFunc(startSlice, nodeSliceComparator)
	copy(start[:], startSlice)
	return moves, n, start
}

func Part2(lines []string) int {
	_, n, start := parseInput(lines)

	// idea1: enumerate all of the possible outcomes (R, L, RR, RL, LR, LL, etc) and find the first that gets to all Z?
	paths := enumeratePaths(n, start)
	for _, p := range paths {
		fmt.Println(p)
	}

	// idea2: find all paths from A^6 to Z^6, and all cycles from A^6 and Z^6 back to themselves

	// idea3: include the history of "seen" nodes in a sequence so they can find cycles

	return 0
}
