package main

import (
	"fmt"
)

func main() {
	solve([]string{"ab", "a"})
	solve([]string{"abb", "aab"})
	solve([]string{"aa", "bb"})
	solve([]string{"a", "b"})
	solve([]string{"a", "b", "c"})
	solve([]string{"asdf"})
}

func solve(words []string) {
	result := makeEqual(words)
	fmt.Println(result)
}

func makeEqual(words []string) bool {
	countMap := make(map[rune]int)
	for _, word := range words {
		for _, char := range word {
			_, ok := countMap[char]
			if !ok {
				countMap[char] = 1
			} else {
				countMap[char]++
			}
		}
	}

	minimumCount := len(words)
	if len(words) == 1 {
		return true
	}

	for _, count := range countMap {
		if count%minimumCount != 0 {
			return false
		}
	}
	return true
}
