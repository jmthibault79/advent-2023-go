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

	// get last group
	if b.Len() > 0 {
		out = append(out, b.String())
	}

	return
}

// assume our splitStrings is ???
// groupLength 1 -> 3 ways
// groupLength 2 -> 2 ways
// groupLength 3 -> 1 ways
// groupLength 4 -> 0 ways
func possibleMatches(splitSprings string, groupLength int) int {
	return int(math.Max(0, float64(len(splitSprings)-groupLength+1)))
}

// 1,2,3 -> len(#.##.###)
func minMatchLength(groups []int) (out int) {
	out = groups[0]
	for _, group := range groups[1:] {
		out += 1 + group
	}
	return
}

func matchOne(s string, groups []int) int {
	if minMatchLength(groups) == len(s) {
		return 1
	} else {
		// I dunno man
		return -1
	}
}

// rows are independent so here's where all the logic is
func oneRow(splitSprings []string, damagedGroups []int) (total int) {
	// general idea for approach: divide and conquer
	// if we can match the beginning and/or the end, we have a smaller problem?

	// assert impossibilities
	if len(splitSprings) == 0 {
		panic("0 splits")
	} else if len(damagedGroups) == 0 {
		panic("0 damagedGroups")
	} else if len(splitSprings) > len(damagedGroups) {
		panic(fmt.Sprintf("Springs split into %d groups but we were given %d",
			len(splitSprings), len(damagedGroups)))
	}

	if len(splitSprings) == len(damagedGroups) {
		// match each separately, and multiply the results
		acc := possibleMatches(splitSprings[0], damagedGroups[0])
		for idx := 1; idx < len(splitSprings); idx++ {
			acc *= possibleMatches(splitSprings[idx], damagedGroups[idx])
		}
		return acc
	} else if len(splitSprings) == 1 {
		return matchOne(splitSprings[0], damagedGroups)
	} else {
		firstSplit, firstGroup := splitSprings[0], damagedGroups[0]
		lastSplit, lastGroup := splitSprings[len(splitSprings)-1], damagedGroups[len(damagedGroups)-1]

		// if the first of each or last of each has matching length, we can use that
		if len(firstSplit) == firstGroup {
			firstMatches := possibleMatches(firstSplit, firstGroup)
			restMatches := oneRow(splitSprings[1:], damagedGroups[1:])
			return firstMatches * restMatches
		}
		if len(lastSplit) == lastGroup {
			lastMatches := possibleMatches(lastSplit, lastGroup)
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
