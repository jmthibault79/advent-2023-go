package main

import (
	"advent/util"
	"fmt"
	"strings"
)

// can't use uint8 because we overflow that before taking the modulus
func hash(in string) (h uint16) {
	for _, char := range in {
		h = ((h + uint16(char)) * 17) % 256
	}
	return
}

func day15part1(row string) (total uint64) {
	for _, split := range strings.Split(row, ",") {
		total += uint64(hash(split))
	}
	return
}

func main() {
	rows := util.ReadInput("input", "15")
	result := day15part1(rows[0])
	fmt.Println("Part1", result)
}
