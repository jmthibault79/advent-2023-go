package main

import (
	"advent/util"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Symbol struct {
	value string
	row   int
	col   int
}

func (s Symbol) String() string {
	return fmt.Sprintf("%s -> [%d, %d]", s.value, s.row, s.col)
}

type Part struct {
	value int
	row   int
	start int
	end   int
}

func (p Part) String() string {
	return fmt.Sprintf("%d -> [%d, %d-%d]", p.value, p.row, p.start, p.end)
}

func completePart(partString string, rowIdx, partStart, partEnd int) Part {
	partValue, err := strconv.Atoi(partString)
	if err != nil {
		fmt.Println("Error:", err)
		return Part{}
	}

	return Part{value: partValue, row: rowIdx, start: partStart, end: partEnd}
}

const notInPart = -1

func parseLine(rowIdx int, line string) (s []Symbol, p []Part) {
	// part accumulator for the current part, if any
	var partAcc strings.Builder
	// part start for the current part, if any
	var partStart = notInPart

	for colIdx, char := range line {
		// case: we're inside a part number
		if unicode.IsDigit(char) {
			partAcc.WriteRune(char)
			// case: we're at the *beginning* of a part
			if partStart == notInPart {
				partStart = colIdx
			}
		} else {
			// case: we just finished parsing a part, so now we need to ship it
			if partStart != notInPart {
				part := completePart(partAcc.String(), rowIdx, partStart, colIdx-1)
				p = append(p, part)

				// reset for next part
				partAcc.Reset()
				partStart = notInPart
			}

			// case: a symbol
			if char != '.' {
				s = append(s, Symbol{value: string(char), row: rowIdx, col: colIdx})
			}
		}
	}

	// account for a part at the end of the line
	if partStart != notInPart {
		part := completePart(partAcc.String(), rowIdx, partStart, len(line)-1)
		p = append(p, part)
	}

	return
}

func symbolAdjacent(part Part, symbolMap map[int]map[int]Symbol, lastRow, lastCol int) (adj bool, s Symbol) {
	// define the search space: a box surrounding the part, clamped to the field edges
	rowStart := max(part.row-1, 0)
	rowEnd := min(part.row+1, lastRow)
	colStart := max(part.start-1, 0)
	colEnd := min(part.end+1, lastCol)

	// man, for loops take me back
	for row := rowStart; row <= rowEnd; row++ {
		symbolMapRow := symbolMap[row]
		if symbolMapRow != nil {
			for col := colStart; col <= colEnd; col++ {
				s := symbolMapRow[col]
				if len(s.value) > 0 {
					return true, s
				}
			}
		}
	}

	return false, Symbol{}
}

// make a 2D map of [row, col] -> symbol for all symbols
func makeSymbolMap(symbols []Symbol) (out map[int]map[int]Symbol) {
	out = make(map[int]map[int]Symbol)
	for _, s := range symbols {
		mapRow := out[s.row]
		if mapRow == nil {
			out[s.row] = make(map[int]Symbol)
		}
		out[s.row][s.col] = s
	}
	return
}

func day3part1(symbols []Symbol, parts []Part, lastRow, lastCol int) int {
	symbolMap := makeSymbolMap(symbols)
	//fmt.Println(symbolMap)

	adjacentPartsSum := 0

	for _, p := range parts {
		adj, s := symbolAdjacent(p, symbolMap, lastRow, lastCol)
		if adj {
			adjacentPartsSum += p.value
			fmt.Println("Adjacent:", p, s)
		}
	}

	return adjacentPartsSum
}

// make a map of row -> [parts in row] for all parts
func makePartMap(parts []Part) (out map[int][]Part) {
	out = make(map[int][]Part)
	for _, p := range parts {
		out[p.row] = append(out[p.row], p)
	}
	return
}

func gearRatio(symbol Symbol, partMap map[int][]Part) int {
	// define the search space: from row before to row after
	// don't need to clamp because there are no gears in the first or last row or column
	rowStart := symbol.row - 1
	rowEnd := symbol.row + 1
	colStart := symbol.col - 1
	colEnd := symbol.col + 1

	var adjParts []Part
	for row := rowStart; row <= rowEnd; row++ {
		parts := partMap[row]
		for _, part := range parts {
			// overlap test: if the two ranges are disjoint, they do not overlap
			// can do this comparison in either direction
			disjoint := (part.end < colStart) || (part.start > colEnd)
			if !disjoint {
				adjParts = append(adjParts, part)
			}
		}
	}

	// exactly 2!
	if len(adjParts) == 2 {
		ratio := adjParts[0].value * adjParts[1].value
		fmt.Println(symbol, "is a gear with ratio", adjParts[0].value, "x", adjParts[1].value, ratio)
		return ratio
	} else {
		return 0
	}
}

const gearSymbol = "*"

func day3part2(symbols []Symbol, parts []Part) (sum int) {
	partMap := makePartMap(parts)
	//fmt.Println(partMap)

	for _, s := range symbols {
		if s.value == gearSymbol {
			sum += gearRatio(s, partMap)
		}
	}

	return sum
}

func main() {
	inputFile := "../../inputs/input3.txt"
	lines := util.ReadLines(inputFile)

	// Day 3 Part 1
	// a 2D grid represents an engine schematic, filled with numbers (multi-digit!) and symbols
	// the symbol . means empty
	// collect any number adjacent to a symbol (including diagonal) and add them up

	// approach:
	// collect part numbers and symbols while scanning line by line
	// record begin and end coords for part numbers
	// record coords for symbols
	// adjacency algorithm TBD

	var symbols []Symbol
	var parts []Part

	// yay finally using the index I've been throwing away so far
	for idx, line := range lines {
		lSymbols, lParts := parseLine(idx, line)
		symbols = append(symbols, lSymbols...)
		parts = append(parts, lParts...)
	}

	fmt.Println("Symbols: ")
	for _, s := range symbols {
		fmt.Println(s)
	}

	fmt.Println("Parts: ")
	for _, p := range parts {
		fmt.Println(p)
	}

	// for Part 1, which parts are adjacent to symbols?  add them.

	// approach: make a 2D map of symbols. for each part, search its bounding box for a symbol

	result := day3part1(symbols, parts, len(lines)-1, len(lines[0])-1)
	fmt.Println("Part1", result)

	// Part 2: we only care about GEARS which are * symbols which are adjacent to exactly 2 parts
	// a GEAR RATIO is the multiplication of those 2 parts.  Add these up.

	// stats: ~400 * in ~140 lines - but not the first and last lines, so I don't need to account for these

	// approach: make a map of line -> parts on that line.  for each * symbol, check -1 to +1 for 2 parts

	result = day3part2(symbols, parts)
	fmt.Println("Part2", result)
}
