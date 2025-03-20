package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("./input.txt")
	check(err)

	content, err := io.ReadAll(file)
	check(err)

	// split sections into rules and sequence
	sections := strings.Split(string(content), "\n\n")

	fmt.Println(part1(sections))
}