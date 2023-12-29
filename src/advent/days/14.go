package main

import (
	"advent/util"
	"fmt"
)

const rollingRock = 'O'
const fixedCube = '#'
const emptySpace = '.'

func tiltNorth(rows []string) (tilted [][]rune) {
	// initially set the top row equal to the input
	tilted = make([][]rune, len(rows))
	tilted[0] = ([]rune)(rows[0])

	rowIdx := 1
	for _, row := range rows[1:] {
		tilted[rowIdx] = ([]rune)(row)
		for colIdx, char := range tilted[rowIdx] {
			// if the current char is a rollingRock
			if char == rollingRock {
				// roll up as far as it can go
				for rollupRowIdx := rowIdx - 1; rollupRowIdx >= 0 && tilted[rollupRowIdx][colIdx] == emptySpace; rollupRowIdx-- {
					// roll this rock up one space
					tilted[rollupRowIdx][colIdx] = rollingRock
					tilted[rollupRowIdx+1][colIdx] = emptySpace
				}
			}
		}
		rowIdx++
	}

	return
}

func totalLoad(tilted [][]rune) (total int) {
	loadForThisRow := len(tilted)
	for _, row := range tilted {
		for _, char := range row {
			if char == rollingRock {
				total += loadForThisRow
			}
		}
		loadForThisRow--
	}
	return
}

func day14part1(rows []string) int {
	northTilted := tiltNorth(rows)
	return totalLoad(northTilted)
}

func main() {
	rows := util.ReadInput("input", "14")
	result := day14part1(rows)
	fmt.Println("Part1", result)
}
