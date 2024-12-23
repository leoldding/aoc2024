package main

import (
	"bufio"
	"fmt"
	"math"
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

	registers := make([]int, 3)
	program := []int{}

	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		ret := strings.Split(line, ": ")

		switch count {
		case 3:
		case 4:
			programStrs := strings.Split(ret[1], ",")
			for _, str := range programStrs {
				num, _ := strconv.Atoi(str)
				program = append(program, num)
			}
		default:
			registers[count], _ = strconv.Atoi(ret[1])
		}

		count++
	}

	output := runProgram(program, registers)

	fmt.Println("Program Output:", output[:len(output)-1])

	// part 2

	queue := []int{1}

	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		for i := cur; i < cur+8; i++ {
			addQueue, found := simulate(program, []int{i, 0, 0})
			if addQueue {
				queue = append(queue, i<<3)
			}
			if found {
				fmt.Println("New Register A:", i)
				return
			}
		}
	}

}

func runProgram(program []int, registers []int) string {
	pt := 0
	output := ""
	for pt < len(program) {
		opcode := program[pt]
		literal := program[pt+1]
		combo := literal
		if combo >= 4 && combo < 7 {
			combo = registers[combo%4]
		}

		switch opcode {
		case 0:
			numerator := registers[0]
			denominator := int(math.Pow(2, float64(combo)))
			registers[0] = int(numerator / denominator)
		case 1:
			registers[1] ^= literal
		case 2:
			registers[1] = combo % 8
		case 3:
			if registers[0] != 0 {
				pt = literal
				continue
			}
		case 4:
			registers[1] ^= registers[2]
		case 5:
			output += strconv.Itoa(combo%8) + ","
		case 6:
			numerator := registers[0]
			denominator := int(math.Pow(2, float64(combo)))
			registers[1] = int(numerator / denominator)
		case 7:
			numerator := registers[0]
			denominator := int(math.Pow(2, float64(combo)))
			registers[2] = int(numerator / denominator)
		}

		pt += 2
	}
	return output
}

func simulate(program []int, registers []int) (bool, bool) {
	pt := 0
	temp := registers[0]
	count := 0
	for temp > 0 {
		temp = temp >> 3
		count++
	}
	outPt := len(program) - count
	for pt < len(program) {
		opcode := program[pt]
		literal := program[pt+1]
		combo := literal
		if combo >= 4 && combo < 7 {
			combo = registers[combo%4]
		}

		switch opcode {
		case 0:
			numerator := registers[0]
			denominator := int(math.Pow(2, float64(combo)))
			registers[0] = int(numerator / denominator)
		case 1:
			registers[1] ^= literal
		case 2:
			registers[1] = combo % 8
		case 3:
			if registers[0] != 0 {
				pt = literal
				continue
			}
		case 4:
			registers[1] ^= registers[2]
		case 5:
			cur := combo % 8
			if cur != program[outPt] {
				return false, false
			}
			outPt++
		case 6:
			numerator := registers[0]
			denominator := int(math.Pow(2, float64(combo)))
			registers[1] = int(numerator / denominator)
		case 7:
			numerator := registers[0]
			denominator := int(math.Pow(2, float64(combo)))
			registers[2] = int(numerator / denominator)
		}

		pt += 2
	}
	return true, count == len(program)
}
