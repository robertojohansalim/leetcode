package main

import (
	"fmt"
)

func main() {
	// https://leetcode.com/problems/container-with-most-water/description/
	fmt.Println("RES:", maxArea([]int{1,8,6,2,5,4,8,3,7})) // 49
}

func maxArea(height []int) int {
	largestVolume := 0
    endIdx := len(height) - 1
	for startIdx:=0 ; startIdx < len(height); startIdx++ { // Iterate per Pivot
        startHeight := height[startIdx]
        endHeight := height[endIdx]
        volume := calculateVolume(startHeight, startIdx, endHeight, endIdx)
        // fmt.Println(startIdx, endIdx , volume)
        if volume > largestVolume {
            largestVolume = volume
        }
		for startIdx < endIdx && startHeight > endHeight{
            endIdx--
			endHeight = height[endIdx]
            // fmt.Println(startIdx, endIdx, volume)
			volume = calculateVolume(startHeight, startIdx, endHeight, endIdx)
			if volume > largestVolume {
				largestVolume = volume
			}
		}
	}

	return largestVolume
}

func calculateVolume(heightA, idxA, heightB, idxB int) int {
    minHeight := heightA
    if minHeight > heightB{
        minHeight = heightB
    }
    volume := minHeight * (idxA - idxB)
    if volume < 0{
        return -volume
    }
    return volume
}