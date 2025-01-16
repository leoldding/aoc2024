package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

type Position struct {
	x int
	y int
}

var moves = [][2]int{[2]int{1, 0}, [2]int{-1, 0}, [2]int{0, 1}, [2]int{0, -1}}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("error opening file")
		return
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	track := [][]rune{}
	start := Position{}
	end := Position{}

	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		track = append(track, []rune(line))
		col := strings.Index(line, "S")
		if col != -1 {
			start.x = row
			start.y = col
		}
		col = strings.Index(line, "E")
		if col != -1 {
			end.x = row
			end.y = col
		}
		row++
	}

	route := getRoute(track, start, end)
	route = append(route, end)
	dists := make(map[Position]map[Position]int)
	calcDist := func(start, end Position) int {
		return int(math.Abs(float64(start.x-end.x)) + math.Abs(float64(start.y-end.y)))
	}
	for i := 0; i < len(route); i++ {
		dists[route[i]] = make(map[Position]int)
		for j := i + 1; j < len(route); j++ {
			dists[route[i]][route[j]] = calcDist(route[i], route[j])
		}
	}

	fmt.Println("Fastest time without cheating:", len(route)-1)

	cheatCount := 0
	for i, start := range route {
		for end, dist := range dists[start] {
			if dist <= 2 {
				if slices.Index(route, end)-i-cheatBFS(track, start, end) >= 100 {
					cheatCount++
				}
			}
		}
	}
	fmt.Println("Cheats (2) that save 100 ps:", cheatCount)

	// part 2

	cheatCount = 0
	for i, start := range route {
		for end, dist := range dists[start] {
			if dist <= 20 {
				if slices.Index(route, end)-i-cheatBFS(track, start, end) >= 100 {
					cheatCount++
				}
			}
		}
	}
	fmt.Println("Cheats (20) that save 100 ps:", cheatCount)
}

func getRoute(track [][]rune, start Position, end Position) []Position {
	type PosRoute struct {
		position Position
		route    []Position
	}
	picoseconds := 0
	queue := []PosRoute{PosRoute{start, []Position{}}}
	visited := make([][]bool, len(track))
	for i := 0; i < len(track); i++ {
		visited[i] = make([]bool, len(track[i]))
	}
	visited[start.x][start.y] = true

	for len(queue) > 0 {
		length := len(queue)

		for i := 0; i < length; i++ {
			cur := queue[i]
			pos := cur.position
			if track[pos.x][pos.y] == 'E' {
				return cur.route
			}
			for _, move := range moves {
				posX := pos.x + move[0]
				posY := pos.y + move[1]
				if posX < 0 || posX >= len(track) || posY < 0 || posY >= len(track[0]) {
					continue
				}
				if track[posX][posY] == '#' {
					continue
				}
				if visited[posX][posY] {
					continue
				}
				newPos := Position{posX, posY}
				queue = append(queue, PosRoute{newPos, append(cur.route, cur.position)})
				visited[posX][posY] = true
			}
		}
		queue = queue[length:]
		picoseconds++
	}
	return []Position{}
}

func cheatBFS(track [][]rune, start Position, end Position) int {
	picoseconds := 0
	queue := []Position{start}
	visited := make([][]bool, len(track))
	for i := 0; i < len(track); i++ {
		visited[i] = make([]bool, len(track[i]))
	}
	visited[start.x][start.y] = true

	for len(queue) > 0 {
		length := len(queue)

		for i := 0; i < length; i++ {
			pos := queue[i]
			if pos.x == end.x && pos.y == end.y {
				return picoseconds
			}
			for _, move := range moves {
				posX := pos.x + move[0]
				posY := pos.y + move[1]
				if posX < 0 || posX >= len(track) || posY < 0 || posY >= len(track[0]) {
					continue
				}
				if visited[posX][posY] {
					continue
				}
				queue = append(queue, Position{posX, posY})
				visited[posX][posY] = true
			}
		}
		queue = queue[length:]
		picoseconds++
	}

	return -1
}
