package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Connection struct {
	wire1 string
	wire2 string
	gate  string
	out   string
	ind   int
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("error opening file")
		return
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	wires := make(map[string]int)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		split := strings.Split(line, ": ")
		val, _ := strconv.Atoi(split[1])
		wires[split[0]] = val
	}

	conns := []Connection{}
	zWires := []string{}
	ind := 0
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		conns = append(conns, Connection{split[0], split[2], split[1], split[4], ind})
		if strings.Index(split[4], "z") == 0 {
			zWires = append(zWires, split[4])
		}
		ind++
	}
	sort.Strings(zWires)

	fmt.Println("Z wires output:", compute(conns, wires, zWires))

	// part 2

	findConn := func(wire string) (Connection, error) {
		for _, conn := range conns {
			if conn.out == wire {
				return conn, nil
			}
		}
		return Connection{}, errors.New("not found")
	}

	reverseConns := make(map[string][][]Connection)
	for _, wire := range zWires {
		zConn, _ := findConn(wire)
		reverseConns[wire] = [][]Connection{[]Connection{zConn}}
		for {
			insert := []Connection{}
			for _, revConn := range reverseConns[wire][len(reverseConns[wire])-1] {
				conn, err := findConn(revConn.wire1)
				if err == nil {
					insert = append(insert, conn)
				}

				conn, err = findConn(revConn.wire2)
				if err == nil {
					insert = append(insert, conn)
				}
			}

			if len(insert) == 0 {
				break
			}
			reverseConns[wire] = append(reverseConns[wire], insert)
		}
	}

	simulate := func() ([46]int, bool) {
		cycles := 10
		complete := 0
		wrong := [46]int{}
		for i := 0; i < cycles; i++ {
			rand1 := rand.Intn(int(math.Pow(2.0, 44)))
			rand2 := rand.Intn(int(math.Pow(2.0, 44)))
			sum := rand1 + rand2
			//fmt.Printf("%046b is: %d\n", rand1, rand1)
			//fmt.Printf("%046b is: %d\n", rand2, rand2)
			//fmt.Printf("%046b is: %d\n", sum, sum)

			ind := 0
			for rand1 > 0 {
				wires["x"+fmt.Sprintf("%02d", ind)] = rand1 & 1
				rand1 >>= 1
				ind++
			}

			ind = 0
			for rand2 > 0 {
				wires["y"+fmt.Sprintf("%02d", ind)] = rand2 & 1
				rand2 >>= 1
				ind++
			}

			test := compute(conns, wires, zWires)
			if test == -1 {
				continue
			}
			//fmt.Printf("%046b is: %d\n", test, test)

			ind = 0
			for sum > 0 {
				if sum&1 != test&1 {
					wrong[ind]++
					break
				}
				sum >>= 1
				test >>= 1
				ind++
			}
			complete++
		}

		return wrong, cycles == complete
	}

	wrong, complete := simulate()
	if !complete {
		return
	}
	fmt.Println(wrong)
	i0, _ := findMax(wrong)

	/*
		used XOR rules to find and validate already found values
		https://www.reddit.com/r/adventofcode/comments/1hl698z/comment/m3k68gd/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button
		mnm XOR gqb -> gwh
		wjf XOR ksf -> jct
		kgk XOR sbs -> rcb
		kgk AND sbs -> z21
		x39 AND y39 -> z39
		dsk OR ptc -> z09
	*/

	// {dsk ptc OR z09 149} {mnm gqb XOR gwh 83}
	fmt.Println(conns[83], conns[161])
	conns[149].out, conns[83].out = conns[83].out, conns[149].out
	i0 = 13

	//{y13 x13 AND wbw 111} {y13 x13 XOR wgb 191}
	fmt.Println(conns[111], conns[191])
	conns[111].out, conns[191].out = conns[191].out, conns[111].out
	i0 = 21

	//{kgk sbs AND z21 68} {kgk sbs XOR rcb 203}
	fmt.Println(conns[68], conns[203])
	conns[68].out, conns[203].out = conns[203].out, conns[68].out
	i0 = 39

	/*
		// x39 AND y39 -> z39 wjf XOR ksf -> jct
		conns[141].out, conns[161].out = conns[161].out, conns[141].out


		SOLUTION: gwh,jct,rcb,wbw,wgb,z09,z21,z39
	*/

	fmt.Println(i0)
	restrictedConns := []Connection{}
	for i := 0; i < i0; i++ {
		for _, revConns := range reverseConns["z"+fmt.Sprintf("%02d", i)] {
			for _, conn := range revConns {
				restrictedConns = append(restrictedConns, conn)
			}
		}
	}

	swapConns := []Connection{}
	for _, revConns := range reverseConns["z"+fmt.Sprintf("%02d", i0)] {
		for _, conn := range revConns {
			swapConns = append(swapConns, conn)
		}
	}

	for _, conn1 := range swapConns {
		for _, conn2 := range conns {
			if slices.Contains(restrictedConns, conn1) || slices.Contains(restrictedConns, conn2) {
				continue
			}

			conns[conn1.ind].out, conns[conn2.ind].out = conns[conn2.ind].out, conns[conn1.ind].out
			wrong, complete = simulate()
			conns[conn1.ind].out, conns[conn2.ind].out = conns[conn2.ind].out, conns[conn1.ind].out
			if complete {
				i, _ := findMax(wrong)
				if wrong[i0] == 0 && i > i0 {
					fmt.Println(conns[conn1.ind], conns[conn2.ind])
					fmt.Println(i, wrong)
				}
			}
		}
	}

}

func findMax(arr [46]int) (int, int) {
	largest, secondLargest := math.MinInt32, math.MinInt32
	i0, i1 := -1, -1

	for i, v := range arr {
		if v > largest {
			secondLargest = largest
			i1 = i0
			largest = v
			i0 = i
		} else if v > secondLargest && v != largest {
			secondLargest = v
			i1 = i
		}
	}

	return i0, i1
}

func compute(baseConns []Connection, baseWires map[string]int, zWires []string) int {
	conns := []Connection{}
	for _, v := range baseConns {
		conns = append(conns, v)
	}

	wires := make(map[string]int)
	for k, v := range baseWires {
		wires[k] = v
	}

	for len(conns) > 0 {
		computed := 0

		for i := len(conns) - 1; i >= 0; i-- {
			if _, ok := wires[conns[i].wire1]; !ok {
				continue
			}
			if _, ok := wires[conns[i].wire2]; !ok {
				continue
			}
			switch conns[i].gate {
			case "AND":
				wires[conns[i].out] = wires[conns[i].wire1] & wires[conns[i].wire2]
			case "OR":
				wires[conns[i].out] = wires[conns[i].wire1] | wires[conns[i].wire2]
			case "XOR":
				wires[conns[i].out] = wires[conns[i].wire1] ^ wires[conns[i].wire2]
			}

			conns[i], conns[len(conns)-1-computed] = conns[len(conns)-1-computed], conns[i]
			computed++
		}
		if computed == 0 {
			return -1
		}

		conns = conns[:len(conns)-computed]
	}
	bits := 0
	count := 0

	for _, wire := range zWires {
		temp := wires[wire] << count
		bits |= temp
		count++
	}

	return bits
}
