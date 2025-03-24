package main

import (
	"sort"
	"strconv"
	"strings"
)

func sortSequence(sequence []string, rules map[string][]string) []string {
    dependencies := make(map[string]map[string]bool)

    for _, num := range sequence {
        dependencies[num] = make(map[string]bool)
    }

    for before, afters := range rules {
        for _, after := range afters {
            if _, exists := dependencies[after]; exists {
                dependencies[after][before] = true
            }
        }
    }

    sort.SliceStable(sequence, func(i, j int) bool {
        a := sequence[i]
        b := sequence[j]

        if dependencies[b][a] {
            return true
        }

        if dependencies[a][b] {
            return false
        }

        return i < j
    })

    return sequence
}

func part2(sections []string) int {
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
        if !ok {
            correctSequence := sortSequence(actualLine, rules)

            mid, err := strconv.Atoi(correctSequence[len(correctSequence)/2])
            check(err)

            count += mid
        }
	}

	return count
}