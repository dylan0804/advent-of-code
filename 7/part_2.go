package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)


func appendNumbers(a, b int) int {
	digits := 1

	temp := b

	for temp >= 10 {
		temp /= 10
		digits *= 10
	}

	return a * digits * 10 + b
}

func checkCombinationTwo(target int, nums []int, index, currentVal int) bool {
	if index == len(nums) {
		return currentVal == target
	}

	if checkCombinationTwo(target, nums, index+1, currentVal*nums[index]) {
		return true
	}

	if checkCombinationTwo(target, nums, index+1, currentVal+nums[index]) {
		return true
	}

	if checkCombinationTwo(target, nums, index+1, appendNumbers(currentVal, nums[index])) {
		return true
	}

	return false
}

func processNumsTwo(target int, nums []int) bool {
	if nums[0] >= target {
		return false
	}

	if len(nums) == 1 && nums[0] != target {
		return false
	}

	return checkCombinationTwo(target, nums, 1, nums[0])
}

func part2(file *os.File) int {
	scanner := bufio.NewScanner(file)

	count := 0

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")

		targetString, nums := line[0], strings.Fields(line[1])

		listOfNums := stringsToInts(nums)

		target, err := strconv.Atoi(targetString)
		check(err)

		valid := processNumsTwo(target, listOfNums)
		if valid {
			count += target
		}
	}

	return count
}