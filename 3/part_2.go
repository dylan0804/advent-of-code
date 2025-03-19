package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func processStringTwo(text string, mulEnabled bool) int {
	calculatedValue := 0

	index := strings.Index(text, "mul")
	dontIndex := strings.Index(text, "don't()")
	doIndex := strings.Index(text, "do()")

	if index == -1 && doIndex == -1 && dontIndex == -1 {
		return 0
	}

	earliestPos := len(text)
	earliestType := ""

	if doIndex != -1 && doIndex < earliestPos {
		earliestPos = doIndex
		earliestType = "do"
	}

	if dontIndex != -1 && dontIndex < earliestPos {
		earliestPos = dontIndex
		earliestType = "dont"
	}

	if index != -1 && index < earliestPos {
		earliestPos = index
		earliestType = "mul"
	}

	fmt.Println(earliestType)

	if earliestType == "do" {
		return processStringTwo(text[earliestPos+3:], true)
	} else if earliestType == "dont" {
		return processStringTwo(text[earliestPos+6:], false)
	} else if earliestType == "mul" {
		if mulEnabled && earliestPos + 3 < len(text) && text[earliestPos+3] == '(' {
			firstNum, firstNumIndex := parseNumber(text, earliestPos+4)
	
			if firstNumIndex < len(text) && text[firstNumIndex] == ',' {
				secondNum, secondNumIndex := parseNumber(text, firstNumIndex+1)
	
				if secondNumIndex < len(text) && text[secondNumIndex] == ')' {
					calculatedValue = firstNum * secondNum

					fmt.Println(firstNum, secondNum)
	
					return calculatedValue + processStringTwo(text[secondNumIndex+1:], mulEnabled)
				}
			}
		}

		return processStringTwo(text[earliestPos+3:], mulEnabled)
	}

	return 0
}

func part2(file *os.File) int {
	scanner := bufio.NewScanner(file)

	var allTexts strings.Builder

	for scanner.Scan() {
		text := scanner.Text()		
		allTexts.WriteString(text)
	}

	return processStringTwo(allTexts.String(), true)

}