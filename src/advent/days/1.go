package main

import (
	"advent/util"
	"fmt"
	"strings"
	"unicode"
)

func converted(digits []int) (out int) {
	if len(digits) == 0 {
		return 0
	}

	last := len(digits) - 1
	return digits[0]*10 + digits[last]
}

func digitsPart1(in string) (out []int) {
	for _, char := range in {
		if unicode.IsDigit(char) {
			out = append(out, int(char)-int('0'))
		}
	}

	return
}

var textToDigit = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	// no zero in text
}

func firstDigitPart2(line string) (digit int) {
	// shouldn't happen but let's fail-safe
	if len(line) == 0 {
		return 0
	}

	char := line[0]
	if unicode.IsDigit(rune(char)) {
		return int(char) - int('0')
	}

	for key := range textToDigit {
		if strings.HasPrefix(line, key) {
			return textToDigit[key]
		}
	}

	lineMinusFirstChar := line[1:]
	return firstDigitPart2(lineMinusFirstChar)
}

func lastDigitPart2(line string) (digit int) {
	// shouldn't happen but let's fail-safe
	if len(line) == 0 {
		return 0
	}

	last := len(line) - 1
	char := line[last]
	if unicode.IsDigit(rune(char)) {
		return int(char) - int('0')
	}

	for key := range textToDigit {
		if strings.HasSuffix(line, key) {
			return textToDigit[key]
		}
	}

	lineMinusLastChar := line[:last]
	return lastDigitPart2(lineMinusLastChar)
}

func digitsPart2(in string) (out []int) {
	return []int{firstDigitPart2(in), lastDigitPart2(in)}
}

func main() {
	input1 := "../../inputs/input1.txt"
	lines := util.ReadLines(input1)

	// debug: confirm that I can read a file
	//for _, value := range lines {
	//	fmt.Println(value)
	//}

	// Part 1 puzzle: the first and last digits in a string form a 2-digit number
	// add these 2-digit numbers

	// Q: more than 2 digits?  A: throw away the middle (clear from example)
	// Q: 1 digit?  A: repeat it (clear from example)
	// Q: 0 digits? A: unclear! I'll assume that I can ignore this / set to 0

	// plan per line:
	// parse digits from string
	// keep first and last
	// convert to 2-digit num

	// final plan: sum lines

	// Part 2: actually text like "five" needs to be considered too, like it's "5"

	var acc = 0
	for _, line := range lines {
		// switch this to change between part 1 and 2
		//digits := digitsPart1(line)
		digits := digitsPart2(line)

		sum := converted(digits)
		acc += sum

		// debug: does this make sense?
		fmt.Println(line, digits, sum)
	}

	fmt.Println("final sum", acc)
}
