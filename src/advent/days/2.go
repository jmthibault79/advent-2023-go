package main

import (
	"advent/util"
	"fmt"
	"strings"
)

type Draw struct {
	redCount   int
	greenCount int
	blueCount  int
}

func (d Draw) String() string {
	return fmt.Sprintf("[%d Red, %d Green, %d Blue]", d.redCount, d.greenCount, d.blueCount)
}

// expected input looks like 8 green, 6 blue, 20 red but not all may be present
func parseDraw(drawString string) Draw {
	colorStrings := strings.Split(drawString, ",")

	redC, greenC, blueC := 0, 0, 0

	var tempCount int
	var tempColor string
	for _, cStr := range colorStrings {
		_, err := fmt.Sscanf(cStr, "%d %s", &tempCount, &tempColor)
		if err != nil {
			fmt.Println("Error:", err)
			return Draw{redCount: 0, greenCount: 0, blueCount: 0}
		}

		// not sure if I'm grabbing any extra spaces,
		// so let's use Contains to be safe
		if strings.Contains(tempColor, "red") {
			redC = tempCount
		} else if strings.Contains(tempColor, "green") {
			greenC = tempCount
		} else if strings.Contains(tempColor, "blue") {
			blueC = tempCount
		}
	}

	return Draw{redCount: redC, greenCount: greenC, blueCount: blueC}
}

type Game struct {
	id      int
	draws   []Draw
	atLeast Draw
}

func (g Game) String() string {
	var drawStrings []string
	for _, d := range g.draws {
		drawStrings = append(drawStrings, d.String())
	}

	return fmt.Sprintf("Game %d: [ %s ] -> %s",
		g.id, strings.Join(drawStrings, ", "), g.atLeast)
}

func parseGame(line string) Game {
	// should split into "Game x" and the draws
	gameIdVsDraws := strings.Split(line, ":")

	var gameId int
	_, err := fmt.Sscanf(gameIdVsDraws[0], "Game %d", &gameId)
	if err != nil {
		fmt.Println("Error:", err)
		return Game{id: 0, draws: []Draw{}, atLeast: Draw{}}
	}

	// should split into the individual draws
	drawStrings := strings.Split(gameIdVsDraws[1], ";")

	var draws []Draw
	atLeastR, atLeastG, atLeastB := 0, 0, 0
	for _, dStr := range drawStrings {
		d := parseDraw(dStr)
		atLeastR = max(atLeastR, d.redCount)
		atLeastG = max(atLeastG, d.greenCount)
		atLeastB = max(atLeastB, d.blueCount)
		draws = append(draws, d)
	}

	return Game{id: gameId, draws: draws,
		atLeast: Draw{redCount: atLeastR, greenCount: atLeastG, blueCount: atLeastB}}
}

func part1(lines []string, redCount, greenCount, blueCount int) int {
	acc := 0

	for _, line := range lines {
		game := parseGame(line)
		atLeast := game.atLeast
		if atLeast.redCount <= redCount &&
			atLeast.greenCount <= greenCount &&
			atLeast.blueCount <= blueCount {
			//fmt.Printf("Game %d Possible\n", game.id)
			acc += game.id
		}
	}

	return acc
}

func main() {
	inputFile := "../../inputs/input2.txt"
	lines := util.ReadLines(inputFile)

	// Part 1
	// each GAME (a line) has "a few" DRAWS which each reveal X red, Y blue, Z green cubes
	// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	// -> G1 = [ [3B,4R], [1R,2G,6B], [2G] ]
	// then Elfis asks:
	// which of these games were possible if the total was A red, B green, C blue?
	// give me the sum of those game IDs
	// Note: the cubes are replaced between draws

	// test my "Stringers"
	//testGame := Game{id: 10, draws: []Draw{
	//	{redCount: 1, greenCount: 2, blueCount: 3},
	//	{redCount: 0, greenCount: 0, blueCount: 3},
	//}}
	//fmt.Println(testGame)
	//for _, l := range lines {
	//	fmt.Println(parseGame(l))
	//}

	// for Part 1, what is the sum of the games that are possible w/ 12 red, 13 green, 14 blue?
	result := part1(lines, 12, 13, 14)
	fmt.Println(result)
}
