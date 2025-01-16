package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("error opening file")
		return
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	patterns := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		patterns = strings.Split(line, ", ")
	}

	patternTrie := Trie{&TrieNode{make(map[rune]*TrieNode), false}}

	for _, pattern := range patterns {
		patternTrie.Insert(pattern)
	}

	designs := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		designs = append(designs, line)
	}

	count := 0
	trieMap := make(map[string]bool)

	for _, design := range designs {
		if triePossible(&patternTrie, trieMap, design) {
			count++
		}
	}
	fmt.Println("Number of possible designs:", count)

	// part 2

	count = 0
	trieMapCount := make(map[string]int)
	for _, design := range designs {
		count += trieCount(&patternTrie, trieMapCount, design)
	}

	fmt.Println("Different ways to make designs:", count)
}

func triePossible(t *Trie, trieMap map[string]bool, word string) bool {
	if word == "" {
		return true
	}
	if truth, ok := trieMap[word]; ok {
		return truth
	}

	for i := 0; i <= len(word); i++ {
		left, right := word[:i], word[i:]
		if t.Search(left) {
			if triePossible(t, trieMap, right) {
				trieMap[right] = true
				return true
			} else {
				trieMap[right] = false
			}
		}
	}
	return false
}

func trieCount(t *Trie, trieMap map[string]int, word string) int {
	if word == "" {
		return 1
	}
	if count, ok := trieMap[word]; ok {
		return count
	}

	count := 0
	for i := 0; i <= len(word); i++ {
		left, right := word[:i], word[i:]
		if t.Search(left) {
			count += trieCount(t, trieMap, right)
		}
	}
	trieMap[word] = count
	return count
}

func (t *Trie) Insert(word string) {
	cur := t.root
	for _, ch := range word {
		if _, ok := cur.children[ch]; !ok {
			cur.children[ch] = &TrieNode{make(map[rune]*TrieNode), false}
		}
		cur = cur.children[ch]
	}
	cur.isEnd = true
}

func (t *Trie) Search(word string) bool {
	cur := t.root
	for _, ch := range word {
		if _, ok := cur.children[ch]; !ok {
			return false
		}
		cur = cur.children[ch]
	}
	return cur.isEnd
}
