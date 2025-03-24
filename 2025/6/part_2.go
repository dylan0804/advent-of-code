package main

import "fmt"

func copyGrid(grid []string) []string {
	testGrid := make([]string, len(grid))

	copy(testGrid, grid)

	return testGrid
}

func replaceAtIndex(s string, col int, char byte) string {
	temp := []byte(s)
	temp[col] = char
	return string(temp)
}

func addObstacle(grid []string, row, col, guardRow, guardCol int) bool {
	testGrid := copyGrid(grid)
	testGrid[row] = replaceAtIndex(testGrid[row], col, '#')

	visited := make(map[string]bool)

	curRow, curCol, dir := guardRow, guardCol, 0

	for {
		state := fmt.Sprintf("%d,%d,%d", curRow, curCol, dir)

		// found a loop
		if visited[state] {
			return true
		}

		visited[state] = true

		nextRow := curRow + directions[dir][0]
		nextCol := curCol + directions[dir][1]

		// guard is out it means that adding an obstacle here is useless
		if nextRow < 0 || nextRow >= len(testGrid) || nextCol < 0 || nextCol >= len(testGrid[0]) {
			return false
		}

		if testGrid[nextRow][nextCol] == '#' {
			dir = (dir + 1) % 4
		} else {
			curRow, curCol = nextRow, nextCol
		}
	}
}

func part2(grid []string) int {
	count := 0
	guardRow, guardCol := 0, 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '^' {
				guardRow, guardCol = i, j
				break
			}
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '#' || grid[i][j] == '^' {
				continue
			}

			if addObstacle(grid, i, j, guardRow, guardCol) {
				count++
			}
		}
	}
	
	return count
}