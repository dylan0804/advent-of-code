package main

import (
	"bufio"
	"container/heap"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type intHeap []int

func (h intHeap) Len() int { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)  { h[i], h[j] = h[j], h[i]}

func (h *intHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() any {
	oldHeap := *h
	oldLength := len(oldHeap)
	poppedValue := oldHeap[oldLength-1]
	*h = oldHeap[0 : oldLength-1]
	return poppedValue
}

func part1() int {
	hl, hr := &intHeap{}, &intHeap{}
	heap.Init(hl)
	heap.Init(hr)

	f, err := os.Open("./input.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		nums := strings.Fields(line)

		num, _ := strconv.Atoi(nums[0])
		heap.Push(hl, num)

		num, _ = strconv.Atoi(nums[1])
		heap.Push(hr, num)
	}

	var result int

	for hr.Len() > 0 && hl.Len() > 0 {
		diff := heap.Pop(hl).(int) - heap.Pop(hr).(int)
		if diff < 0 {
			diff *= -1
		}

		result += diff
	}

	return result
}