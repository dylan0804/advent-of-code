package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func checkCombination(target int, nums []int, index, currentVal int) bool {
	if index == len(nums) {
		return currentVal == target
	}

	if checkCombination(target, nums, index+1, currentVal*nums[index]) {
		return true
	}

	if checkCombination(target, nums, index+1, currentVal+nums[index]) {
		return true
	}

	return false
}

func processNums(target int, nums []int) bool {
	if nums[0] >= target {
		return false
	}

	if len(nums) == 1 && nums[0] != target {
		return false
	}

	return checkCombination(target, nums, 1, nums[0])
}

func part1(file *os.File) int {
	scanner := bufio.NewScanner(file)

	count := 0

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")

		targetString, nums := line[0], strings.Fields(line[1])

		listOfNums := stringsToInts(nums)

		target, err := strconv.Atoi(targetString)
		check(err)

		valid := processNums(target, listOfNums)
		if valid {
			count += target
		}
	}

	return count
}