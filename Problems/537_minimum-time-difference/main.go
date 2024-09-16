package main

import (
	"fmt"
	"sort"
)

func main() {
	// https://leetcode.com/problems/minimum-time-difference
	fmt.Println("RES:", findMinDifference([]string{"23:59","00:00"})) // 1
	fmt.Println("RES:", findMinDifference([]string{"00:00", "23:59","00:00"})) // 0
}

// Naive Sort Method O(n Log n + n)
func findMinDifference(timePoints []string) int {
	timeMinutes := make([]int, len(timePoints))
	for i, timeStr := range timePoints{
		timeMinutes[i] = timeStringToMinutes(timeStr)
	}

	sort.Ints(timeMinutes)

	minDiff := (timeMinutes[0] + 24 * 60) - timeMinutes[len(timeMinutes) - 1]
	for i := range timeMinutes{
		if i == 0{
			continue
		}
		diff := timeMinutes[i] - timeMinutes[i - 1]
		if minDiff > diff{
			minDiff =  diff
		}
	}
	return minDiff
}


func timeStringToMinutes(str string) int{
	var hour, minutes int
    fmt.Sscanf(str, "%d:%d", &hour, &minutes)
	// fmt.Println(hour, minutes)
	return hour * 60 + minutes
}