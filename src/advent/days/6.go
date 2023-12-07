package main

import (
	"advent/util"
	"fmt"
	"strconv"
	"strings"
)

func countWinners(time, record int) (winners int) {
	middle := time / 2

	if time%2 == 0 {
		// even has a midpoint to check
		if middle*middle > record {
			winners = 1
			// shift by one to set up for the rest of the algorithm
			middle--
		} else {
			// if even the midpoint didn't win, there are no winners
			return 0
		}
	}

	// now check from the midpoint down
	for buttonTime := middle; buttonTime > 0; buttonTime-- {
		raceTime := time - buttonTime
		if buttonTime*raceTime > record {
			// count this victory and its opposite
			winners += 2
		}
	}

	return winners
}

func countAllWinners(times, records []int) int {
	winnerAcc := 1
	for idx, time := range times {
		record := records[idx]
		winnerCount := countWinners(time, record)
		fmt.Println("Race", idx, "has", winnerCount, "winners")
		winnerAcc *= winnerCount
	}

	return winnerAcc
}

func day6part1(lines []string) int {
	times, err := util.ParseSpacedInts(strings.Split(lines[0], ":")[1])
	util.MaybePanic(err)
	fmt.Println("Times", times)

	records, err := util.ParseSpacedInts(strings.Split(lines[1], ":")[1])
	util.MaybePanic(err)
	fmt.Println("Records", records)

	return countAllWinners(times, records)
}

func day6part2(lines []string) int {

	time, err := strconv.Atoi(strings.Replace(strings.Split(lines[0], ":")[1], " ", "", -1))
	util.MaybePanic(err)
	fmt.Println("Time", time)

	record, err := strconv.Atoi(strings.Replace(strings.Split(lines[1], ":")[1], " ", "", -1))
	util.MaybePanic(err)
	fmt.Println("Record", record)

	return countWinners(time, record)
}

func main() {
	inputFile := "../../inputs/input6.txt"
	lines := util.ReadLines(inputFile)

	// boats, see.  Holding their go-buttons for N ms mean they travel at N ms/mm for the next (T-N) ms
	// so you want to hold the button for a while but not the whole while, see

	// we want to beat the previous record distance.  How many values of N would cause us to beat the given record?

	// According to maths and examples:
	// the best N is T/2 and it decreases symmetrically from there. Odd T means even count(N), & v.v.
	// it reminds me of combinatorics/factorial/etc because Dist(N=1) = 1 x (T-1), Dist(N=2) = 2 x (T-2), etc

	result := day6part1(lines)
	fmt.Println("Part1", result)

	// lol the kerning was off so actually we want to smash all the numbers together

	result = day6part2(lines)
	fmt.Println("Part2", result)
}
