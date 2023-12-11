package day8

import (
	"fmt"
	"strings"
)

const lastCharPos = 2
const startRune = 'A'
const targetRune = 'Z'
const leftChar = 'L'

type Node struct {
	id       string
	endsWith rune
	left     string
	right    string
}

func (n Node) String() string {
	return fmt.Sprintf("%s (%s, %s)", n.id, n.left, n.right)
}

type NodeMap map[string]*Node

type Path struct {
	seen        map[string]int
	start       string
	end         string
	steps       string
	cycle       bool
	targetSteps map[string]int
}

func (gp Path) String() string {
	return fmt.Sprintf("%s -> %s in %d steps: %s", gp.start, gp.end, len(gp.steps), gp.steps)
}

func addNode(p Path, n *Node, direction string) (out Path) {

	// init

	seen := make(map[string]int)
	steps := p.steps + direction
	cycle := false
	targetSteps := make(map[string]int)

	// copy targetSteps map - looks like there are reference problems if you don't
	for k, v := range p.targetSteps {
		targetSteps[k] = v
	}

	if n.endsWith == targetRune {
		targetSteps[n.id] = len(p.steps)
	}

	if p.seen[n.id] > 0 {
		cycle = true
	} else {
		// copy seen map - looks like there are reference problems if you don't
		for k, v := range p.seen {
			seen[k] = v
		}
		seen[n.id] = 1
	}
	return Path{seen: seen, start: p.start, end: n.id, steps: steps, cycle: cycle, targetSteps: targetSteps}
}

//
//func enumeratePaths(m NodeMap, start GhostNode) (cycles, winners []GhostPath) {
//	// init path queue with path at the start with 0 steps
//	initSeen := make(map[GhostNode]int)
//	initSeen[start] = 1
//	pathConsiderationQueue := []GhostPath{{seen: initSeen, start: start, end: start}}
//
//	// I don't want `range pathConsiderationQueue` here because I'm going to be expanding it
//	// so I want to check len() each time
//	for queueIdx := 0; queueIdx < len(pathConsiderationQueue); queueIdx++ {
//		if queueIdx%1_000_000 == 0 {
//			fmt.Println("idx ", queueIdx, " len ", len(pathConsiderationQueue))
//		}
//		currentPath := pathConsiderationQueue[queueIdx]
//		leftPath := addGhostNode(currentPath, currentPath.end.left(m), "L", m)
//		rightPath := addGhostNode(currentPath, currentPath.end.right(m), "R", m)
//
//		// attempt to help the garbage collector by removing something we don't need
//		pathConsiderationQueue[queueIdx] = GhostPath{}
//
//		for _, p := range []GhostPath{leftPath, rightPath} {
//			if p.winner {
//				fmt.Println("Winner", p)
//				winners = append(winners, p)
//			} else if p.cycle {
//				if len(cycles)%100 == 0 {
//					fmt.Println("Cycles", len(cycles))
//				}
//				cycles = append(cycles, p)
//			} else {
//				// add to queue
//				pathConsiderationQueue = append(pathConsiderationQueue, p)
//			}
//		}
//	}
//
//	return
//}

func findWinners(p Path, m NodeMap) (winners []Path) {
	if p.cycle {
		if len(p.targetSteps) > 0 {
			return []Path{p}
		} else {
			return nil
		}
	} else {
		leftWinners := findWinners(addNode(p, m[m[p.end].left], "L"), m)
		winners = append(winners, leftWinners...)
		rightWinners := findWinners(addNode(p, m[m[p.end].right], "R"), m)
		winners = append(winners, rightWinners...)
		return
	}
}

func findWinningPaths(startNodes []Node, m NodeMap) (allWinners []Path) {
	//	n := startNodes[0]
	for _, n := range startNodes {
		initSeen := make(map[string]int)
		initSeen[n.id] = 1
		initPath := Path{seen: initSeen, start: n.id, end: n.id}
		allWinners = append(allWinners, findWinners(initPath, m)...)
	}
	return
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

func parseInput(lines []string) (moves []bool, m NodeMap, startNodes []Node) {
	for _, char := range lines[0] {
		moves = append(moves, char == leftChar)
	}

	m = make(NodeMap)
	for _, line := range lines[2:] {
		node := parseNode(line)
		if node.endsWith == startRune {
			startNodes = append(startNodes, node)
		}
		m[node.id] = &node
	}

	return moves, m, startNodes
}

func Part2(lines []string) int {

	// idea1: enumerate all of the possible outcomes (R, L, RR, RL, LR, LL, etc) and find the first that gets to all Z?

	// idea2: find all paths from A^6 to Z^6, and all cycles from A^6 and Z^6 back to themselves

	// idea3: include the history of "seen" nodes in a sequence so they can find cycles

	//cycles, winners := enumeratePaths(m, start)
	//fmt.Println("Winners:", len(winners))
	//for _, p := range winners {
	//	fmt.Println(p)
	//}
	//fmt.Println("Cycles:", len(cycles))
	//for _, p := range cycles {
	//	fmt.Println(p)
	//}

	_, m, startNodes := parseInput(lines)
	winners := findWinningPaths(startNodes, m)
	if len(winners) == 0 {
		fmt.Println("No Winners")
	}
	for _, winner := range winners {
		fmt.Println("Winner:", winner)
	}
	return 0
}
