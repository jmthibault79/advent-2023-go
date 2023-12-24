package main

import (
	"advent/util"
	"fmt"
	"math"
	"strings"
)

type Galaxy struct {
	row int
	col int
}

const empty = '.'
const galaxy = '#'

func expand(inMap []string) (outMap []string) {
	const yes = 1

	// a new map with duplicated rows
	var withDupedRows []string
	// record those columns which are not empty
	hasGalaxyCols := make(map[int]int)

	for _, row := range inMap {
		allEmpty := true
		for cIdx, char := range row {
			if char == galaxy {
				allEmpty = false
				hasGalaxyCols[cIdx] = yes
			}
		}
		if allEmpty {
			// double it
			withDupedRows = append(withDupedRows, row, row)
		} else {
			withDupedRows = append(withDupedRows, row)
		}
	}

	for _, row := range withDupedRows {
		var outRow strings.Builder
		for cIdx, char := range row {
			if hasGalaxyCols[cIdx] == yes {
				outRow.WriteString(string(char))
			} else {
				// if empty col, dupe it
				outRow.WriteString(string(char) + string(char))
			}
		}
		outMap = append(outMap, outRow.String())
	}

	return outMap
}

func parseGalaxies(inMap []string) (galaxies []Galaxy) {
	for rIdx, row := range inMap {
		for cIdx, char := range row {
			if char == galaxy {
				galaxies = append(galaxies, Galaxy{
					row: rIdx,
					col: cIdx,
				})
			}
		}
	}
	return
}

func distance(gal1, gal2 Galaxy) int {
	rowDist := math.Abs(float64(gal1.row - gal2.row))
	colDist := math.Abs(float64(gal1.col - gal2.col))
	return int(rowDist + colDist)
}

func day11part1(inMap []string) (total int) {
	galaxies := parseGalaxies(inMap)
	for gIdx, gal := range galaxies {
		for g2Idx := gIdx + 1; g2Idx < len(galaxies); g2Idx++ {
			gal2 := galaxies[g2Idx]
			total += distance(gal, gal2)
		}
	}
	return
}

func main() {
	initialUniverse := util.ReadInput("input", "11")
	expandedUniverse := expand(initialUniverse)
	result := day11part1(expandedUniverse)
	fmt.Println("Part1", result)
}
