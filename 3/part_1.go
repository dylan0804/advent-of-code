package main

import (
	"bufio"
	"os"
	"strings"
)

func processString(text string) int {
	index := strings.Index(text, "mul")

	if index == -1 {
		return 0
	}

	calculatedValue := 0

	if index + 4 < len(text) && text[index+3] == '(' {
		firstNum, firstNumIndex := parseNumber(text, index+4)

		if firstNumIndex < len(text) && text[firstNumIndex] == ',' {
			secondNum, secondNumIndex := parseNumber(text, firstNumIndex+1)

			if secondNumIndex < len(text) && text[secondNumIndex] == ')' {
				calculatedValue = firstNum * secondNum

				index = secondNumIndex
			}
		}
	}

	remainingText := text[index+1:]

	return calculatedValue + processString(remainingText)
}

func parseNumber(text string, i int) (int, int) {
	num := 0

	for i < len(text) && text[i] >= '0' && text[i] <= '9' {
		num = num * 10 + int(text[i] - '0')
		i++
	}
	
	return num, i
}

func part1(file *os.File) int {
	scanner := bufio.NewScanner(file)

	result := 0

	for scanner.Scan() {
		text := scanner.Text()		
		result += processString(text)
	}

	return result
}