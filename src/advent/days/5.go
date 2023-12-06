package main

import (
	"advent/util"
	"fmt"
	"strings"
)

// example -> seeds: 79 14 55 13
func parseSeeds(seedLine string) (seeds []int, err error) {
	return util.ParseSpacedInts(strings.Replace(seedLine, "seeds: ", "", 1))
}

// example -> seed-to-soil map:
func tryParseMapHeader(maybeHeaderLine string) (from, to string, success bool) {
	firstSplit := strings.Fields(maybeHeaderLine)
	if len(firstSplit) != 2 {
		return "", "", false
	}
	secondSplit := strings.Split(firstSplit[0], "-")
	if len(secondSplit) != 3 || secondSplit[1] != "to" {
		return "", "", false
	}
	return secondSplit[0], secondSplit[2], true
}

func tryParseMapRow(maybeRowLine string) (srcStart, destStart, rangeLength int, success bool) {
	vals, err := util.ParseSpacedInts(maybeRowLine)
	if err != nil || len(vals) != 3 {
		return 0, 0, 0, false
	} else {
		return vals[0], vals[1], vals[2], true
	}
}

type ConversionMap struct {
	fromType string
	toType   string
	values   [][]int
}

func (m ConversionMap) String() string {
	var out strings.Builder

	out.WriteString(fmt.Sprintf("Map from %s to %s\n", m.fromType, m.toType))
	for _, row := range m.values {
		out.WriteString(util.IntsToString(row) + "\n")
	}

	return out.String()
}

func day5part1(lines []string) int {
	seeds, err := parseSeeds(lines[0])
	util.MaybePanic(err)
	fmt.Println("Seeds", seeds)

	var maps []ConversionMap

	// map builder values
	rowValues := make([][]int, 0)
	inMap, from, to := false, "", ""

	// skip blank line 1
	for _, line := range lines[2:] {
		if inMap {
			srcStart, destStart, rangeLength, success := tryParseMapRow(line)
			if success {
				rowValues = append(rowValues, []int{srcStart, destStart, rangeLength})
			} else {
				// complete the map
				maps = append(maps, ConversionMap{fromType: from, toType: to, values: rowValues})
				// reset
				inMap, from, to = false, "", ""
				rowValues = make([][]int, 0)
			}
		} else {
			from, to, inMap = tryParseMapHeader(line)
		}
	}

	// complete the last map if we haven't already
	if inMap {
		maps = append(maps, ConversionMap{fromType: from, toType: to, values: rowValues})
	}

	fmt.Println(len(maps), "Maps")
	for _, m := range maps {
		fmt.Println(m)
	}

	return 0
}

func main() {
	inputFile := "../../inputs/test5.txt"
	lines := util.ReadLines(inputFile)

	// ok wow day 5 is special
	// there are these maps, see, and you use them to convert numbers of category X to category Y

	result := day5part1(lines)
	fmt.Println("Part1", result)

}
