package day8

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
)

const lastCharPos = 2
const startCollectionRune = 'A'
const targetCollectionRune = 'Z'
const leftChar = 'L'
const leftBool = true
const rightBool = false

// empirical values from input
const nodesInCollection = 2

type Node struct {
	id       string
	endsWith rune
	left     string
	right    string
}
type GhostNode [nodesInCollection]Node
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
	lefts := make([]Node, 0, nodesInCollection)
	for _, n := range gn {
		lefts = append(lefts, m[n.left])
	}
	slices.SortFunc(lefts, nodeSliceComparator)
	copy(out[:], lefts)
	return
}

func (gn GhostNode) right(m NodeMap) (out GhostNode) {
	rights := make([]Node, 0, nodesInCollection)
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
	seenNodeCollections := make(map[GhostNode]int)

	// init traversal queue with the startNodes
	traversalQueue := []GhostPath{{start: start, end: start, steps: nil}}

	// I don't want `range traversalQueue` here because I'm going to be expanding it
	// so I want to check len() each time
	for queueIdx := 0; queueIdx < len(traversalQueue); queueIdx++ {
		nct := traversalQueue[queueIdx]
		seenNodeCollections[nct.end] = 1

		leftNC := nct.end.left(n)
		leftTraversal := addGhostNode(nct, leftNC, leftBool)

		// if we've seen this GhostNode before, then we've completed a cycle
		// add this traversal to output
		if seenNodeCollections[leftNC] > 0 {
			out = append(out, leftTraversal)
		} else {
			// add to queue
			traversalQueue = append(traversalQueue, leftTraversal)
		}

		rightNC := traversalQueue[queueIdx].end.right(n)
		rightTraversal := addGhostNode(nct, rightNC, rightBool)

		if seenNodeCollections[rightNC] > 0 {
			out = append(out, rightTraversal)
		} else {
			traversalQueue = append(traversalQueue, rightTraversal)
		}
	}

	return
}

func parseInput(lines []string) (moves []bool, n NodeMap, startNodes GhostNode) {
	for _, char := range lines[0] {
		moves = append(moves, char == leftChar)
	}

	n = make(NodeMap)
	startSlice := make([]Node, 0, nodesInCollection)
	for _, line := range lines[2:] {
		node := parseNode(line)
		if node.endsWith == startCollectionRune {
			startSlice = append(startSlice, node)
		}
		n[node.id] = node
	}
	slices.SortFunc(startSlice, nodeSliceComparator)
	copy(startNodes[:], startSlice)
	return moves, n, startNodes
}

func Part2(lines []string) int {
	_, n, startNodes := parseInput(lines)

	ts := enumeratePaths(n, startNodes)
	for _, t := range ts {
		fmt.Println(t)
	}
	// idea1: enumerate all of the possible outcomes (R, L, RR, RL, LR, LL, etc) and find the first that gets to all Z?

	// idea2: find all paths from A^6 to Z^6, and all cycles from A^6 and Z^6 back to themselves

	// idea3: include the history of "seen" nodes in a sequence so they can find cycles

	return 0
}
