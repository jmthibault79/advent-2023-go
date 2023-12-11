package day8

import (
	"advent/util"
	"fmt"
	"strings"
)

const lastCharPos = 2
const startRune = 'A'
const targetRune = 'Z'
const leftChar = 'L'

// how many times should we cycle through before we're convinced the result is periodic?
const periodicConfirmation = 3

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
	seen         map[string]int
	start        string
	end          string
	steps        string
	cycle        bool
	targetSteps  []int
	targetAtStep map[int]string
	targetCounts map[string]int
}

func (p Path) String() string {
	return fmt.Sprintf("%s -> %s in %d steps", p.start, p.end, len(p.steps))
	//	return fmt.Sprintf("%s -> %s in %d steps: %s", p.start, p.end, len(p.steps), p.steps)
}

func (p Path) maxTargetCount() (maxCount int) {
	maxCount = 0
	for _, count := range p.targetCounts {
		maxCount = max(maxCount, count)
	}
	return
}

func (p Path) getPeriod() (period int) {
	period = p.targetSteps[0]
	for idx := 1; idx > len(p.targetSteps); idx++ {
		thisPeriod := p.targetSteps[idx] - p.targetSteps[idx-1]
		if thisPeriod != period {
			panic("not periodic!")
		}
	}
	return
}

func addNode(p Path, n *Node, direction string) (out Path) {

	// init

	cycle := false
	steps := p.steps + direction
	seen := make(map[string]int)
	targetSteps := make([]int, 0)
	targetAtStep := make(map[int]string)
	targetCounts := make(map[string]int)

	// copy maps and slices - looks like there are reference problems if you don't
	for k, v := range p.seen {
		seen[k] = v
	}
	for k, v := range p.targetAtStep {
		targetAtStep[k] = v
	}
	for k, v := range p.targetCounts {
		targetCounts[k] = v
	}
	for _, step := range p.targetSteps {
		targetSteps = append(targetSteps, step)
	}

	if n.endsWith == targetRune {
		step := len(steps)
		targetSteps = append(targetSteps, step)
		targetAtStep[step] = n.id
		targetCounts[n.id]++
	}

	if p.seen[n.id] > 0 {
		cycle = true
	} else {
		seen[n.id] = 1
	}
	return Path{seen: seen, start: p.start, end: n.id, steps: steps, cycle: cycle,
		targetSteps: targetSteps, targetAtStep: targetAtStep, targetCounts: targetCounts}
}

func findPaths(startNodes []Node, m NodeMap, moves string) (paths []Path) {
	//	n := startNodes[0]
	for _, n := range startNodes {
		initSeen := make(map[string]int)
		initSeen[n.id] = 1
		path := Path{seen: initSeen, start: n.id, end: n.id}
		movIdx := 0
		for path.maxTargetCount() < periodicConfirmation {
			if moves[movIdx%len(moves)] == leftChar {
				path = addNode(path, m[m[path.end].left], "L")
			} else {
				path = addNode(path, m[m[path.end].right], "R")
			}
			movIdx++
		}
		paths = append(paths, path)
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

func parseInput(lines []string) (moves string, m NodeMap, startNodes []Node) {
	m = make(NodeMap)
	for _, line := range lines[2:] {
		node := parseNode(line)
		if node.endsWith == startRune {
			startNodes = append(startNodes, node)
		}
		m[node.id] = &node
	}

	return strings.TrimSpace(lines[0]), m, startNodes
}

func Part2(lines []string) int {
	moves, m, startNodes := parseInput(lines)
	paths := findPaths(startNodes, m, moves)
	periods := make([]int, 0)
	for _, path := range paths {
		period := path.getPeriod()
		periods = append(periods, period)
		fmt.Println(path.start, "with period", period)
	}
	return util.LCM(periods)
}
