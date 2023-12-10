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

func (gn GhostNode) allTarget() bool {
	for idx := 0; idx < numGhosts; idx++ {
		if gn[idx].endsWith != targetRune {
			return false
		}
	}
	return true
}

type GhostPath struct {
	seen   map[GhostNode]int
	start  GhostNode
	end    GhostNode
	steps  string
	cycle  bool
	winner bool
}

func (gp GhostPath) String() string {
	return fmt.Sprintf("%s -> %s in %d steps:%s", gp.start.String(), gp.end.String(), len(gp.steps), gp.steps)
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

func addGhostNode(gp GhostPath, gn GhostNode, direction string) (out GhostPath) {
	cycle, winner := false, false
	seen := gp.seen
	if gn.allTarget() {
		winner = true
	} else if gp.seen[gn] > 0 {
		cycle = true
	} else {
		seen[gn] = 1
	}
	return GhostPath{seen: seen, start: gp.start, end: gn, steps: gp.steps + direction, cycle: cycle, winner: winner}
}

func enumeratePaths(n NodeMap, start GhostNode) (cycles, winners []GhostPath) {
	// init path queue with path at the start with 0 steps
	initSeen := make(map[GhostNode]int)
	initSeen[start] = 1
	pathConsiderationQueue := []GhostPath{{seen: initSeen, start: start, end: start}} //, steps: "", complete: false}}

	// I don't want `range pathConsiderationQueue` here because I'm going to be expanding it
	// so I want to check len() each time
	for queueIdx := 0; queueIdx < len(pathConsiderationQueue); queueIdx++ {
		currentPath := pathConsiderationQueue[queueIdx]
		leftPath := addGhostNode(currentPath, currentPath.end.left(n), "L")
		rightPath := addGhostNode(currentPath, currentPath.end.right(n), "R")

		for _, p := range []GhostPath{leftPath, rightPath} {
			if p.winner {
				winners = append(winners, p)
			} else if p.cycle {
				cycles = append(cycles, p)
			} else {
				// add to queue
				pathConsiderationQueue = append(pathConsiderationQueue, p)
			}
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
	cycles, winners := enumeratePaths(n, start)
	fmt.Println("Winners")
	for _, p := range winners {
		fmt.Println(p)
	}
	fmt.Println("Cycles")
	for _, p := range cycles {
		fmt.Println(p)
	}

	// idea2: find all paths from A^6 to Z^6, and all cycles from A^6 and Z^6 back to themselves

	// idea3: include the history of "seen" nodes in a sequence so they can find cycles

	return 0
}
