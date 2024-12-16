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

	topMap := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		topMap = append(topMap, []int{})
		last := len(topMap) - 1
		for _, height := range line {
			topMap[last] = append(topMap[last], int(height-'0'))
		}
	}

	score := 0

	for i := 0; i < len(topMap); i++ {
		for j := 0; j < len(topMap[i]); j++ {
			if topMap[i][j] == 0 {
				score += len(searchTrailhead(topMap, i, j, 0, [][]int{}))
			}
		}
	}

	fmt.Println("Score of all trailheads:", score)

	// part 2

	rating := 0

	for i := 0; i < len(topMap); i++ {
		for j := 0; j < len(topMap[i]); j++ {
			if topMap[i][j] == 0 {
				rating += searchDistinctTrailheads(topMap, i, j, 0)
			}
		}
	}

	fmt.Println("Rating of all trailheads:", rating)
}

func searchTrailhead(topMap [][]int, i int, j int, cur int, peaks [][]int) [][]int {
	if i < 0 || i >= len(topMap) || j < 0 || j >= len(topMap[i]) {
		return peaks
	}
	if topMap[i][j] != cur {
		return peaks
	}
	if cur == 9 {
		if !peakExists(i, j, peaks) {
			peaks = append(peaks, []int{i, j})
		}
		return peaks
	}
	moves := [][]int{[]int{1, 0}, []int{0, 1}, []int{-1, 0}, []int{0, -1}}

	for _, move := range moves {
		peaks = searchTrailhead(topMap, i+move[0], j+move[1], cur+1, peaks)
	}

	return peaks
}

func peakExists(i int, j int, peaks [][]int) bool {
	for _, peak := range peaks {
		if peak[0] == i && peak[1] == j {
			return true
		}
	}
	return false
}

func searchDistinctTrailheads(topMap [][]int, i int, j int, cur int) int {
	if i < 0 || i >= len(topMap) || j < 0 || j >= len(topMap[i]) {
		return 0
	}
	if topMap[i][j] != cur {
		return 0
	}
	if cur == 9 {
		return 1
	}
	moves := [][]int{[]int{1, 0}, []int{0, 1}, []int{-1, 0}, []int{0, -1}}
	trails := 0

	for _, move := range moves {
		trails += searchDistinctTrailheads(topMap, i+move[0], j+move[1], cur+1)
	}

	return trails
}
