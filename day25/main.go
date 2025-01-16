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

	newSchem := true
	isLock := false
	heights := [5]int{}

	locks := [][5]int{}
	keys := [][5]int{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			if isLock {
				locks = append(locks, heights)
			} else {
				keys = append(keys, heights)
			}

			newSchem = true
			isLock = false
			heights = [5]int{}
			continue
		}

		if newSchem {
			if line == "#####" {
				isLock = true
			}
			newSchem = false
		}

		for i, ch := range line {
			if ch == '#' {
				heights[i]++
			}
		}
	}

	pairs := 0

	for _, lock := range locks {
		for _, key := range keys {
			canFit := true
			for i := 0; i < 5; i++ {
				if lock[i]+key[i] > 7 {
					canFit = false
					break
				}
			}
			if canFit {
				pairs++
			}
		}
	}

	fmt.Println("Lock and key pairs:", pairs)
}
