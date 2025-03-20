package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("./input.txt")
	check(err)

	fmt.Println(part1(file))
	fmt.Println(part2(file))
}