package main

import (
	"advent/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const cardsInHand = 5
const jokerLabel = 'J'
const jokerValue = 1

const fiveOfAKind = 6
const fourOfAKind = 5
const fullHouse = 4
const threeOfAKind = 3
const twoPair = 2
const onePair = 1
const highCard = 0

var labelValue = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

type Hand struct {
	cards       string
	bid         int
	handType    int
	labelValues [cardsInHand]int
}

func handType(cards string, part2 bool) int {
	typeMap := make(map[rune]int)
	for _, label := range cards {
		typeMap[label]++
	}

	highestCount := 0
	var jokerCount int
	var counts []int
	for label, count := range typeMap {
		if part2 && label == jokerLabel {
			jokerCount = count
		} else {
			counts = append(counts, count)
			highestCount = max(count, highestCount)
		}
	}

	switch highestCount + jokerCount {
	case 5:
		return fiveOfAKind
	case 4:
		return fourOfAKind
	case 3:
		if jokerCount == 0 || jokerCount == 1 {
			// if no jokers, the counts are 3,2 or 3,1,1
			// if 1 joker, the counts are 2,2 or 2,1,1
			if len(counts) == 2 {
				return fullHouse
			} else {
				return threeOfAKind
			}
		} else {
			// if 2 jokers, the counts are 1,1,1
			return threeOfAKind
		}
	case 2:
		// if no jokers, the counts are 2,2,1 or 2,1,1,1
		// if 1 joker, the counts are 1,1,1,1
		if len(counts) == 3 {
			return twoPair
		} else {
			return onePair
		}
	default:
		return highCard
	}
}

func newHand(handStr string, bid int, part2 bool) Hand {
	var labelValues [cardsInHand]int
	for idx, char := range handStr {
		if part2 && char == jokerLabel {
			labelValues[idx] = jokerValue
		} else {
			labelValues[idx] = labelValue[char]
		}
	}
	return Hand{cards: handStr, bid: bid, handType: handType(handStr, part2), labelValues: labelValues}
}

type CamelGameComparator []Hand

// these are directly from ChatGPT but they're trivial so I don't care
func (c CamelGameComparator) Len() int      { return len(c) }
func (c CamelGameComparator) Swap(i, j int) { c[i], c[j] = c[j], c[i] }

func (c CamelGameComparator) Less(i, j int) bool {
	if c[i].handType != c[j].handType {
		return c[i].handType < c[j].handType
	}

	for idx := 0; idx < cardsInHand; idx++ {
		if c[i].labelValues[idx] != c[j].labelValues[idx] {
			return c[i].labelValues[idx] < c[j].labelValues[idx]
		}
	}
	// they're the same
	return false
}

func parseHands(lines []string, part2 bool) (hands []Hand) {
	for _, line := range lines {
		splitLine := strings.Fields(line)
		bid, err := strconv.Atoi(splitLine[1])
		util.MaybePanic(err)
		hands = append(hands, newHand(splitLine[0], bid, part2))
	}
	return
}

func day7(lines []string, part2 bool) (acc int) {
	hands := parseHands(lines, part2)
	sort.Sort(CamelGameComparator(hands))

	for idx, hand := range hands {
		acc += hand.bid * (idx + 1)
	}

	return
}

func main() {
	lines := util.ReadInput("input", "7")

	result := day7(lines, false)
	fmt.Println("Part1", result)

	result = day7(lines, true)
	fmt.Println("Part2", result)
}
