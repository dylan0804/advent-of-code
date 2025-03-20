package main

import (
	"bufio"
	"fmt"
	"os"
)

func processFile(file *os.File) [][]rune {
	scanner := bufio.NewScanner(file)

	var grid [][]rune

	for scanner.Scan() {
		line := scanner.Text()

		grid = append(grid, []rune(line))
	}

	return grid
}

func checkAllDirections(grid [][]rune, row, col, dr, dc int) bool {
	word := "XMAS"

	rows := len(grid)
	cols := len(grid[0])

	for k := 0; k < len(word); k++ {
		rr := row + k*dr
		cc := col + k*dc

		if rr < 0 || rr >= rows || cc < 0 || cc >= cols {
			return false
		}

		if grid[rr][cc] != rune(word[k]) {
			return false
		}
	}

	return true
}

func part1(file *os.File) {
	grid := processFile(file)

	directions := [][]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 1},
		{1, 1},
		{1, -1},
		{-1, -1},
	}
	
	rows := len(grid)
	cols := len(grid[0])

	var count int

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 'X' {
				for _, d := range directions {
					dr, dc := d[0], d[1]
					if checkAllDirections(grid, i, j, dr, dc) {
						count++
					}
				}
			}
		}
	}

	fmt.Println(count)
}