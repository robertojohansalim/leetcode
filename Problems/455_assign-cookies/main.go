package main

import (
	"fmt"
	"sort"
)

func main() {
	var res int
	res = findContentChildren([]int{1, 2, 3}, []int{1, 1})
	fmt.Println("Result:", res)
	res = findContentChildren([]int{1, 2}, []int{2, 3})
	fmt.Println("Result:", res)
	res = findContentChildren([]int{10, 9, 8, 7}, []int{5, 6, 7, 8})
	fmt.Println("Result:", res)
	res = findContentChildren([]int{1, 2, 3}, []int{3})
	fmt.Println("Result:", res)
}

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	fmt.Println(s)
	fmt.Println(g)

	i := 0
	satisfied := 0
	for _, greed := range g {
		for i < len(s) && s[i] < greed {
			i++
		}
		if i < len(s) {
			satisfied++
			i++ // If valid then nexted cookie (current is used)
		} else {
			break
		}
	}
	return satisfied
}
