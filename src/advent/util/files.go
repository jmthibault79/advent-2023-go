package util

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(filePath string) (lines []string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	// defer means to run it after function completes but before it returns
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Iterate through each line in the file
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return
}

func ReadInput(kind string, day int) (lines []string) {
	dir := "../../inputs/"
	filename := fmt.Sprintf("%s%d.txt", kind, day)
	return ReadLines(dir + filename)
}
