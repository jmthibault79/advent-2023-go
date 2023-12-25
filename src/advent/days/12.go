package main

import (
	"advent/util"
	"fmt"
	"math"
	"strings"
)

const operational = '.'
const damaged = '#'
const unknown = '?'

func splitByOperational(in string) (out []string) {
	// runs of operational are equivalent to 1 - split on these
	var b strings.Builder
	for _, char := range in {
		if char == operational {
			// if we accumulated a string before this, ship it
			if b.Len() > 0 {
				out = append(out, b.String())
				b.Reset()
			}
		} else {
			b.WriteString(string(char))
		}
	}

	// ship last group
	if b.Len() > 0 {
		out = append(out, b.String())
	}

	return
}

// assume our springsLen is 3 (???, ?##, etc)
// groupLength 1 -> 3 ways
// groupLength 2 -> 2 ways
// groupLength 3 -> 1 ways
// groupLength 4 -> 0 ways
func groupInOneString(springsLen int, damagedGroupLen int) int {
	return int(math.Max(0, float64(springsLen-damagedGroupLen+1)))
}

// 1,2,3 -> len(#.##.###)
func minMatchLength(damagedGroups []int) (out int) {
	out = damagedGroups[0]
	for _, group := range damagedGroups[1:] {
		out += 1 + group
	}
	return
}

func manyGroupsOneString(springs string, damagedGroups []int) int {
	if len(damagedGroups) == 1 {
		return groupInOneString(len(springs), damagedGroups[0])
	}

	if len(damagedGroups) == 0 || minMatchLength(damagedGroups) > len(springs) {
		return 0
	}

	// can we match the first plus a gap?
	// before the gap, we always match because it's damaged or unknown
	firstDamagedGap := damagedGroups[0]
	if springs[firstDamagedGap] == damaged {
		// not a match - let's try again starting from the next position
		return manyGroupsOneString(springs[1:], damagedGroups)
	} else {
		// a match!  Do we still match if we check what's left?
		subMatches := manyGroupsOneString(springs[firstDamagedGap+1:], damagedGroups[1:])
		// add that to what happens if we shift this by one
		shiftMatches := manyGroupsOneString(springs[1:], damagedGroups)
		return subMatches + shiftMatches
	}
}

func oneRow(splitSprings []string, damagedGroups []int) (total int) {
	// general idea for approach: divide and conquer
	// if we can match the beginning and/or the end, we have a smaller problem?

	// assert impossibilities
	if len(splitSprings) == 0 {
		panic("0 splits")
	} else if len(damagedGroups) == 0 {
		panic("0 damagedGroups")
	} else if len(splitSprings) > len(damagedGroups) {
		panic(fmt.Sprintf("Springs split into %d damagedGroups but we were given %d",
			len(splitSprings), len(damagedGroups)))
	}

	if len(splitSprings) == len(damagedGroups) {
		// match each separately, and multiply the results
		acc := groupInOneString(len(splitSprings[0]), damagedGroups[0])
		for idx := 1; idx < len(splitSprings); idx++ {
			acc *= groupInOneString(len(splitSprings[idx]), damagedGroups[idx])
		}
		return acc
	} else if len(splitSprings) == 1 {
		return manyGroupsOneString(splitSprings[0], damagedGroups)
	} else {
		firstSplit, firstGroup := splitSprings[0], damagedGroups[0]
		lastSplit, lastGroup := splitSprings[len(splitSprings)-1], damagedGroups[len(damagedGroups)-1]

		// if the first of each or last of each has matching length, we can use that
		if len(firstSplit) == firstGroup {
			firstMatches := groupInOneString(len(firstSplit), firstGroup)
			restMatches := oneRow(splitSprings[1:], damagedGroups[1:])
			return firstMatches * restMatches
		}
		if len(lastSplit) == lastGroup {
			lastMatches := groupInOneString(len(lastSplit), lastGroup)
			restMatches := oneRow(splitSprings[:len(splitSprings)-1], damagedGroups[:len(damagedGroups)-1])
			return lastMatches * restMatches
		}

		// are there cases beyond these?
		return -1
	}
}

func day12part1(rows []string) (total int) {
	for _, row := range rows {
		split := strings.Fields(row)
		// what spring divisions do we have in our input string?
		splitSprings := splitByOperational(split[0])

		damagedGroups, err := util.ParseSeparatedInts(split[1], ",")
		util.MaybePanic(err)

		// rows are independent so here's where all the logic is
		rowMatches := oneRow(splitSprings, damagedGroups)
		total += rowMatches
	}
	return
}

func main() {
	rows := util.ReadInput("test", "12")
	result := day12part1(rows)
	fmt.Println("Part1", result)
}
