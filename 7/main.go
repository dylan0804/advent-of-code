package main

import (
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func stringsToInts(s []string) []int {
	ints := make([]int, len(s))

	for i, str := range s {
		num, err := strconv.Atoi(str)
		check(err)

		ints[i] = num
 	}

	return ints
}

func main() {
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

	fmt.Println(part1(file))
	fmt.Println(part2(file))
}