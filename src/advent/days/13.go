package main

import (
	"advent/util"
	"fmt"
)

func findVertical(pattern []string) (colsLeft int, found bool) {
	lastCol := len(pattern[0]) - 1
	candidates := make(map[int]bool, lastCol)

	// start by finding where adjacent columns match in the first row
	for idx := 0; idx < lastCol; idx++ {
		if pattern[0][idx] == pattern[0][idx+1] {
			// only record matches
			candidates[idx] = true
		}
	}

	// narrow down these possibilities for each successive row
	for _, row := range pattern[1:] {
		for idx := range candidates {
			if row[idx] != row[idx+1] {
				delete(candidates, idx)
			}
		}
	}

	// now we need to *confirm* the candidates
	for idx := range candidates {
		found = true
		// start at the candidate (left) and its adjacent (right); step both outward by one
	nextCandidate:
		for leftCol, rightCol := idx, idx+1; leftCol >= 0 && rightCol <= lastCol; leftCol, rightCol = leftCol-1, rightCol+1 {
			for _, row := range pattern {
				if row[leftCol] != row[rightCol] {
					found = false
					break nextCandidate
				}
			}
		}
		if found {
			return idx + 1, true
		}
	}

	// none matched
	return -1, false
}

func findHorizontal(pattern []string) (rowsAbove int, found bool) {
	lastRow := len(pattern) - 1
	candidates := make(map[int]bool, lastRow)

	// start by finding where adjacent rows match in the first column
	for idx := 0; idx < lastRow; idx++ {
		if pattern[idx][0] == pattern[idx+1][0] {
			// only record matches
			candidates[idx] = true
		}
	}

	// narrow down these possibilities for each successive column
	for colIdx := range pattern[0] {
		for idx := range candidates {
			if pattern[idx][colIdx] != pattern[idx+1][colIdx] {
				delete(candidates, idx)
			}
		}
	}

	// now we need to *confirm* the candidates
	for idx := range candidates {
		found = true
		// start at the candidate (top) and its adjacent (bottom); step both outward by one
	nextCandidate:
		for topRow, bottomRow := idx, idx+1; topRow >= 0 && bottomRow <= lastRow; topRow, bottomRow = topRow-1, bottomRow+1 {
			for colIdx := range pattern[0] {
				if pattern[topRow][colIdx] != pattern[bottomRow][colIdx] {
					found = false
					break nextCandidate
				}
			}
		}
		if found {
			return idx + 1, true
		}
	}

	// none matched
	return -1, false
}

func summary(pattern []string) (out int) {
	columnsLeft, found := findVertical(pattern)
	if found {
		out += columnsLeft
	}
	rowsAbove, found := findHorizontal(pattern)
	if found {
		out += rowsAbove * 100
	}
	return
}

func day13part1(rows []string) (total int) {
	var onePattern []string
	for _, row := range rows {
		if row == "" {
			total += summary(onePattern)
			onePattern = nil
		} else {
			onePattern = append(onePattern, row)
		}
	}

	if len(onePattern) > 0 {
		total += summary(onePattern)
	}
	return
}

func main() {
	rows := util.ReadInput("input", "13")
	result := day13part1(rows)
	fmt.Println("Part1", result)
}
