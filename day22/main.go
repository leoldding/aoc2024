package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type seq struct {
	first  int
	second int
	third  int
	fourth int
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("error opening file")
		return
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	secrets := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		secret, _ := strconv.Atoi(line)
		secrets = append(secrets, secret)
	}

	secretSum := 0
	for _, secret := range secrets {
		for i := 0; i < 2000; i++ {
			secret = simulate(secret)
		}
		secretSum += secret
	}
	fmt.Println("Secret sum:", secretSum)

	// part 2

	seqMap := make(map[seq]int)

	first, second, third, fourth := 0, 0, 0, 0
	for _, secret := range secrets {
		prev := secret % 10

		secret = simulate(secret)
		first = secret%10 - prev
		prev = secret % 10

		secret = simulate(secret)
		second = secret%10 - prev
		prev = secret % 10

		secret = simulate(secret)
		third = secret%10 - prev
		prev = secret % 10

		secret = simulate(secret)
		fourth = secret%10 - prev
		prev = secret % 10

		sequence := seq{first, second, third, fourth}
		seqMap[sequence] += secret % 10
		seqUsed := make(map[seq]struct{})
		seqUsed[sequence] = struct{}{}

		for i := 4; i < 2000; i++ {
			first = second
			second = third
			third = fourth
			secret = simulate(secret)
			fourth = secret%10 - prev
			prev = secret % 10

			sequence := seq{first, second, third, fourth}
			if _, ok := seqUsed[sequence]; !ok {
				seqMap[sequence] += secret % 10
				seqUsed[sequence] = struct{}{}
			}
		}
	}

	maxBananas := 0
	for _, bananas := range seqMap {
		maxBananas = max(maxBananas, bananas)
	}

	fmt.Println("Max bananas:", maxBananas)
}

func simulate(num int) int {
	num = prune(mix(num*64, num))
	num = prune(mix(int(num/32), num))
	num = prune(mix(num*2048, num))
	return num
}

func simulate2(num int, seqMap map[seq]int) {

}

func mix(num1 int, num2 int) int {
	return num1 ^ num2
}

func prune(num int) int {
	return num % 16777216
}
