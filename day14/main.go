package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("error opening file")
		return
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	positions := [][]int{}
	velocities := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		separators := regexp.MustCompile("[=, ]")
		result := separators.Split(line, -1)

		posX, _ := strconv.Atoi(result[1])
		posY, _ := strconv.Atoi(result[2])
		positions = append(positions, []int{posX, posY})

		velX, _ := strconv.Atoi(result[4])
		velY, _ := strconv.Atoi(result[5])
		velocities = append(velocities, []int{velX, velY})
	}

	seconds, height, width := 100, 103, 101
	for second := 0; second < seconds; second++ {
		for i, _ := range positions {
			positions[i][0] = (positions[i][0] + velocities[i][0]) % width
			if positions[i][0] < 0 {
				positions[i][0] += width
			}
			positions[i][1] = (positions[i][1] + velocities[i][1]) % height
			if positions[i][1] < 0 {
				positions[i][1] += height
			}
		}
	}

	safety := 1

	quads := []int{0, 0, 0, 0}

	vertMid, horMid := int(width/2), int(height/2)

	for _, position := range positions {
		if position[0] < vertMid {
			if position[1] < horMid {
				quads[0]++
			} else if position[1] > horMid {
				quads[1]++
			}
		} else if position[0] > vertMid {
			if position[1] < horMid {
				quads[2]++
			} else if position[1] > horMid {
				quads[3]++
			}
		}
	}

	for _, quad := range quads {
		safety *= quad
	}

	fmt.Println("Safety score:", safety)

	// part 2

	grid := [][]string{}
	for {
		grid = [][]string{}
		for i := 0; i < width; i++ {
			grid = append(grid, make([]string, height))
		}

		for i := 0; i < len(positions); i++ {
			positions[i][0] = (positions[i][0] + velocities[i][0]) % width
			if positions[i][0] < 0 {
				positions[i][0] += width
			}
			positions[i][1] = (positions[i][1] + velocities[i][1]) % height
			if positions[i][1] < 0 {
				positions[i][1] += height
			}

			grid[positions[i][0]][positions[i][1]] = "#"
		}
		seconds++

		for _, row := range grid {
			consecutive := 0
			prev := ""
			for _, col := range row {
				if prev == "#" && col == "#" {
					consecutive++
				}
				prev = col
			}
			if consecutive >= 25 {
				goto end
			}
		}
	}
end:
	for _, row := range grid {
		for i, col := range row {
			if col == "" {
				row[i] = "."
			}
		}
		fmt.Println(row)
	}
	fmt.Println("Time to tree:", seconds)
}
