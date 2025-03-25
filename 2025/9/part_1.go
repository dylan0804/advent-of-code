package main

var data []int

func part1(content string) int {
	// 12345
	id := 0

	for i, char := range content {
		digit := char - '0'
		processChar(int(digit), id, i)

		if i % 2 == 0 {
			id++
		}
	}

	moveBlocks(data)
	
	checkSum := calculateChecksum(data)

	return checkSum
}

func calculateChecksum(nums []int) int {
	count := 0

	for i, num := range nums {
		if num != -1 {
			count += i * num
		}
	}

	return count
}

func moveBlocks(nums []int) {
	left := 0
	right := len(nums) - 1

	for {
		// find the next empty space
		for left < len(nums) && nums[left] != -1 {
			left++
		}

		// find the next available space
		for right >= 0 && nums[right] == -1 {
			right--
		}

		if left >= right {
			break
		}

		nums[left], nums[right] = nums[right], nums[left]

		left++
		right--
	}
}

func processChar(digit int, id, flag int) {
	i := 0
	
	for i < digit {
		if flag % 2 == 0 {
			data = append(data, id)
		} else {
			data = append(data, -1)
		}
		i++
	}
}