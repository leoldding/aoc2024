package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("error opening file")
		return
	}

	scanner := bufio.NewScanner(input)
	scanner.Scan()
	line := scanner.Text()

	fileSystem := []int{}
	id := 0

	for i, val := range line {
		num, _ := strconv.Atoi(string(val))
		if i%2 == 0 {
			for j := 0; j < num; j++ {
				fileSystem = append(fileSystem, id)
			}
			id++
		} else {
			for j := 0; j < num; j++ {
				fileSystem = append(fileSystem, -1)
			}
		}
	}

	left, right := 0, len(fileSystem)-1

	for left < right {
		if fileSystem[left] != -1 {
			left++
			continue
		}

		if fileSystem[right] == -1 {
			right--
			continue
		}

		fileSystem[left], fileSystem[right] = fileSystem[right], fileSystem[left]
	}

	checksum := 0

	for i, num := range fileSystem {
		if num == -1 {
			break
		}
		checksum += i * num
	}

	input.Close()

	fmt.Println("File system checksum:", checksum)

	// part 2

	input, err = os.Open("input.txt")
	if err != nil {
		fmt.Println("error opening file")
		return
	}
	defer input.Close()

	scanner = bufio.NewScanner(input)
	scanner.Scan()
	line = scanner.Text()

	fileSystem2 := [][]int{}
	id = 0
	idx := 0

	for i, val := range line {
		num, _ := strconv.Atoi(string(val))
		if i%2 == 0 {
			fileSystem2 = append(fileSystem2, []int{id, idx, num}) // id, start index, size
			id++
			idx += num
		} else {
			fileSystem2 = append(fileSystem2, []int{-1, idx, num}) // id, start index, size
			idx += num
		}
	}

	right = len(fileSystem2) - 1

	for right > 0 {
		left = 1
		for left < right {
			if fileSystem2[left][2] >= fileSystem2[right][2] { // check empty space is big enough
				fileSystem2[right][1] = fileSystem2[left][1]  // change start index of file block to start of empty space
				fileSystem2[left][1] += fileSystem2[right][2] // update empty space start index
				fileSystem2[left][2] -= fileSystem2[right][2] // update empty space size
				break
			}
			left += 2
		}
		right -= 2
	}

	checksum = 0

	for i := 0; i < len(fileSystem2); i += 2 {
		start := fileSystem2[i][1]
		size := fileSystem2[i][2]
		for idx := start; idx < start+size; idx++ {
			checksum += fileSystem2[i][0] * idx
		}
	}

	fmt.Println("New file system checksum:", checksum)
}
