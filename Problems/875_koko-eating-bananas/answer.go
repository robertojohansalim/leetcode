package main

import "fmt"

func minEatingSpeed(piles []int, h int) int {
	max := 0
	for _, p := range piles {
		if p > max {
			max = p
		}
	}

	return binarySearch(piles, max, h)
}

func canBeEaten(piles []int, k, h int) bool {
	fmt.Println(piles, "k:", k, "h:", h)
	pilesIdx := 0
	currPile := piles[0]
	for i := 0; i < h; i++ {
		currPile -= k
		if currPile <= 0 {
			pilesIdx++
			if pilesIdx >= len(piles) {
				return true
			}
			currPile = piles[pilesIdx]
		}

	}
	return false
}

func isSolutionInTheLeftSide(piles []int, k, n int) bool {
	// fmt.Println("K:", k, piles[k])
	isSolveable := canBeEaten(piles, k, n)
	fmt.Println(isSolveable)
	return isSolveable
}

func binarySearch(piles []int, max, h int) int {
	left := 0
	right := max
	for left < right {
		mid := left + (right-left)/2
		if isSolutionInTheLeftSide(piles, mid, h) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
