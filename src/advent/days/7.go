package main

import (
	"advent/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const cardsInHand = 5

type Hand struct {
	cards       string
	bid         int
	handType    int
	labelValues [cardsInHand]int
}

const fiveOfAKind = 6
const fourOfAKind = 5
const fullHouse = 4
const threeOfAKind = 3
const twoPair = 2
const onePair = 1
const highCard = 0

func handType(cards string) int {
	typeMap := make(map[rune]int)
	for _, char := range cards {
		typeMap[char]++
	}

	highestCount := 0
	var counts []int
	for _, count := range typeMap {
		counts = append(counts, count)
		highestCount = max(count, highestCount)
	}

	switch highestCount {
	case 5:
		return fiveOfAKind
	case 4:
		return fourOfAKind
	case 3:
		if len(counts) == 2 {
			// only possibly is a 3x and a 2x
			return fullHouse
		} else {
			return threeOfAKind
		}
	case 2:
		if len(counts) == 3 {
			// only possibility is 2,2,1
			return twoPair
		} else {
			return onePair
		}
	default:
		return highCard
	}
}

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

func newHand(handStr string, bid int) Hand {
	var labelValues [cardsInHand]int
	for idx, char := range handStr {
		labelValues[idx] = labelValue[char]
	}
	return Hand{cards: handStr, bid: bid, handType: handType(handStr), labelValues: labelValues}
}

func (h Hand) worseHandThan(otherHand Hand) bool {
	if h.handType == otherHand.handType {
		for idx := 0; idx < cardsInHand; idx++ {
			if h.labelValues[idx] != otherHand.labelValues[idx] {
				return h.labelValues[idx] < otherHand.labelValues[idx]
			}
		}
		// they're the same
		return false
	} else {
		return h.handType < otherHand.handType
	}
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

func parseHands(lines []string) (hands []Hand) {
	for _, line := range lines {
		splitLine := strings.Fields(line)
		bid, err := strconv.Atoi(splitLine[1])
		util.MaybePanic(err)
		hands = append(hands, newHand(splitLine[0], bid))
	}
	return
}

func day7part1(lines []string) (acc int) {
	hands := parseHands(lines)
	sort.Sort(CamelGameComparator(hands))

	for idx, hand := range hands {
		acc += hand.bid * (idx + 1)
	}

	return
}

func main() {
	lines := util.ReadInput("input", 7)

	result := day7part1(lines)
	fmt.Println("Part1", result)
}
