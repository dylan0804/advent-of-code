package main

import (
	"bufio"
	"os"
	"strings"
)

func processLineTwo(line []string) bool {
	if (isAscending(line) || isDescending(line)) && isSafe(line) {
		return true
	}

	for i := 0; i < len(line); i++ {
		newLine := make([]string, 0, len(line)-1)
		newLine = append(newLine, line[:i]...)

		if i < len(line)-1 {
			newLine = append(newLine, line[i+1:]...)
		}

		if (isAscending(newLine) || isDescending(newLine)) && isSafe(newLine) {
			return true
		}
	}
	
	return false
}

func part2() int {
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var count int

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())

		ok := processLineTwo(line)

		if ok {
			count++
		}
	}

	return count
}