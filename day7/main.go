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

	targets := []int{}
	numbers := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		res := strings.Split(line, ": ")

		target, _ := strconv.Atoi(res[0])
		targets = append(targets, target)

		nums := strings.Split(res[1], " ")
		numbers = append(numbers, []int{})
		last := len(numbers) - 1
		for _, num := range nums {
			number, _ := strconv.Atoi(num)
			numbers[last] = append(numbers[last], number)
		}
	}

	calibrationSum := 0

	for i, target := range targets {
		if doesCompute(target, numbers[i][1:], numbers[i][0]) {
			calibrationSum += target
		}
	}

	fmt.Println("Total calibration result:", calibrationSum)

	// part 2

	calibrationSum2 := 0

	for i, target := range targets {
		if doesCompute2(target, numbers[i][1:], numbers[i][0]) {
			calibrationSum2 += target
		}
	}

	fmt.Println("Total second calibration result:", calibrationSum2)
}

func doesCompute(target int, numbers []int, total int) bool {
	if len(numbers) == 0 {
		if target == total {
			return true
		}
		return false
	}

	truth := false
	truth = truth || doesCompute(target, numbers[1:], total+numbers[0])
	truth = truth || doesCompute(target, numbers[1:], total*numbers[0])

	return truth
}

func doesCompute2(target int, numbers []int, total int) bool {
	if len(numbers) == 0 {
		if target == total {
			return true
		}
		return false
	}

	truth := false
	truth = truth || doesCompute2(target, numbers[1:], total+numbers[0])
	truth = truth || doesCompute2(target, numbers[1:], total*numbers[0])
	newTotal, _ := strconv.Atoi(strconv.Itoa(total) + strconv.Itoa(numbers[0]))
	truth = truth || doesCompute2(target, numbers[1:], newTotal)

	return truth
}
