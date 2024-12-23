package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	size := 71
	bytes := 1024
	grid := [][]rune{}
	for i := 0; i < size; i++ {
		grid = append(grid, make([]rune, size))
		for j := 0; j < size; j++ {
			grid[i][j] = '.'
		}
	}

	for scanner.Scan() {
		if bytes == 0 {
			break
		}
		line := scanner.Text()
		split := strings.Split(line, ",")
		x, _ := strconv.Atoi(split[1])
		y, _ := strconv.Atoi(split[0])
		grid[x][y] = '#'
		bytes--
	}

	grid[size-1][size-1] = '*'

	for _, row := range grid {
		fmt.Println(string(row))
	}

	visited := make([][]bool, size)
	for i := range visited {
		visited[i] = make([]bool, size)
	}

	queue := [][]int{[]int{0, 0}}
	visited[0][0] = true
	steps := 1

	for len(queue) > 0 {
		queueLength := len(queue)
		for i := 0; i < queueLength; i++ {
			pos := queue[i]
			moves := [][]int{[]int{1, 0}, []int{-1, 0}, []int{0, 1}, []int{0, -1}}

			for _, move := range moves {
				posX := pos[0] + move[0]
				posY := pos[1] + move[1]
				if canMove(grid, posX, posY, size, visited) {
					queue = append(queue, []int{posX, posY})
					visited[posX][posY] = true
					if grid[posX][posY] == '*' {
						goto done
					}
				}
			}
		}

		queue = queue[queueLength:]
		steps++
	}

done:

	fmt.Println("Minimum steps to exit:", steps)

	// part 2

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ",")
		x, _ := strconv.Atoi(split[1])
		y, _ := strconv.Atoi(split[0])
		grid[x][y] = '#'

		if !canExit(grid, size) {
			fmt.Println("Last byte:", y, x)
			break
		}
	}

}

func canMove(grid [][]rune, posX int, posY int, size int, visited [][]bool) bool {
	if posX < 0 || posX >= size || posY < 0 || posY >= size || grid[posX][posY] == '#' || visited[posX][posY] {
		return false
	}
	return true
}

func canExit(grid [][]rune, size int) bool {
	visited := make([][]bool, size)
	for i := range visited {
		visited[i] = make([]bool, size)
	}

	queue := [][]int{[]int{0, 0}}
	visited[0][0] = true

	for len(queue) > 0 {
		queueLength := len(queue)
		for i := 0; i < queueLength; i++ {
			pos := queue[i]
			moves := [][]int{[]int{1, 0}, []int{-1, 0}, []int{0, 1}, []int{0, -1}}

			for _, move := range moves {
				posX := pos[0] + move[0]
				posY := pos[1] + move[1]
				if canMove(grid, posX, posY, size, visited) {
					queue = append(queue, []int{posX, posY})
					visited[posX][posY] = true
					if grid[posX][posY] == '*' {
						return true
					}
				}
			}
		}

		queue = queue[queueLength:]
	}

	return false
}
