package main

import (
	"advent/util"
	"fmt"
)

// 1 2 3
// 1 1
// 0
func getDiffs(seq []int) (diffs [][]int) {
	diffs = make([][]int, 1)
	diffs[0] = seq
	for i := 0; !util.AllEqual(diffs[i], 0); i++ {
		diffLen := len(diffs[i]) - 1
		diffs = append(diffs, make([]int, diffLen))
		for j := 0; j < diffLen; j++ {
			diffs[i+1][j] = diffs[i][j+1] - diffs[i][j]
		}
	}
	return
}

func nextVal(seq []int) (out int) {
	diffs := getDiffs(seq)

	// diffs has N rows.  diffs[N-1] is all-zero, so N-2 is the first useful one
	// we want the sum of the last vals
	for ii := len(diffs) - 2; ii >= 0; ii-- {
		lastPos := len(diffs[0]) - 1 - ii
		out += diffs[ii][lastPos]
	}
	return
}

func prevVal(seq []int) (out int) {
	diffs := getDiffs(seq)

	// diffs has N rows.  diffs[N-1] is all-zero
	// do this: curVal = 0.  for each diff [N-2 ... 0] curVal = diff[0] - curVal
	for ii := len(diffs) - 2; ii >= 0; ii-- {
		out = diffs[ii][0] - out
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

func day9part2(seqs [][]int) (out int) {
	for _, seq := range seqs {
		next := prevVal(seq)
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
	lines := util.ReadInput("input", "9")
	sequences := parseInput9(lines)
	result := day9part1(sequences)
	fmt.Println("Part1", result)

	result = day9part2(sequences)
	fmt.Println("Part2", result)
}
