package binarysearch

func BinarySearch(max int) int {
	left := 0
	right := max

	for left < right {
		mid := left + (right-left)/2

		if isSolutionAtTheLeftOfMid(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func isSolutionAtTheLeftOfMid(mid int) bool {
	return true
}
