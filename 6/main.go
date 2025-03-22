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
	defer file.Close()

	content, err := io.ReadAll(file)
	check(err)

	grid := strings.Split(string(content), "\n")

	fmt.Println(part1(grid))
}