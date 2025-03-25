package main

import (
	"fmt"
	"io"
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
	defer file.Close()

	content, err := io.ReadAll(file)
	check(err)

	fmt.Println(part1(string(content)))
}