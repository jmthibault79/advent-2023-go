package main

import (
	"advent/util"
	"fmt"
)

const splitVert = '|'
const splitHoriz = '-'
const nwseMirror = '\\'
const neswMirror = '/'
const passThrough = '.'

// arbitrary
const fromWest = 1
const fromNorth = 2
const fromEast = 3
const fromSouth = 4

// this day calls for globals
var day16MirrorMap []string
var day16Seen [][]map[int]bool

func inBounds(rowIdx, colIdx int) bool {
	return rowIdx >= 0 && rowIdx < len(day16MirrorMap) && colIdx >= 0 && colIdx < len(day16MirrorMap[0])
}

func moveStraight(fromDir, rowIdx, colIdx int) (newRI, newCI int, keepGoing bool) {
	newRI, newCI = -1, -1

	switch fromDir {
	case fromWest:
		newRI, newCI = rowIdx, colIdx+1
	case fromEast:
		newRI, newCI = rowIdx, colIdx-1
	case fromNorth:
		newRI, newCI = rowIdx+1, colIdx
	case fromSouth:
		newRI, newCI = rowIdx-1, colIdx
	}
	return newRI, newCI, inBounds(newRI, newCI)
}

func bounceNesw(fromDir, rowIdx, colIdx int) (newRI, newCI, newDir int, keepGoing bool) {
	newRI, newCI = -1, -1

	switch fromDir {
	case fromWest:
		newRI, newCI, newDir = rowIdx-1, colIdx, fromSouth
	case fromEast:
		newRI, newCI, newDir = rowIdx+1, colIdx, fromNorth
	case fromNorth:
		newRI, newCI, newDir = rowIdx, colIdx-1, fromEast
	case fromSouth:
		newRI, newCI, newDir = rowIdx, colIdx+1, fromWest
	}

	return newRI, newCI, newDir, inBounds(newRI, newCI)
}

func bounceNwse(fromDir, rowIdx, colIdx int) (newRI, newCI, newDir int, keepGoing bool) {
	newRI, newCI = -1, -1

	switch fromDir {
	case fromWest:
		newRI, newCI, newDir = rowIdx+1, colIdx, fromNorth
	case fromEast:
		newRI, newCI, newDir = rowIdx-1, colIdx, fromSouth
	case fromNorth:
		newRI, newCI, newDir = rowIdx, colIdx+1, fromWest
	case fromSouth:
		newRI, newCI, newDir = rowIdx, colIdx-1, fromEast
	}

	return newRI, newCI, newDir, inBounds(newRI, newCI)
}

func maybeSplitVert(fromDir, rowIdx, colIdx int) (newRI, newCI int, keepGoing bool) {
	switch fromDir {
	case fromNorth:
		fallthrough
	case fromSouth:
		return moveStraight(fromDir, rowIdx, colIdx)
	case fromWest:
		fallthrough
	case fromEast:
		// we don't want to return keepGoing here: instead fire two new lasers north and south, if possible
		for _, newFromDir := range []int{fromSouth, fromNorth} {
			newRI, newCI, keepGoing = moveStraight(newFromDir, rowIdx, colIdx)
			if keepGoing {
				pewPew(newFromDir, newRI, newCI)
			}
		}
	}

	return -1, -1, false
}

func maybeSplitHoriz(fromDir, rowIdx, colIdx int) (newRI, newCI int, keepGoing bool) {
	switch fromDir {
	case fromWest:
		fallthrough
	case fromEast:
		return moveStraight(fromDir, rowIdx, colIdx)
	case fromNorth:
		fallthrough
	case fromSouth:
		// we don't want to return keepGoing here: instead fire two new lasers east and west, if possible
		for _, newFromDir := range []int{fromEast, fromWest} {
			newRI, newCI, keepGoing = moveStraight(newFromDir, rowIdx, colIdx)
			if keepGoing {
				pewPew(newFromDir, newRI, newCI)
			}
		}
	}

	return -1, -1, false
}

func pewPew(fromDir, rowIdx, colIdx int) {
	// have we already shot a laser at this space from this direction?
	if day16Seen[rowIdx][colIdx][fromDir] {
		return
	} else {
		day16Seen[rowIdx][colIdx][fromDir] = true
	}

	newRI, newCI, newDir, keepGoing := -1, -1, fromDir, false
	switch day16MirrorMap[rowIdx][colIdx] {
	case passThrough:
		newRI, newCI, keepGoing = moveStraight(fromDir, rowIdx, colIdx)
	case neswMirror:
		newRI, newCI, newDir, keepGoing = bounceNesw(fromDir, rowIdx, colIdx)
	case nwseMirror:
		newRI, newCI, newDir, keepGoing = bounceNwse(fromDir, rowIdx, colIdx)
	case splitVert:
		newRI, newCI, keepGoing = maybeSplitVert(fromDir, rowIdx, colIdx)
	case splitHoriz:
		newRI, newCI, keepGoing = maybeSplitHoriz(fromDir, rowIdx, colIdx)
	}

	if keepGoing {
		pewPew(newDir, newRI, newCI)
	}
}

func day16part1() (total int) {
	// init day16Seen
	day16Seen = make([][]map[int]bool, len(day16MirrorMap))
	for rowIdx, _ := range day16MirrorMap {
		day16Seen[rowIdx] = make([]map[int]bool, len(day16MirrorMap[0]))
		for colIdx, _ := range day16MirrorMap[0] {
			day16Seen[rowIdx][colIdx] = make(map[int]bool, 0)
		}
	}

	// shoot a laser at 0,0 from the west
	pewPew(fromWest, 0, 0)

	// total the energized spaces
	for _, r := range day16Seen {
		for _, seenMap := range r {
			if len(seenMap) > 0 {
				total++
			}
		}
	}

	return
}

func main() {
	day16MirrorMap = util.ReadInput("input", "16")
	result := day16part1()
	fmt.Println("Part1", result)
}
