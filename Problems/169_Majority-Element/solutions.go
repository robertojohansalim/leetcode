package main

func majorityElement(nums []int) int {
	var candidate, vote int
	for _, num := range nums {
		if vote == 0 {
			candidate = num
		}

		if candidate == num {
			vote++
		} else {
			vote--
		}

	}
	return candidate
}
