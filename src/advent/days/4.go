package main

import (
	"advent/util"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func parseInts(s string) (ints []int) {
	var num int
	for _, token := range strings.Split(s, " ") {
		scanned, err := fmt.Sscanf(token, "%d", &num)
		// skip if blank or other error
		if err == nil && scanned == 1 {
			ints = append(ints, num)
		}
	}
	return
}

type Card struct {
	id    int
	lucky []int
	mine  []int
}

func (c Card) String() string {
	var luckyStr strings.Builder
	for _, n := range c.lucky {
		luckyStr.WriteString(strconv.Itoa(n) + " ")
	}
	var mineStr strings.Builder
	for _, n := range c.mine {
		mineStr.WriteString(strconv.Itoa(n) + " ")
	}

	return fmt.Sprintf("C%d: Lucky %s | Mine %s", c.id, luckyStr.String(), mineStr.String())
}

func (c Card) matches() (matches int) {
	// make a map of the lucky nums for easy lookup
	// nothing was said about duplicates and I don't think there are any, so let's ignore 'em
	m := make(map[int]int)
	for idx, num := range c.lucky {
		m[num] = idx + 1 // need to set to something nonzero
	}

	for _, num := range c.mine {
		if m[num] > 0 {
			matches++
		}
	}

	return
}

func parseCard(line string) Card {
	// should split into "Card x" and the numbers
	cardIdVsNumbers := strings.Split(line, ":")

	var cardId int
	_, err := fmt.Sscanf(cardIdVsNumbers[0], "Card %d", &cardId)
	util.MaybePanic(err)

	luckyVsMine := strings.Split(cardIdVsNumbers[1], "|")
	lucky := parseInts(luckyVsMine[0])
	mine := parseInts(luckyVsMine[1])

	return Card{id: cardId, lucky: lucky, mine: mine}
}

func score(matches int) int {
	if matches == 0 {
		return 0
	}
	return int(math.Exp2(float64(matches - 1)))
}

func day4part1(lines []string) int {
	acc := 0

	for _, line := range lines {
		card := parseCard(line)
		fmt.Println(card)
		acc += score(card.matches())
	}

	return acc
}

type CardAndCount struct {
	card  Card
	count int
}

func day4part2(lines []string) int {
	// keeps track of the hand of cards and also the final score
	cardsAndCounts := make(map[int]CardAndCount)

	// parse and count the originals
	for idx, line := range lines {
		card := parseCard(line)
		cardsAndCounts[idx] = CardAndCount{card: card, count: 1}
	}

	// score and multiply
	acc := 0
	for idx := 0; idx < len(lines); idx++ {
		cc := cardsAndCounts[idx]
		acc += cc.count
		matches := cc.card.matches()
		// add `count` copies of each of the cards corresponding to `matches`
		for copyIdx := idx + 1; copyIdx < len(lines) && copyIdx <= idx+matches; copyIdx++ {
			copied := cardsAndCounts[copyIdx]
			cardsAndCounts[copyIdx] = CardAndCount{
				card:  copied.card,
				count: copied.count + cc.count,
			}
		}
	}

	return acc
}

func main() {
	inputFile := "../../inputs/input4.txt"
	lines := util.ReadLines(inputFile)

	// Part 1
	// Scratch cards have lucky numbers on them, and I get points if I match them with my numbers
	// 0 matches = 0 pts, 1 match = 1 pt, and each beyond that is x2 (2 -> 2, 3 -> 4, 4 -> 8)

	result := day4part1(lines)
	fmt.Println("Part1", result)

	// Part 2 keeps the same idea of matches, but what you win is more cards
	// if card N has 3 matches, you win an extra copy of N+1, N+2, N+3
	// so: when it comes time to score N+1, you do it twice, and so on.
	// How many CARDS do you have, originals and copies?

	result = day4part2(lines)
	fmt.Println("Part2", result)
}
