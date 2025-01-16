package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type pos struct {
	x    int
	y    int
	dir  int
	pts  int
	path []coord
}

type coord struct {
	x int
	y int
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("error opening file")
		return
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	maze := [][]rune{}
	start := pos{}
	start.dir = 1 // 0123 -> nesw
	row := 0

	for scanner.Scan() {
		line := scanner.Text()
		maze = append(maze, []rune(line))

		if strings.Index(line, "S") != -1 {
			start.x = row
			start.y = strings.Index(line, "S")
		}
		row++
	}

	queue := []pos{start}
	visited := make(map[int]map[int]map[int]int) // x, y, dir
	paths := [][]coord{}
	score := math.MaxInt32

	for len(queue) > 0 {
		queueLength := len(queue)

		for i := 0; i < queueLength; i++ {
			curPos := queue[i]

			if maze[curPos.x][curPos.y] == 'E' {
				if curPos.pts < score {
					score = curPos.pts
					paths = [][]coord{curPos.path}
				} else if curPos.pts == score {
					paths = append(paths, curPos.path)
				}
			}

			newPos := pos{-1, -1, -1, -1, append(curPos.path, coord{curPos.x, curPos.y})}
			switch curPos.dir {
			case 0:
				if maze[curPos.x-1][curPos.y] != '#' {
					newPos.x = curPos.x - 1
					newPos.y = curPos.y
					newPos.dir = curPos.dir
					newPos.pts = curPos.pts + 1
				}
			case 1:
				if maze[curPos.x][curPos.y+1] != '#' {
					newPos.x = curPos.x
					newPos.y = curPos.y + 1
					newPos.dir = curPos.dir
					newPos.pts = curPos.pts + 1
				}
			case 2:
				if maze[curPos.x+1][curPos.y] != '#' {
					newPos.x = curPos.x + 1
					newPos.y = curPos.y
					newPos.dir = curPos.dir
					newPos.pts = curPos.pts + 1
				}
			case 3:
				if maze[curPos.x][curPos.y-1] != '#' {
					newPos.x = curPos.x
					newPos.y = curPos.y - 1
					newPos.dir = curPos.dir
					newPos.pts = curPos.pts + 1
				}
			}

			if newPos.x != -1 {
				if _, ok := visited[newPos.x][newPos.y][newPos.dir]; !ok {
					queue = append(queue, newPos)
					if _, ok := visited[newPos.x]; !ok {
						visited[newPos.x] = make(map[int]map[int]int)
					}
					if _, ok := visited[newPos.x][newPos.y]; !ok {
						visited[newPos.x][newPos.y] = make(map[int]int)
					}
					visited[newPos.x][newPos.y][newPos.dir] = newPos.pts
				} else {
					if newPos.pts <= visited[newPos.x][newPos.y][newPos.dir] {
						visited[newPos.x][newPos.y][newPos.dir] = newPos.pts
						queue = append(queue, newPos)
					}
				}
			}

			left := pos{curPos.x, curPos.y, curPos.dir - 1, curPos.pts + 1000, append([]coord{}, curPos.path...)}
			if left.dir == -1 {
				left.dir = 3
			}
			if _, ok := visited[left.x][left.y][left.dir]; !ok {
				queue = append(queue, left)
				if _, ok := visited[left.x]; !ok {
					visited[left.x] = make(map[int]map[int]int)
				}
				if _, ok := visited[left.x][left.y]; !ok {
					visited[left.x][left.y] = make(map[int]int)
				}
				visited[left.x][left.y][left.dir] = left.pts
			} else {
				if left.pts <= visited[left.x][left.y][left.dir] {
					visited[left.x][left.y][left.dir] = left.pts
					queue = append(queue, left)
				}
			}

			right := pos{curPos.x, curPos.y, curPos.dir + 1, curPos.pts + 1000, append([]coord{}, curPos.path...)}
			if right.dir == 4 {
				right.dir = 0
			}
			if _, ok := visited[right.x][right.y][right.dir]; !ok {
				queue = append(queue, right)
				if _, ok := visited[right.x]; !ok {
					visited[right.x] = make(map[int]map[int]int)
				}
				if _, ok := visited[right.x][right.y]; !ok {
					visited[right.x][right.y] = make(map[int]int)
				}
				visited[right.x][right.y][right.dir] = right.pts
			} else {
				if right.pts <= visited[right.x][right.y][right.dir] {
					visited[right.x][right.y][right.dir] = right.pts
					queue = append(queue, right)
				}
			}
		}

		queue = queue[queueLength:]
	}

	fmt.Println("Lowest score:", score)

	tiles := 1 // start tile
	used := make(map[coord]struct{})
	for _, path := range paths {
		for _, tile := range path {
			if _, ok := used[tile]; !ok {
				tiles++
				used[tile] = struct{}{}
			}
		}
	}
	fmt.Println("Tiles on best paths:", tiles)
}
