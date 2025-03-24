package main

import (
	"strconv"
	"strings"
)

func processSequence(sequence map[string]int, rules map[string][]string) bool {
	for key, val := range sequence {
		setOfRules := rules[key]

		for _, num := range setOfRules {
			currentVal, ok := sequence[num]
			if currentVal < val && ok {
				return false
			}
		}
	}

	return true
}

func part1(sections []string) int {
	var count int

	rulesLine := strings.Split(sections[0], "\n")

	rules := make(map[string][]string)
	// compute rules
	for _, line := range rulesLine {
		splitLine := strings.Split(line, "|")
		rules[splitLine[0]] = append(rules[splitLine[0]], splitLine[1])
	}

	sequenceLine := strings.Split(sections[1], "\n")

	// read each sequence
	for _, line := range sequenceLine {
		actualLine := strings.Split(line, ",")

		sequence := make(map[string]int)
		
		for j, actLine := range actualLine {
			sequence[actLine] = j
		}

		ok := processSequence(sequence, rules)
		if ok {
			indexLen := len(actualLine)

			val, err := strconv.Atoi(actualLine[indexLen/2])
			check(err)

			count += val
		}
	}

	return count
}