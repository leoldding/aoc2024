package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"
)

type triConn struct {
	conn1 string
	conn2 string
	conn3 string
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("error opening file")
		return
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	connSlice := []string{}
	pairs := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()

		pcs := strings.Split(line, "-")

		if !slices.Contains(connSlice, pcs[0]) {
			connSlice = append(connSlice, pcs[0])
		}
		if !slices.Contains(connSlice, pcs[1]) {
			connSlice = append(connSlice, pcs[1])
		}

		pairs = append(pairs, []string{pcs[0], pcs[1]})
	}

	sort.Strings(connSlice)
	connMat := make([][]bool, len(connSlice))
	for i := 0; i < len(connSlice); i++ {
		connMat[i] = make([]bool, len(connSlice))
	}

	for _, pair := range pairs {
		c1 := slices.Index(connSlice, pair[0])
		c2 := slices.Index(connSlice, pair[1])
		connMat[c1][c2] = true
		connMat[c2][c1] = true
	}

	/*
		fmt.Print("   ")
		for _, conn := range connSlice {
			fmt.Print(conn + " ")
		}
		fmt.Println()
		for i, row := range connMat {
			fmt.Print(connSlice[i] + " ")
			for _, c := range row {
				if c {
					fmt.Print("1  ")
				} else {
					fmt.Print("0  ")
				}
			}
			fmt.Println()
		}
	*/

	triCount := 0
	for i, connRow := range connMat {
		for j := i + 1; j < len(connRow); j++ {
			if connMat[i][j] {
				for k := j + 1; k < len(connRow); k++ {
					if connMat[i][k] && connMat[j][k] {
						if strings.Index(connSlice[i], "t") == 0 || strings.Index(connSlice[j], "t") == 0 || strings.Index(connSlice[k], "t") == 0 {
							triCount++
						}
					}
				}
			}
		}
	}

	fmt.Println("Tri-conns with T:", triCount)

	// part 2

	bigParty := []int{}
	var traverse func(int, []int)
	traverse = func(conn int, party []int) {
		if conn >= len(connMat) {
			return
		}
		for i := conn + 1; i < len(connMat); i++ {
			if connMat[conn][i] {
				for _, partyConn := range party {
					if !connMat[partyConn][i] {
						goto next
					}
				}
				traverse(i, append(party, i))
			next:
				traverse(i, party)
			}
		}
		if len(party) > len(bigParty) {
			bigParty = []int{}
			for _, partyConn := range party {
				bigParty = append(bigParty, partyConn)
			}
		}
	}
	for i, _ := range connMat {
		traverse(i, []int{i})
	}

	password := ""
	for _, conn := range bigParty {
		password += connSlice[conn] + ","
	}

	fmt.Println("LAN party password:", password[:len(password)-1])
}
