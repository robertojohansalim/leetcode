package main

import (
	"fmt"
	"math"
	"sort"
)

func maxProductDifference(nums []int) int {

	// return maxProductDifference_sort(nums)
	// return maxProductDifference_TrackSecondMaxMin(nums)
	return maxProductDifference_SecondLargestTracking_editorialAnswer(nums)
}

// I know this is a slower solution of O(n log n)
// But this is a nice refresh on the "sort" package
func maxProductDifference_sort(nums []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	length := len(nums)
	return (nums[0] * nums[1]) - (nums[length-1] * nums[length-2])
}

// The O(n) solution that need to track the seconds max and min value
// This solution actually use O(m * n) where m is the number of nLargest,
// but due to the m being constant of 2 (always fidn second largest) therefore be should be written as O(2n) or simplified to O(n)
func maxProductDifference_TrackSecondMaxMin(nums []int) int {
	firstLargest, secondLargest, firstSmallest, secondSmallest := -1, -1, -1, -1

	firstLargest = nums[0]
	firstLargestIdx := 0
	firstSmallest = nums[0]
	firstSmallestIdx := 0
	for idx, num := range nums {
		if num > firstLargest {
			firstLargest = num
			firstLargestIdx = idx
		}

		if num < firstSmallest {
			firstSmallest = num
			firstSmallestIdx = idx
		}
	}

	secondLargest = firstSmallest
	secondSmallest = firstLargest
	for idx, num := range nums {
		if idx == firstLargestIdx || idx == firstSmallestIdx {
			continue
		}

		if num > secondLargest {
			secondLargest = num
		}

		if num < secondSmallest {
			secondSmallest = num
		}
	}

	fmt.Printf("(%d * %d) - (%d * %d)\n", firstLargest, secondLargest, firstSmallest, secondSmallest)

	return (firstLargest * secondLargest) - (firstSmallest * secondSmallest)
}

func maxProductDifference_SecondLargestTracking_editorialAnswer(nums []int) int {
	firstLargest, secondLargest := -math.MaxInt, -math.MaxInt
	firstSmallest, secondSmallest := math.MaxInt, math.MaxInt

	for _, num := range nums {
		// If largest is changed, transfer current largest to second largest
		if num > firstLargest {
			secondLargest = firstLargest
			firstLargest = num
		} else {
			// Only perform check on second Largest when not assigning the first Largest
			if num > secondLargest {
				secondLargest = num
			}
		}

		if num < firstSmallest {
			secondSmallest = firstSmallest
			firstSmallest = num
		} else {
			if num < secondSmallest {
				secondSmallest = num
			}
		}
	}

	fmt.Printf("(%d * %d) - (%d * %d)\n", firstLargest, secondLargest, firstSmallest, secondSmallest)

	return (firstLargest * secondLargest) - (firstSmallest * secondSmallest)
}
