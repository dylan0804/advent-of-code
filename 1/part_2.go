package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func part2() int {
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := make(map[string]int, 1000)
	firstValues := make([]int, 1000)

	for scanner.Scan() {
		lines := strings.Fields(scanner.Text())
		num, err := strconv.Atoi(lines[0])
		check(err)
		firstValues = append(firstValues, num)
		count[lines[1]]++
	}

	var result int

	for _, val := range firstValues {
		key := strconv.Itoa(val)
		result += val * count[key]
	}

	return result
}