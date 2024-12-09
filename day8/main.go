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

	antennas := make(map[rune][][]int)
	row := 0
	col := 0

	for scanner.Scan() {
		line := scanner.Text()
		for i, val := range line {
			col = i
			if val == '.' {
				continue
			}
			antennas[val] = append(antennas[val], []int{row, i})
		}

		row++
	}
	col++

	antinodes := 0
	positions := make(map[int]map[int]struct{})

	for _, antennaPos := range antennas {
		for i := 0; i < len(antennaPos); i++ {
			for j := i + 1; j < len(antennaPos); j++ {
				xDiff := antennaPos[j][0] - antennaPos[i][0]
				yDiff := antennaPos[j][1] - antennaPos[i][1]

				posX := antennaPos[i][0] - xDiff
				posY := antennaPos[i][1] - yDiff

				if posX >= 0 && posX < col && posY >= 0 && posY < row {
					if _, ok := positions[posY]; !ok {
						positions[posY] = make(map[int]struct{})
					}
					if _, ok := positions[posY][posX]; !ok {
						positions[posY][posX] = struct{}{}
						antinodes++
					}
				}

				posX = antennaPos[j][0] + xDiff
				posY = antennaPos[j][1] + yDiff

				if posX >= 0 && posX < col && posY >= 0 && posY < row {
					if _, ok := positions[posY]; !ok {
						positions[posY] = make(map[int]struct{})
					}
					if _, ok := positions[posY][posX]; !ok {
						positions[posY][posX] = struct{}{}
						antinodes++
					}
				}
			}
		}
	}

	fmt.Println("Number of unique antinode positions:", antinodes)

	// part 2

	antinodes = 0
	positions = make(map[int]map[int]struct{})

	for _, antennaPos := range antennas {
		for i := 0; i < len(antennaPos); i++ {
			for j := i + 1; j < len(antennaPos); j++ {
				xDiff := antennaPos[j][0] - antennaPos[i][0]
				yDiff := antennaPos[j][1] - antennaPos[i][1]

				posX := antennaPos[i][0]
				posY := antennaPos[i][1]

				for posX >= 0 && posX < col && posY >= 0 && posY < row {
					if _, ok := positions[posY]; !ok {
						positions[posY] = make(map[int]struct{})
					}
					if _, ok := positions[posY][posX]; !ok {
						positions[posY][posX] = struct{}{}
						antinodes++
					}

					posX -= xDiff
					posY -= yDiff
				}

				posX = antennaPos[i][0] + xDiff
				posY = antennaPos[i][1] + yDiff

				for posX >= 0 && posX < col && posY >= 0 && posY < row {
					if _, ok := positions[posY]; !ok {
						positions[posY] = make(map[int]struct{})
					}
					if _, ok := positions[posY][posX]; !ok {
						positions[posY][posX] = struct{}{}
						antinodes++
					}

					posX += xDiff
					posY += yDiff
				}
			}
		}
	}

	fmt.Println("New number of unique antinode positions:", antinodes)
}
