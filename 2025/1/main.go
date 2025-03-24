package main

import "fmt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}