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

	scanner := bufio.NewScanner(input)

	safe := 0

	for scanner.Scan() {
		line := scanner.Text()
		report := strings.Fields(line)
		levels := []int{}
		for _, level := range report {
			num, _ := strconv.Atoi(level)
			levels = append(levels, num)
		}

		if safeInc(levels) || safeDec(levels) {
			safe++
		}
	}

	input.Close()
	fmt.Println("Safe reports:", safe)

	// part 2

	input, err = os.Open("input.txt")
	if err != nil {
		fmt.Println("error opening file")
		return
	}
	defer input.Close()

	scanner2 := bufio.NewScanner(input)

	safeDamp := 0

	for scanner2.Scan() {
		line := scanner2.Text()
		report := strings.Fields(line)
		levels := []int{}
		for _, level := range report {
			num, _ := strconv.Atoi(level)
			levels = append(levels, num)
		}

		if safeInc(levels) || safeDec(levels) {
			safeDamp++
		} else {
			for i := 0; i < len(levels); i++ {
				dampenedLevels := make([]int, len(levels)-1)
				copy(dampenedLevels, levels[:i])
				copy(dampenedLevels[i:], levels[i+1:])
				if safeInc(dampenedLevels) || safeDec(dampenedLevels) {
					safeDamp++
					break
				}
			}
		}
	}

	fmt.Println("Safe dampened reports:", safeDamp)
}

func safeInc(arr []int) bool {
	prev := arr[0]
	for i := 1; i < len(arr); i++ {
		diff := arr[i] - prev
		if diff < 1 || diff > 3 {
			return false
		}
		prev = arr[i]
	}

	return true
}

func safeDec(arr []int) bool {
	prev := arr[0]
	for i := 1; i < len(arr); i++ {
		diff := prev - arr[i]
		if diff < 1 || diff > 3 {
			return false
		}
		prev = arr[i]
	}

	return true
}
