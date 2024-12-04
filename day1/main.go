package main

import (
	"bufio"
	"fmt"
	"math"
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

	list1, list2 := []int{}, []int{}

	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Fields(line)
		num, _ := strconv.Atoi(nums[0])
		list1 = append(list1, num)
		num, _ = strconv.Atoi(nums[1])
		list2 = append(list2, num)
	}

	sort.Ints(list1)
	sort.Ints(list2)

	diff := 0

	for i := 0; i < len(list1); i++ {
		diff += int(math.Abs(float64(list1[i] - list2[i])))
	}

	fmt.Println("Difference in lists:", diff)

	// part 2

	counter := make(map[int]int)
	for _, num := range list2 {
		counter[num]++
	}

	simScore := 0

	for _, num := range list1 {
		if count, ok := counter[num]; ok {
			simScore += count * num
		}
	}

	fmt.Println("Similarity score:", simScore)

	return
}
