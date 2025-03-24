package main

import (
	"os"
)

func checkDiagonal(grid [][]rune, row, col int, upr, btl []int) bool {
	rows := len(grid)
	cols := len(grid[0])

	rupr := row + upr[0]
	cupr := col + upr[1]

	if rupr < 0 || rupr >= rows || cupr < 0 || cupr >= cols {
		return false
	}

	rbtl := row + btl[0]
	cbtl := col + btl[1]

	if rbtl < 0 || rbtl >= rows || cbtl < 0 || cbtl >= cols {
		return false
	}

	if (grid[rupr][cupr] == 'S' && grid[rbtl][cbtl] == 'M') || (grid[rupr][cupr] == 'M' && grid[rbtl][cbtl] == 'S') {
		return true
	}

	return false
}

func part2(file *os.File) int {
	grid := processFile(file)

	directions := [][]int{
		{-1, 1},
		{1, -1},
		{-1, -1},
		{1, 1},
	}
	
	rows := len(grid)
	cols := len(grid[0])

	var count int

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 'A' {
				// up right and bottom left
				upr := directions[0]
				btl := directions[1]

				// up left and bottom right
				upl := directions[2]
				btr := directions[3]

				if checkDiagonal(grid, i, j, upr, btl) && checkDiagonal(grid, i, j, upl, btr) {
					count++
				}
			}
		}
	}

	return count
}