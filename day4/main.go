package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("error opening file")
		return
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	grid := [][]rune{}

	for scanner.Scan() {
		line := scanner.Text()

		grid = append(grid, []rune{})
		last := len(grid) - 1

		for _, char := range line {
			grid[last] = append(grid[last], char)
		}
	}

	count := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 'X' {
				count += searchXMAS(i, j, grid)
			}
		}
	}

	fmt.Println("XMAS found:", count)

	count = 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 'X' {
				count += betterSearchXMAS(i, j, grid)
			}
		}
	}

	fmt.Println("Better XMAS found:", count)

	// part 2

	count = 0

	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			if grid[i][j] == 'A' {
				if checkCrossMAS(i, j, grid) {
					count++
				}
			}
		}
	}

	fmt.Println("CrossMAS found:", count)
}

func searchXMAS(row int, col int, grid [][]rune) int {
	word := []rune{'X', 'M', 'A', 'S'}
	count := 0

	// right
	for i := 1; i < 4; i++ {
		if isOutOfBounds(row, col+i, grid) {
			goto left
		}
		if grid[row][col+i] != word[i] {
			goto left
		}
	}
	count++

left:
	for i := 1; i < 4; i++ {
		if isOutOfBounds(row, col-i, grid) {
			goto up
		}
		if grid[row][col-i] != word[i] {
			goto up
		}
	}
	count++

up:
	for i := 1; i < 4; i++ {
		if isOutOfBounds(row-i, col, grid) {
			goto down
		}
		if grid[row-i][col] != word[i] {
			goto down
		}
	}
	count++

down:
	for i := 1; i < 4; i++ {
		if isOutOfBounds(row+i, col, grid) {
			goto upright
		}
		if grid[row+i][col] != word[i] {
			goto upright
		}
	}
	count++

upright:
	for i := 1; i < 4; i++ {
		if isOutOfBounds(row-i, col+i, grid) {
			goto downright
		}
		if grid[row-i][col+i] != word[i] {
			goto downright
		}
	}
	count++

downright:
	for i := 1; i < 4; i++ {
		if isOutOfBounds(row+i, col+i, grid) {
			goto upleft
		}
		if grid[row+i][col+i] != word[i] {
			goto upleft
		}
	}
	count++

upleft:
	for i := 1; i < 4; i++ {
		if isOutOfBounds(row-i, col-i, grid) {
			goto downleft
		}
		if grid[row-i][col-i] != word[i] {
			goto downleft
		}
	}
	count++

downleft:
	for i := 1; i < 4; i++ {
		if isOutOfBounds(row+i, col-i, grid) {
			goto done
		}
		if grid[row+i][col-i] != word[i] {
			goto done
		}
	}
	count++

done:
	return count
}

func betterSearchXMAS(row int, col int, grid [][]rune) int {
	word := []rune{'X', 'M', 'A', 'S'}
	count := 0
	moves := [][]int{[]int{1, 0}, []int{0, 1}, []int{-1, 0}, []int{0, -1}, []int{1, 1}, []int{-1, -1}, []int{1, -1}, []int{-1, 1}}

	for _, move := range moves {
		for i := 1; i < 4; i++ {
			newRow, newCol := row+move[0]*i, col+move[1]*i
			if isOutOfBounds(newRow, newCol, grid) {
				break
			}
			if grid[newRow][newCol] != word[i] {
				break
			}
			if i == 3 {
				count++
			}
		}
	}

	return count
}

func isOutOfBounds(row int, col int, grid [][]rune) bool {
	if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[0]) {
		return true
	}
	return false
}

func checkCrossMAS(row int, col int, grid [][]rune) bool {
	if (grid[row-1][col-1] == 'M' && grid[row+1][col+1] == 'S') || (grid[row-1][col-1] == 'S' && grid[row+1][col+1] == 'M') {
		if (grid[row-1][col+1] == 'M' && grid[row+1][col-1] == 'S') || (grid[row-1][col+1] == 'S' && grid[row+1][col-1] == 'M') {
			return true
		}
	}

	return false
}
