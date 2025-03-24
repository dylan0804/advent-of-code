package main

var directions = [][]int{
	{-1, 0}, // up
	{0, 1}, // right
	{1, 0}, // bottom
	{0, -1}, // left
}

func trackPosition(count, startDirection int, grid []string, row, col int, visited [][]bool) int {
	if !visited[row][col] {
		visited[row][col] = true
		count++
	}

	dir := directions[startDirection]
	
	rr := row + dir[0]
	cc := col + dir[1]	

	if rr < 0 || rr >= len(grid) || cc < 0 || cc >= len(grid[0]) {
		return count
	}

	if grid[rr][cc] == '#' {
		startDirection = (startDirection + 1) % 4
		dir = directions[startDirection]
		rr = row + dir[0]
		cc = col + dir[1]
	}

	return trackPosition(count, startDirection, grid, rr, cc, visited)
}

func part1(grid []string) int {
	count, startDirection := 0, 0

	rows := len(grid)
	cols := len(grid[0])

	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}
	
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == '^' {
				count = trackPosition(0, startDirection, grid, row, col, visited)
				break
			}
		}
	}

	return count
}