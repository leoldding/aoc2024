package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	rules := make(map[string][]string)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		rule := strings.Split(line, "|")
		rules[rule[0]] = append(rules[rule[0]], rule[1])
	}

	updateSum := 0
	incorrectUpdates := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()
		updates := strings.Split(line, ",")
		printed := make(map[string]struct{})
		add := 0
		for _, update := range updates {
			for _, page := range rules[update] {
				if _, ok := printed[page]; ok {
					incorrectUpdates = append(incorrectUpdates, updates)
					goto next
				}
			}
			printed[update] = struct{}{}
		}

		add, _ = strconv.Atoi(updates[len(updates)/2])
		updateSum += add

	next:
	}

	fmt.Println("Sum of update middle page:", updateSum)

	newUpdateSum := 0

	for _, update := range incorrectUpdates {
		sort.Slice(update, func(i, j int) bool {
			for _, page := range rules[update[i]] {
				if page == update[j] {
					return true
				}
			}
			return false
		})
		add, _ := strconv.Atoi(update[len(update)/2])
		newUpdateSum += add
	}

	fmt.Println("Sum of newly update middle page:", newUpdateSum)
}
