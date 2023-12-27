package main

import "fmt"

func main() {
	solve("abaac", []int{1, 2, 3, 4, 5})
	solve("abc", []int{1, 2, 3})
	solve("aabaa", []int{1, 2, 3, 4, 1})
}

func solve(colors string, neededTime []int) {
	res := minCost(colors, neededTime)
	fmt.Println(colors, neededTime)
	fmt.Println("Result:", res)
}

// PSEUDO CODE:
// while repeating letter, keep track of largestMinimum
// if current neededTime is less than largestMinimum
//
//	add neededTime
//
// else
//
//	add largstMinimum, and update largestMinimum to current neededTime value
func minCost(colors string, neededTime []int) int {
	totalCost := 0

	var runningColor rune
	var largestMinimum int

	for i, c := range colors {
		time := neededTime[i]
		if runningColor == c {
			removedTime := time
			if time < largestMinimum {
				removedTime = time
			} else {
				removedTime = largestMinimum
				largestMinimum = time
			}
			totalCost += removedTime
		} else {
			runningColor = c
			largestMinimum = time
		}
	}

	return totalCost
}
