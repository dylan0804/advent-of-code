package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func isAscending(line[] string) bool {
	for i := 1; i < len(line); i++ {
		second, err := strconv.Atoi(line[i])
		check(err)

		first, err := strconv.Atoi(line[i-1])
		check(err)

		if second < first {
			return false
		}
	}

	return true
}

func isDescending(line[] string) bool {
	for i := 1; i < len(line); i++ {
		second, err := strconv.Atoi(line[i])
		check(err)

		first, err := strconv.Atoi(line[i-1])
		check(err)

		if second > first {
			return false
		}
	}

	return true
}

func isSafe(line []string) bool {
	for i := 1; i < len(line); i++ {
		second, err := strconv.Atoi(line[i])
		check(err)

		first, err := strconv.Atoi(line[i-1])
		check(err)

		diff := second - first
		if diff < 0 {
			diff *= -1
		}

		if diff > 3 || diff < 1 {
			return false
		}
	}

	return true
}

func processLine(line []string) bool {
	if (isAscending(line) || isDescending(line)) && isSafe(line) {
		return true
	}

	
	
	return false
}

func part1() int {
	file, err := os.Open("./input.txt")
	check(err)

	scanner := bufio.NewScanner(file)

	var count int

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())

		ok := processLine(line)

		if ok {
			count++
		}
	}

	return count
}