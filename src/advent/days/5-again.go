package main

import (
	"advent/util"
	"fmt"
	"sort"
	"strings"
)

const cmRowVals = 3

type RangedValue struct {
	start int
	end   int
	span  int
}

func (rv RangedValue) String() string {
	return fmt.Sprintf("%d-%d (%d)", rv.start, rv.end, rv.span)
}

func createRangedValue(start, span int) RangedValue {
	return RangedValue{start: start, span: span, end: start + span - 1}
}

type RVComparator []RangedValue

func (c RVComparator) Len() int           { return len(c) }
func (c RVComparator) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c RVComparator) Less(i, j int) bool { return c[i].start < c[j].start }

type ConversionMapRow struct {
	destStart int
	rv        RangedValue
}

func (c ConversionMapRow) String() string {
	return fmt.Sprintf("%s->%d", c.rv.String(), c.destStart)
}

type CMRComparator []ConversionMapRow

func (c CMRComparator) Len() int           { return len(c) }
func (c CMRComparator) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c CMRComparator) Less(i, j int) bool { return c[i].rv.start < c[j].rv.start }

func parseRangedSeeds(seedLine string) (seeds []RangedValue) {
	vals, err := util.ParseSpacedInts(strings.Replace(seedLine, "seeds: ", "", 1))
	util.MaybePanic(err)
	for idx := 0; idx < len(vals); idx += 2 {
		seeds = append(seeds, createRangedValue(vals[idx], vals[idx+1]))
	}
	sort.Sort(RVComparator(seeds))
	return seeds
}

func parseRangedMaps(mapLines []string) (maps [][]ConversionMapRow) {
	inMap := false
	var oneMapRows []ConversionMapRow
	for _, line := range mapLines {
		if inMap {
			vals, err := util.ParseSpacedInts(line)
			if err == nil && len(vals) == cmRowVals {
				oneMapRows = append(oneMapRows, ConversionMapRow{destStart: vals[0], rv: createRangedValue(vals[1], vals[2])})
			} else {
				// complete the map
				sort.Sort(CMRComparator(oneMapRows))
				maps = append(maps, oneMapRows)
				// reset
				inMap = false
				oneMapRows = make([]ConversionMapRow, 0)
			}
		} else {
			inMap = strings.Contains(line, "map:")
		}
	}

	// complete the last map if we haven't already
	if inMap {
		sort.Sort(CMRComparator(oneMapRows))
		maps = append(maps, oneMapRows)
	}

	return
}

func toString(rows []ConversionMapRow) string {
	var b strings.Builder
	b.WriteString("Map:\n")
	for _, row := range rows {
		b.WriteString(row.String() + "\n")
	}
	return b.String()
}

func convert(rVals []RangedValue, mapRows []ConversionMapRow) (out []RangedValue) {
	// plan: find matches by iterating through two ordered lists in tandem via cursors
	rValCursor, mapCursor := 0, 0
	for rValCursor < len(rVals) {
		// if we have already completed the mapCursor, append the remaining valueRanges and exit
		if mapCursor >= len(mapRows) {
			out = append(out, rVals[rValCursor:]...)
			break
		}

		currentVal, currentMap := rVals[rValCursor], mapRows[mapCursor]
		valStart, valEnd, mapStart, mapEnd := currentVal.start, currentVal.end, currentMap.rv.start, currentMap.rv.end

		// if all of the value range is before the map range, pass it all through
		if valEnd < mapStart {
			out = append(out, currentVal)
			rValCursor++
		} else if valStart > mapEnd {
			// if all the map range is before the value range, skip to next map
			mapCursor++
		} else {
			// there is some overlap - split into non-overlap and overlap

			// is any of the rangedValue before the start of the map? pass it through
			if valStart < mapStart {
				beforeSpan := mapStart - valStart
				out = append(out, createRangedValue(valStart, beforeSpan))
			}

			// convert the overlap

			overlapStart := max(valStart, mapStart)
			overlapEnd := min(valEnd, mapEnd)
			overlapSpan := overlapEnd - overlapStart + 1

			mapDelta := currentMap.destStart - mapStart
			out = append(out, createRangedValue(overlapStart+mapDelta, overlapSpan))

			// is any of the rangedValue after the end of the map?
			if valEnd > mapEnd {
				// create a new rangedValue for the portion after this map and set as the current
				afterSpan := valEnd - mapEnd
				rVals[rValCursor] = createRangedValue(mapEnd+1, afterSpan)

				// map is consumed; advance cursor
				mapCursor++
			} else {
				// rangedValue is consumed; advance cursor
				rValCursor++
			}
		}
	}

	// should be sorted already, but let's make sure
	sort.Sort(RVComparator(out))
	return out
}

func day5part2Again(lines []string) int {
	seeds := parseRangedSeeds(lines[0])
	fmt.Println("Seeds", seeds)

	// skip blank line 1
	maps := parseRangedMaps(lines[2:])

	rVals := seeds
	for _, m := range maps {
		fmt.Println(toString(m))
		rVals = convert(rVals, m)
	}

	return rVals[0].start
}

func main() {
	// while I did complete Day 5 Part 2, it wasn't pretty
	// let's see if I can do better

	lines := util.ReadInput("input", "5")

	result := day5part2Again(lines)
	fmt.Println("Part2 Again", result)
}
