package main

import (
	"advent/util"
	"fmt"
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

func splitAllByOperational(in []string) (out [][]int) {
	for _, springs := range in {
		split := splitByOperational(springs)
		counts := make([]int, 0)
		for _, segment := range split {
			counts = append(counts, len(segment))
		}
		out = append(out, counts)
	}
	return
}

func generateAll(springs string) (allPossible []string) {
	var s strings.Builder
	for idx, char := range springs {
		if char == unknown {
			damagedPrefix := s.String() + string(damaged)
			operationalPrefix := s.String() + string(operational)
			nextCharIdx := idx + 1
			if nextCharIdx == len(springs) {
				// there are no more chars
				return []string{damagedPrefix, operationalPrefix}
			} else {
				allPossibleRest := generateAll(springs[nextCharIdx:])
				allPossible = make([]string, len(allPossibleRest)*2)
				for restIdx, restString := range allPossibleRest {
					allPossible[restIdx] = damagedPrefix + restString
					allPossible[restIdx+len(allPossibleRest)] = operationalPrefix + restString
				}
				return allPossible
			}
		} else {
			s.WriteString(string(char))
		}
	}

	// if we have reached here, we have a single unambiguous string.
	return []string{s.String()}
}

func matches(allCounts [][]int, groups []int) (total int) {
	for _, counts := range allCounts {
		if len(counts) == len(groups) {
			match := true
			for idx, count := range counts {
				if count != groups[idx] {
					match = false
				}
			}
			if match {
				total++
			}
		}
	}
	return
}

func day12part1(rows []string) (total int) {
	for _, row := range rows {
		split := strings.Fields(row)

		damagedGroups, err := util.ParseSeparatedInts(split[1], ",")
		util.MaybePanic(err)

		allPossible := generateAll(split[0])
		allSplitCounts := splitAllByOperational(allPossible)

		rowMatches := matches(allSplitCounts, damagedGroups)
		total += rowMatches
	}
	return
}

func main() {
	rows := util.ReadInput("input", "12")
	result := day12part1(rows)
	fmt.Println("Part1", result)
}
