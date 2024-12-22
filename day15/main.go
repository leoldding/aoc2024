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

	getMap := true
	warehouse := [][]rune{}
	warehouse2 := [][]rune{}
	pos := []int{}
	pos2 := []int{}
	moves := ""
	row := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			getMap = false
		} else if getMap {
			ind := strings.Index(line, "@")
			warehouse = append(warehouse, []rune(line))
			if ind != -1 {
				pos = []int{row, ind}
				pos2 = []int{row, ind * 2}
			}
			row++

			// build warehouse for part 2
			newLine := ""
			for _, ch := range line {
				switch ch {
				case '#':
					newLine += "##"
				case 'O':
					newLine += "[]"
				case '.':
					newLine += ".."
				case '@':
					newLine += "@."
				}
			}
			warehouse2 = append(warehouse2, []rune(newLine))
		} else {
			moves += line
		}
	}

	for _, move := range moves {
		switch move {
		case '<':
			for i := pos[1]; i >= 0; i-- {
				if warehouse[pos[0]][i] == '.' {
					for j := i; j <= pos[1]; j++ {
						warehouse[pos[0]][j] = warehouse[pos[0]][j+1]
					}
					warehouse[pos[0]][pos[1]] = '.'
					pos[1]--
					break
				} else if warehouse[pos[0]][i] == '#' {
					break
				}
			}
		case '>':
			for i := pos[1]; i < len(warehouse[pos[0]]); i++ {
				if warehouse[pos[0]][i] == '.' {
					for j := i; j >= pos[1]; j-- {
						warehouse[pos[0]][j] = warehouse[pos[0]][j-1]
					}
					warehouse[pos[0]][pos[1]] = '.'
					pos[1]++
					break
				} else if warehouse[pos[0]][i] == '#' {
					break
				}
			}
		case '^':
			for i := pos[0]; i >= 0; i-- {
				if warehouse[i][pos[1]] == '.' {
					for j := i; j <= pos[0]; j++ {
						warehouse[j][pos[1]] = warehouse[j+1][pos[1]]
					}
					warehouse[pos[0]][pos[1]] = '.'
					pos[0]--
					break
				} else if warehouse[i][pos[1]] == '#' {
					break
				}
			}
		case 'v':
			for i := pos[0]; i < len(warehouse); i++ {
				if warehouse[i][pos[1]] == '.' {
					for j := i; j >= pos[0]; j-- {
						warehouse[j][pos[1]] = warehouse[j-1][pos[1]]
					}
					warehouse[pos[0]][pos[1]] = '.'
					pos[0]++
					break
				} else if warehouse[i][pos[1]] == '#' {
					break
				}
			}
		}
	}

	gps := 0

	for i := 0; i < len(warehouse); i++ {
		for j := 0; j < len(warehouse[i]); j++ {
			if warehouse[i][j] == 'O' {
				gps += i*100 + j
			}
		}
	}

	fmt.Println("Sum of GPS:", gps)

	// part 2

	for _, move := range moves {
		switch move {
		case '<':
			for i := pos2[1]; i >= 0; i-- {
				if warehouse2[pos2[0]][i] == '.' {
					for j := i; j <= pos2[1]; j++ {
						warehouse2[pos2[0]][j] = warehouse2[pos2[0]][j+1]
					}
					warehouse2[pos2[0]][pos2[1]] = '.'
					pos2[1]--
					break
				} else if warehouse2[pos2[0]][i] == '#' {
					break
				}
			}
		case '>':
			for i := pos2[1]; i < len(warehouse2[pos2[0]]); i++ {
				if warehouse2[pos2[0]][i] == '.' {
					for j := i; j >= pos2[1]; j-- {
						warehouse2[pos2[0]][j] = warehouse2[pos2[0]][j-1]
					}
					warehouse2[pos2[0]][pos2[1]] = '.'
					pos2[1]++
					break
				} else if warehouse2[pos2[0]][i] == '#' {
					break
				}
			}
		case '^':
			spaces := []int{pos2[1]}
			boxes := [][]int{}
			for i := pos2[0] - 1; i >= 0; i-- {
				length := len(spaces)
				for j := 0; j < length; j++ {
					if warehouse2[i][spaces[j]] == '[' {
						boxes = append(boxes, []int{i, spaces[j]})
						spaces = append(spaces, []int{spaces[j], spaces[j] + 1}...)
					} else if warehouse2[i][spaces[j]] == ']' {
						boxes = append(boxes, []int{i, spaces[j] - 1})
						spaces = append(spaces, []int{spaces[j] - 1, spaces[j]}...)
					} else if warehouse2[i][spaces[j]] == '#' {
						goto end
					}
				}
				spaces = spaces[length:]
				if len(spaces) == 0 {
					break
				}
			}
			for i := len(boxes) - 1; i >= 0; i-- {
				box := boxes[i]
				warehouse2[box[0]-1][box[1]], warehouse2[box[0]-1][box[1]+1] = '[', ']'
				warehouse2[box[0]][box[1]], warehouse2[box[0]][box[1]+1] = '.', '.'
			}
			warehouse2[pos2[0]-1][pos2[1]] = '@'
			warehouse2[pos2[0]][pos2[1]] = '.'
			pos2[0]--
		case 'v':
			spaces := []int{pos2[1]}
			boxes := [][]int{}
			for i := pos2[0] + 1; i < len(warehouse2); i++ {
				length := len(spaces)
				for j := 0; j < length; j++ {
					if warehouse2[i][spaces[j]] == '[' {
						boxes = append(boxes, []int{i, spaces[j]})
						spaces = append(spaces, []int{spaces[j], spaces[j] + 1}...)
					} else if warehouse2[i][spaces[j]] == ']' {
						boxes = append(boxes, []int{i, spaces[j] - 1})
						spaces = append(spaces, []int{spaces[j] - 1, spaces[j]}...)
					} else if warehouse2[i][spaces[j]] == '#' {
						goto end
					}
				}
				spaces = spaces[length:]
				if len(spaces) == 0 {
					break
				}
			}
			for i := len(boxes) - 1; i >= 0; i-- {
				box := boxes[i]
				warehouse2[box[0]+1][box[1]], warehouse2[box[0]+1][box[1]+1] = '[', ']'
				warehouse2[box[0]][box[1]], warehouse2[box[0]][box[1]+1] = '.', '.'
			}
			warehouse2[pos2[0]+1][pos2[1]] = '@'
			warehouse2[pos2[0]][pos2[1]] = '.'
			pos2[0]++
		}
	end:
	}

	gps = 0

	for i := 0; i < len(warehouse2); i++ {
		for j := 0; j < len(warehouse2[i]); j++ {
			if warehouse2[i][j] == '[' {
				gps += i*100 + j
			}
		}
	}

	fmt.Println("Sum of WideGPS:", gps)
}
