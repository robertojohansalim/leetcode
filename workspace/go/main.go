package main

import (
	"fmt"
)

func main() {
	fmt.Println("Result:", lexicalOrder(13)) // [0,2]

	fmt.Println("Result:", lexicalOrder(2)) // [-34,-14,-10,-10,10]
	fmt.Println("Result:", lexicalOrder(100))
}

func lexicalOrder(n int) []int {
	arr := make([]int, n)

	idx := 0
	for i := 1; i < 10; i++ {
		if i > n || idx >= n {
			break
		}
		arr[idx] = i
		idx++
		// Solve Inner Lexico With head of i
		recursion(i*10, n, &idx, arr)
	}
	return arr
}

func recursion(lead, n int, idx *int, arr []int) {
	if lead > n {
		return
	}

	for i := 0; i < 10; i++ {
		curr := lead + i
		if curr > n || *idx >= n {
			return
		}
		// fmt.Println()
		// fmt.Println(arr)
		(arr[*idx]) = curr
		*idx++
		recursion(curr*10, n, idx, arr)
	}
}
