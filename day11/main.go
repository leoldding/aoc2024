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

	scanner.Scan()
	line := scanner.Text()
	stones := strings.Split(line, " ")

	fmt.Println("Number of stones:", simulation(stones))

	stoneMap := make(map[string]int)
	for _, stone := range stones {
		stoneMap[stone]++
	}

	fmt.Println("Number of stones:", betterSimulation(stoneMap))
}

func simulation(stones []string) int {
	for i := 0; i < 25; i++ {
		length := len(stones)
		for _, stone := range stones {
			if stone == "0" {
				stones = append(stones, "1")
			} else if len(stone)%2 == 0 {
				half := len(stone) / 2

				stoneNum, _ := strconv.Atoi(stone[:half])
				newStone1 := strconv.Itoa(stoneNum)

				stoneNum, _ = strconv.Atoi(stone[half:])
				newStone2 := strconv.Itoa(stoneNum)

				stones = append(stones, []string{newStone1, newStone2}...)
			} else {
				stoneNum, _ := strconv.Atoi(stone)
				newStone := strconv.Itoa(stoneNum * 2024)
				stones = append(stones, newStone)
			}
		}
		stones = stones[length:]
	}
	return len(stones)
}

func betterSimulation(stoneMap map[string]int) int {
	for i := 0; i < 75; i++ {
		newStoneMap := map[string]int{}
		for stone, count := range stoneMap {
			if stone == "0" {
				newStoneMap["1"] += count
			} else if len(stone)%2 == 0 {
				half := len(stone) / 2

				stoneNum, _ := strconv.Atoi(stone[:half])
				newStone1 := strconv.Itoa(stoneNum)
				newStoneMap[newStone1] += count

				stoneNum, _ = strconv.Atoi(stone[half:])
				newStone2 := strconv.Itoa(stoneNum)
				newStoneMap[newStone2] += count
			} else {
				stoneNum, _ := strconv.Atoi(stone)
				newStone := strconv.Itoa(stoneNum * 2024)
				newStoneMap[newStone] += count
			}
		}
		stoneMap = newStoneMap
	}

	stones := 0
	for _, count := range stoneMap {
		stones += count
	}

	return stones
}
