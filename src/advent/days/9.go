package main

import (
	"advent/util"
	"fmt"
)

// 1 2 3
// 1 1
// 0
func nextVal(seq []int) (out int) {
	diffs := make([][]int, 1)
	diffs[0] = seq
	for i := 0; !util.AllEqual(diffs[i], 0); i++ {
		diffLen := len(diffs[i]) - 1
		diffs = append(diffs, make([]int, diffLen))
		for j := 0; j < diffLen; j++ {
			diffs[i+1][j] = diffs[i][j+1] - diffs[i][j]
		}
	}

	// diffs has N rows.  diffs[N-1] is all-zero, so N-2 is the first useful one
	// we want the sum of the last vals
	for ii := len(diffs) - 2; ii >= 0; ii-- {
		lastPos := len(diffs[0]) - 1 - ii
		out += diffs[ii][lastPos]
	}
	return
}

func day9part1(seqs [][]int) (out int) {
	for _, seq := range seqs {
		next := nextVal(seq)
		out += next
		fmt.Println(seq, " -> ", next)
	}
	return
}

func parseInput9(lines []string) (out [][]int) {
	for _, line := range lines {
		nums, err := util.ParseSpacedInts(line)
		util.MaybePanic(err)
		out = append(out, nums)
	}
	return
}

func main() {
	lines1 := util.ReadInput("input", "9")
	sequences := parseInput9(lines1)
	result := day9part1(sequences)
	fmt.Println("Part1", result)

}
