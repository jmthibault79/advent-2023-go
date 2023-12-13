package main

import (
	"advent/util"
	"fmt"
)

const snek = 'S'
const ns = '|'
const ew = '-'
const ne = 'L'
const nw = 'J'
const sw = '7'
const se = 'F'
const ground = '.'

const north = 1
const south = 2
const east = 3
const west = 4

type MapPoint struct {
	x             int
	y             int
	value         uint8
	stepToGetHere uint8
}

func (p MapPoint) String() string {
	return fmt.Sprintf("%s %d,%d", string(p.value), p.y, p.x)
}

// my coordinate system starts at 0,0 in the nw and max,max in the se
// x is within a line and y is the lines

func westOf(srcX, srcY int, pipeMap []string) (point MapPoint, compatible bool) {
	if srcX > 0 {
		x, y := srcX-1, srcY
		westPoint := pipeMap[y][x]
		return MapPoint{x: x, y: y, value: westPoint, stepToGetHere: west},
			westPoint == snek || westPoint == ew || westPoint == ne || westPoint == se
	}
	return MapPoint{x: -1, y: -1, value: ground}, false
}

func eastOf(srcX, srcY int, pipeMap []string) (point MapPoint, compatible bool) {
	if srcX < len(pipeMap[0])-1 {
		x, y := srcX+1, srcY
		eastPoint := pipeMap[y][x]
		return MapPoint{x: x, y: y, value: eastPoint, stepToGetHere: east},
			eastPoint == snek || eastPoint == ew || eastPoint == nw || eastPoint == sw
	}
	return MapPoint{x: -1, y: -1, value: ground}, false
}

func northOf(srcX, srcY int, pipeMap []string) (point MapPoint, compatible bool) {
	if srcY > 0 {
		x, y := srcX, srcY-1
		northPoint := pipeMap[y][x]
		return MapPoint{x: x, y: y, value: northPoint, stepToGetHere: north},
			northPoint == snek || northPoint == ns || northPoint == se || northPoint == sw
	}
	return MapPoint{x: -1, y: -1, value: ground}, false
}

func southOf(srcX, srcY int, pipeMap []string) (point MapPoint, compatible bool) {
	if srcY < len(pipeMap)-1 {
		x, y := srcX, srcY+1
		southPoint := pipeMap[y][x]
		return MapPoint{x: x, y: y, value: southPoint, stepToGetHere: south},
			southPoint == snek || southPoint == ns || southPoint == ne || southPoint == nw
	}
	return MapPoint{x: -1, y: -1, value: ground}, false
}

type directionChecker func(int, int, []string) (MapPoint, bool)

func nextStep(fromPoint MapPoint, pipeMap []string) MapPoint {
	directionsToTry := make([]directionChecker, 0)

	// 3 possibilities:
	// starting from snek: can be any direction
	// ns or ew: continue going that direction
	// nw,sw,se,ne: go around the corner

	switch fromPoint.value {
	case snek:
		directionsToTry = []directionChecker{northOf, westOf, southOf, eastOf}

		// straight through

	case ew:
		if fromPoint.stepToGetHere == west {
			directionsToTry = []directionChecker{westOf}
		} else if fromPoint.stepToGetHere == east {
			directionsToTry = []directionChecker{eastOf}
		}
	case ns:
		if fromPoint.stepToGetHere == north {
			directionsToTry = []directionChecker{northOf}
		} else if fromPoint.stepToGetHere == south {
			directionsToTry = []directionChecker{southOf}
		}

		// corners

	case nw:
		if fromPoint.stepToGetHere == south {
			directionsToTry = []directionChecker{westOf}
		} else if fromPoint.stepToGetHere == east {
			directionsToTry = []directionChecker{northOf}
		}
	case sw:
		if fromPoint.stepToGetHere == north {
			directionsToTry = []directionChecker{westOf}
		} else if fromPoint.stepToGetHere == east {
			directionsToTry = []directionChecker{southOf}
		}
	case se:
		if fromPoint.stepToGetHere == north {
			directionsToTry = []directionChecker{eastOf}
		} else if fromPoint.stepToGetHere == west {
			directionsToTry = []directionChecker{southOf}
		}
	case ne:
		if fromPoint.stepToGetHere == south {
			directionsToTry = []directionChecker{eastOf}
		} else if fromPoint.stepToGetHere == west {
			directionsToTry = []directionChecker{northOf}
		}
	}

	for _, dtt := range directionsToTry {
		point, compatible := dtt(fromPoint.x, fromPoint.y, pipeMap)
		if compatible {
			return point
		}
	}

	panic("no path for snek")
}

func day10part1(pipeMap []string) (out int) {
	// ok first find that S
	Sx, Sy := -1, -1
snekloop:
	for yi := 0; yi < len(pipeMap); yi++ {
		for xi := 0; xi < len(pipeMap[yi]); xi++ {
			if pipeMap[yi][xi] == snek {
				Sx, Sy = xi, yi
				break snekloop
			}
		}
	}

	snekPoint := MapPoint{x: Sx, y: Sy, value: snek}
	pipePath := []MapPoint{snekPoint}

	// assumption: there won't be any compatible but false steps

	step := nextStep(snekPoint, pipeMap)
	pipePath = append(pipePath, step)
	for step.value != snek {
		step = nextStep(step, pipeMap)
		pipePath = append(pipePath, step)
	}

	for _, point := range pipePath {
		fmt.Println(point)
	}
	return len(pipePath) / 2
}

func main() {
	pipeMap := util.ReadInput("input", "10")
	result := day10part1(pipeMap)
	fmt.Println("Part1", result)
	//
	//pipeMap = util.ReadInput("test", "10b")
	//result = day10part1(pipeMap)
	//fmt.Println("Part1b", result)
}
