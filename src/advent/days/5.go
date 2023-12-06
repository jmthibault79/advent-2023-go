package main

import (
	"advent/util"
	"fmt"
	"math"
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

func tryParseMapRow(maybeRowLine string) ([]int, bool) {
	vals, err := util.ParseSpacedInts(maybeRowLine)
	if err != nil || len(vals) != 3 {
		return nil, false
	} else {
		return vals, true
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

func (m ConversionMap) convert(in int) int {
	for _, row := range m.values {
		// let's follow along with this nonsense by using an example: ds = 10, ss = 20, rl = 30
		// input matches are therefore 20-49, and should be mapped to 10-39
		destStart, srcStart, rangeLength := row[0], row[1], row[2]
		// for matches, this value would be in the range [0, 29]
		deltaFromStart := in - srcStart
		if deltaFromStart >= 0 && deltaFromStart < rangeLength {
			// so matching outputs are in the range [10, 39]
			return destStart + deltaFromStart
		}
	}
	// no matches: return original
	return in
}

func day5part1(lines []string) int {
	seeds, err := parseSeeds(lines[0])
	util.MaybePanic(err)
	fmt.Println("Seeds", seeds)

	// skip blank line 1
	maps := parseMaps(lines[2:])

	lowest := math.MaxInt

	for _, seed := range seeds {
		fmt.Print("seed ", seed)
		val := seed
		for _, m := range maps {
			val = m.convert(val)
			fmt.Print(" -> ", m.toType, " ", val)
		}
		fmt.Println()
		lowest = min(lowest, val)
	}

	return lowest
}

func parseMaps(mapLines []string) (maps []ConversionMap) {
	// map builder values
	rowValues := make([][]int, 0)
	inMap, from, to := false, "", ""

	for _, line := range mapLines {
		if inMap {
			row, success := tryParseMapRow(line)
			if success {
				rowValues = append(rowValues, row)
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
	return maps
}

func day5part2(lines []string) int {
	part1Seeds, err := parseSeeds(lines[0])
	util.MaybePanic(err)
	fmt.Println("Seeds", part1Seeds)

	var seeds []int
	for seedIdx := 0; seedIdx < len(part1Seeds); seedIdx += 2 {
		start := part1Seeds[seedIdx]
		rl := part1Seeds[seedIdx+1]
		for newSeedIdx := start; newSeedIdx < start+rl; newSeedIdx++ {
			seeds = append(seeds, newSeedIdx)
			if len(seeds)%1_000_000_000 == 0 {
				fmt.Println("Produced a billion seeds")
			}
		}
	}
	fmt.Println("New Seed Count", len(seeds))

	// skip blank line 1
	maps := parseMaps(lines[2:])

	lowest := math.MaxInt

	fmt.Println("Seed count", len(seeds))
	for idx, seed := range seeds {
		val := seed
		for _, m := range maps {
			val = m.convert(val)
		}
		lowest = min(lowest, val)
		if idx%1_000_000_000 == 0 {
			fmt.Println("Processed a billion seeds")
		}
	}

	return lowest
}

func main() {
	inputFile := "../../inputs/input5.txt"
	lines := util.ReadLines(inputFile)

	// ok wow day 5 is special
	// there are these maps, see, and you use them to convert numbers of category X to category Y

	result := day5part1(lines)
	fmt.Println("Part1", result)

	// part 2: now the seeds are ranges
	result = day5part2(lines)
	fmt.Println("Part2", result)
}
