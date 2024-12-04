package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	memory := ""

	for scanner.Scan() {
		line := scanner.Text()
		memory += line
	}

	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := re.FindAllString(memory, -1)

	mult := 0

	for _, match := range matches {
		split := strings.Index(match, ",")
		num1, _ := strconv.Atoi(match[4:split])
		num2, _ := strconv.Atoi(match[split+1 : len(match)-1])
		mult += num1 * num2
	}

	fmt.Println("Memory multiplication:", mult)

	// part 2

	re = regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	matches = re.FindAllString(memory, -1)
	do := true
	condMult := 0

	for _, match := range matches {
		switch match {
		case "do()":
			do = true
		case "don't()":
			do = false
		default:
			if !do {
				continue
			}
			split := strings.Index(match, ",")
			num1, _ := strconv.Atoi(match[4:split])
			num2, _ := strconv.Atoi(match[split+1 : len(match)-1])
			condMult += num1 * num2
		}
	}

	fmt.Println("Memory conditional multiplication:", condMult)
}
