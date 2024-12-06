package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	origX, origY := -1, -1 // use in part 2
	posX, posY := -1, -1
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "^") {
			origX = strings.Index(line, "^")
			origY = row
			posX = strings.Index(line, "^")
			posY = row
			line = strings.Replace(line, "^", "X", 1)
		}
		grid = append(grid, []rune(line))
		row++
	}

	spaces := 1
	direction := 0 // 0-up, 1-right, 2-down, 3-left

	for {
		newX, newY := posX, posY
		switch direction {
		case 0:
			newY = newY - 1
		case 1:
			newX = newX + 1
		case 2:
			newY = newY + 1
		case 3:
			newX = newX - 1
		}

		if newX < 0 || newX >= len(grid[0]) || newY < 0 || newY >= len(grid) {
			break
		}

		mark := grid[newY][newX]
		switch mark {
		case '.':
			posX, posY = newX, newY
			spaces++
			grid[newY][newX] = 'X'
		case 'X':
			posX, posY = newX, newY
		case '#':
			direction = (direction + 1) % 4
		}
	}

	fmt.Println("Guard unique positions:", spaces)

	// part 2

	for i, row := range grid {
		for j, char := range row {
			if char == 'X' {
				grid[i][j] = '.'
			}
		}
	}

	loops := 0
	moves := 0
	maxMoves := len(grid) * len(grid[0]) * 2

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '#' {
				continue
			}
			grid[i][j] = '#'
			posX, posY := origX, origY
			direction = 0
			moves = 0

			for {
				newX, newY := posX, posY
				switch direction {
				case 0:
					newY = newY - 1
				case 1:
					newX = newX + 1
				case 2:
					newY = newY + 1
				case 3:
					newX = newX - 1
				}

				if newX < 0 || newX >= len(grid[0]) || newY < 0 || newY >= len(grid) {
					break
				}

				mark := grid[newY][newX]
				switch mark {
				case '.':
					posX, posY = newX, newY
					moves++
				case '#':
					direction = (direction + 1) % 4
				}

				if moves > maxMoves {
					loops++
					break
				}
			}
			grid[i][j] = '.'
		}
	}

	fmt.Println("Obstruction possible positions:", loops)
}
