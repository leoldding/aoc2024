package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/mat"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("error opening file")
		return
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	aButtons := [][]int{}
	bButtons := [][]int{}
	prizes := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		split := strings.Split(line, ":")
		instruction, dist := split[0], split[1]

		switch instruction {
		case "Button A":
			split = strings.Split(dist, ",")
			tempX := strings.Split(split[0], "+")[1]
			x, _ := strconv.Atoi(tempX)
			tempY := strings.Split(split[1], "+")[1]
			y, _ := strconv.Atoi(tempY)
			aButtons = append(aButtons, []int{x, y})
		case "Button B":
			split = strings.Split(dist, ",")
			tempX := strings.Split(split[0], "+")[1]
			x, _ := strconv.Atoi(tempX)
			tempY := strings.Split(split[1], "+")[1]
			y, _ := strconv.Atoi(tempY)
			bButtons = append(bButtons, []int{x, y})
		case "Prize":
			split = strings.Split(dist, ",")
			tempX := strings.Split(split[0], "=")[1]
			x, _ := strconv.Atoi(tempX)
			tempY := strings.Split(split[1], "=")[1]
			y, _ := strconv.Atoi(tempY)
			prizes = append(prizes, []int{x, y})
		}
	}

	tokens := 0

	for i, _ := range prizes {
		token := playMachine(0, 0, 0, 0, 100, aButtons[i], bButtons[i], prizes[i], make(map[int]map[int]int))
		if token == math.MaxInt32 {
			continue
		}
		tokens += token
	}

	fmt.Println("Tokens spent:", tokens)

	// part 2

	newPrizes := [][]int{}
	for _, prize := range prizes {
		newPrizes = append(newPrizes, []int{prize[0] + 10000000000000, prize[1] + 10000000000000})
	}

	tokens = 0

	for i, _ := range newPrizes {
		A := mat.NewDense(2, 2, []float64{float64(aButtons[i][0]), float64(bButtons[i][0]), float64(aButtons[i][1]), float64(bButtons[i][1])})
		b := mat.NewVecDense(2, []float64{float64(newPrizes[i][0]), float64(newPrizes[i][1])})
		var x mat.VecDense
		if err := x.SolveVec(A, b); err != nil {
			fmt.Println(err)
			return
		}

		if math.Abs(x.At(0, 0)-math.Round(x.At(0, 0))) < 0.01 && math.Abs(x.At(1, 0)-math.Round(x.At(1, 0))) < 0.01 {
			tokens += int(math.Round(x.At(0, 0))*3 + math.Round(x.At(1, 0)))
		}

	}
	fmt.Println("New tokens spent:", tokens)
}

func playMachine(x int, y int, aCount int, bCount int, countCap int, aButton []int, bButton []int, prize []int, memo map[int]map[int]int) int {
	if x == prize[0] && y == prize[1] {
		return aCount*3 + bCount
	}
	if res, ok := memo[aCount][bCount]; ok {
		return res
	}
	if countCap != -1 {
		if aCount > countCap || bCount > countCap {
			return math.MaxInt32
		}
	}
	if x > prize[0] || y > prize[1] {
		return math.MaxInt32
	}
	pressA := playMachine(x+aButton[0], y+aButton[1], aCount+1, bCount, countCap, aButton, bButton, prize, memo)
	pressB := playMachine(x+bButton[0], y+bButton[1], aCount, bCount+1, countCap, aButton, bButton, prize, memo)
	if _, ok := memo[aCount]; !ok {
		memo[aCount] = make(map[int]int)
	}
	if pressA < pressB {
		memo[aCount][bCount] = pressA
		return pressA
	}
	memo[aCount][bCount] = pressB
	return pressB
}
