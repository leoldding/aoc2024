package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("error opening file")
		return
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	complex2 := 0
	complex25 := 0

	for scanner.Scan() {
		code := scanner.Text()

		num, _ := strconv.Atoi(code[:len(code)-1])

		shortest2 := math.MaxInt
		shortest25 := math.MaxInt

		shortest2 = min(shortest2, dirToDir(code, 3))
		shortest25 = min(shortest25, dirToDir(code, 26))

		complex2 += num * shortest2
		complex25 += num * shortest25
	}

	fmt.Println("Total complexity (depth 2):", complex2)
	fmt.Println("Total complexity (depth 25):", complex25)
}

type codeSeq struct {
	code  string
	depth int
}

var memo = make(map[codeSeq]int)

type pair struct {
	first  rune
	second rune
}

func dirToDir(code string, depth int) int {
	seq := codeSeq{code, depth}
	if val, ok := memo[seq]; ok {
		return val
	}

	length := 0
	if depth == 0 {
		length = len(code)
		memo[seq] = length
		return length
	}
	cur := 'A'
	for _, next := range code {
		if cur == next {
			length += 1
		} else {
			newSeq := paths[pair{cur, next}]
			length += dirToDir(newSeq, depth-1)
		}
		cur = next
	}
	memo[seq] = length
	return length
}

var paths = map[pair]string{
	{'A', '0'}: "<A",
	{'0', 'A'}: ">A",
	{'A', '1'}: "^<<A",
	{'1', 'A'}: ">>vA",
	{'A', '2'}: "<^A",
	{'2', 'A'}: "v>A",
	{'A', '3'}: "^A",
	{'3', 'A'}: "vA",
	{'A', '4'}: "^^<<A",
	{'4', 'A'}: ">>vvA",
	{'A', '5'}: "<^^A",
	{'5', 'A'}: "vv>A",
	{'A', '6'}: "^^A",
	{'6', 'A'}: "vvA",
	{'A', '7'}: "^^^<<A",
	{'7', 'A'}: ">>vvvA",
	{'A', '8'}: "<^^^A",
	{'8', 'A'}: "vvv>A",
	{'A', '9'}: "^^^A",
	{'9', 'A'}: "vvvA",
	{'0', '1'}: "^<A",
	{'1', '0'}: ">vA",
	{'0', '2'}: "^A",
	{'2', '0'}: "vA",
	{'0', '3'}: "^>A",
	{'3', '0'}: "<vA",
	{'0', '4'}: "^<^A",
	{'4', '0'}: ">vvA",
	{'0', '5'}: "^^A",
	{'5', '0'}: "vvA",
	{'0', '6'}: "^^>A",
	{'6', '0'}: "<vvA",
	{'0', '7'}: "^^^<A",
	{'7', '0'}: ">vvvA",
	{'0', '8'}: "^^^A",
	{'8', '0'}: "vvvA",
	{'0', '9'}: "^^^>A",
	{'9', '0'}: "<vvvA",
	{'1', '2'}: ">A",
	{'2', '1'}: "<A",
	{'1', '3'}: ">>A",
	{'3', '1'}: "<<A",
	{'1', '4'}: "^A",
	{'4', '1'}: "vA",
	{'1', '5'}: "^>A",
	{'5', '1'}: "<vA",
	{'1', '6'}: "^>>A",
	{'6', '1'}: "<<vA",
	{'1', '7'}: "^^A",
	{'7', '1'}: "vvA",
	{'1', '8'}: "^^>A",
	{'8', '1'}: "<vvA",
	{'1', '9'}: "^^>>A",
	{'9', '1'}: "<<vvA",
	{'2', '3'}: ">A",
	{'3', '2'}: "<A",
	{'2', '4'}: "<^A",
	{'4', '2'}: "v>A",
	{'2', '5'}: "^A",
	{'5', '2'}: "vA",
	{'2', '6'}: "^>A",
	{'6', '2'}: "<vA",
	{'2', '7'}: "<^^A",
	{'7', '2'}: "vv>A",
	{'2', '8'}: "^^A",
	{'8', '2'}: "vvA",
	{'2', '9'}: "^^>A",
	{'9', '2'}: "<vvA",
	{'3', '4'}: "<<^A",
	{'4', '3'}: "v>>A",
	{'3', '5'}: "<^A",
	{'5', '3'}: "v>A",
	{'3', '6'}: "^A",
	{'6', '3'}: "vA",
	{'3', '7'}: "<<^^A",
	{'7', '3'}: "vv>>A",
	{'3', '8'}: "<^^A",
	{'8', '3'}: "vv>A",
	{'3', '9'}: "^^A",
	{'9', '3'}: "vvA",
	{'4', '5'}: ">A",
	{'5', '4'}: "<A",
	{'4', '6'}: ">>A",
	{'6', '4'}: "<<A",
	{'4', '7'}: "^A",
	{'7', '4'}: "vA",
	{'4', '8'}: "^>A",
	{'8', '4'}: "<vA",
	{'4', '9'}: "^>>A",
	{'9', '4'}: "<<vA",
	{'5', '6'}: ">A",
	{'6', '5'}: "<A",
	{'5', '7'}: "<^A",
	{'7', '5'}: "v>A",
	{'5', '8'}: "^A",
	{'8', '5'}: "vA",
	{'5', '9'}: "^>A",
	{'9', '5'}: "<vA",
	{'6', '7'}: "<<^A",
	{'7', '6'}: "v>>A",
	{'6', '8'}: "<^A",
	{'8', '6'}: "v>A",
	{'6', '9'}: "^A",
	{'9', '6'}: "vA",
	{'7', '8'}: ">A",
	{'8', '7'}: "<A",
	{'7', '9'}: ">>A",
	{'9', '7'}: "<<A",
	{'8', '9'}: ">A",
	{'9', '8'}: "<A",
	{'<', '^'}: ">^A",
	{'^', '<'}: "v<A",
	{'<', 'v'}: ">A",
	{'v', '<'}: "<A",
	{'<', '>'}: ">>A",
	{'>', '<'}: "<<A",
	{'<', 'A'}: ">>^A",
	{'A', '<'}: "v<<A",
	{'^', 'v'}: "vA",
	{'v', '^'}: "^A",
	{'^', '>'}: "v>A",
	{'>', '^'}: "<^A",
	{'^', 'A'}: ">A",
	{'A', '^'}: "<A",
	{'v', '>'}: ">A",
	{'>', 'v'}: "<A",
	{'v', 'A'}: "^>A",
	{'A', 'v'}: "<vA",
	{'>', 'A'}: "^A",
	{'A', '>'}: "vA",
}
