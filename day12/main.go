package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("error opening file")
		return
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	garden := [][]rune{}

	for scanner.Scan() {
		line := scanner.Text()
		garden = append(garden, []rune(line))
	}

	price := 0
	visited := make(map[int]map[int]struct{})

	for i := 0; i < len(garden); i++ {
		for j := 0; j < len(garden[0]); j++ {
			if _, ok := visited[i][j]; ok {
				continue
			}
			area, perimeter := calcPlot(garden, garden[i][j], i, j, visited)
			price += area * perimeter
		}
	}

	fmt.Println("Total price of fencing:", price)

	// part 2

	price = 0
	visited = make(map[int]map[int]struct{})
	totalArea := 0

	for i := 0; i < len(garden); i++ {
		for j := 0; j < len(garden[0]); j++ {
			if _, ok := visited[i][j]; ok {
				continue
			}
			horSides := make(map[int]map[int][]int)
			vertSides := make(map[int]map[int][]int)
			area := calcDiscountPlot(garden, garden[i][j], i, j, i, j, horSides, vertSides, visited)
			sides := 0
			for _, rowDir := range horSides {
				for _, row := range rowDir {
					sort.Ints(row)
					prev := row[0]
					for _, num := range row {
						if prev != num-1 {
							sides++
						}
						prev = num
					}
				}
			}
			for _, colDir := range vertSides {
				for _, col := range colDir {
					sort.Ints(col)
					prev := col[0]
					for _, num := range col {
						if prev != num-1 {
							sides++
						}
						prev = num
					}
				}
			}
			price += area * sides
			totalArea += area
		}
	}

	fmt.Println("Total discounted price of fencing:", price)
}

func calcPlot(garden [][]rune, plant rune, i int, j int, visited map[int]map[int]struct{}) (int, int) {
	if i < 0 || i >= len(garden) || j < 0 || j >= len(garden[i]) {
		return 0, 1
	}
	if garden[i][j] != plant {
		return 0, 1
	}
	if _, ok := visited[i][j]; ok {
		return 0, 0
	}
	if _, ok := visited[i]; !ok {
		visited[i] = make(map[int]struct{})
	}
	visited[i][j] = struct{}{}
	moves := [][]int{[]int{1, 0}, []int{0, 1}, []int{-1, 0}, []int{0, -1}}
	area, perimeter := 1, 0

	for _, move := range moves {
		tempArea, tempPerimeter := calcPlot(garden, plant, i+move[0], j+move[1], visited)
		area += tempArea
		perimeter += tempPerimeter
	}
	return area, perimeter
}

func calcDiscountPlot(garden [][]rune, plant rune, i int, j int, prevI int, prevJ int, horSides map[int]map[int][]int, vertSides map[int]map[int][]int, visited map[int]map[int]struct{}) int {
	if i < 0 {
		if _, ok := horSides[0]; !ok {
			horSides[0] = make(map[int][]int)
		}
		horSides[0][-1] = append(horSides[0][-1], j)
		return 0
	}
	if i >= len(garden) {
		row := len(garden)
		if _, ok := horSides[row]; !ok {
			horSides[row] = make(map[int][]int)
		}
		horSides[row][1] = append(horSides[row][1], j)
		return 0
	}
	if j < 0 {
		if _, ok := vertSides[0]; !ok {
			vertSides[0] = make(map[int][]int)
		}
		vertSides[0][-1] = append(vertSides[0][-1], i)
		return 0
	}
	if j >= len(garden[i]) {
		col := len(garden[i])
		if _, ok := vertSides[col]; !ok {
			vertSides[col] = make(map[int][]int)
		}
		vertSides[col][1] = append(vertSides[col][1], i)
		return 0
	}
	if garden[i][j] != plant {
		if i != prevI {
			row := max(i, prevI)
			if _, ok := horSides[row]; !ok {
				horSides[row] = make(map[int][]int)
			}
			horSides[row][i-prevI] = append(horSides[row][i-prevI], j)
			return 0
		}
		if j != prevJ {
			col := max(j, prevJ)
			if _, ok := vertSides[col]; !ok {
				vertSides[col] = make(map[int][]int)
			}
			vertSides[col][j-prevJ] = append(vertSides[col][j-prevJ], i)
			return 0
		}
	}
	if _, ok := visited[i][j]; ok {
		return 0
	}
	if _, ok := visited[i]; !ok {
		visited[i] = make(map[int]struct{})
	}
	visited[i][j] = struct{}{}
	moves := [][]int{[]int{1, 0}, []int{0, 1}, []int{-1, 0}, []int{0, -1}}
	area := 1
	for _, move := range moves {
		area += calcDiscountPlot(garden, plant, i+move[0], j+move[1], i, j, horSides, vertSides, visited)
	}
	return area
}
