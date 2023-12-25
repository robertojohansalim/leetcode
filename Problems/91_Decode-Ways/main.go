package main

import (
	"fmt"
)

func main() {
	var res int
	res = numDecodings("1")
	fmt.Println("Output:", res) // 1
	res = numDecodings("11")
	fmt.Println("Output:", res) // 2
	res = numDecodings("111")
	fmt.Println("Output:", res) // 3
	res = numDecodings("1111")
	fmt.Println("Output:", res) // 5
	res = numDecodings("11111")
	fmt.Println("Output:", res) // 8
	res = numDecodings("11110")
	fmt.Println("Output:", res) // 3
}

func numDecodings(s string) int {
	arr := make([]int, len(s))
	for i := range arr {
		arr[i] = -1
	}
	return dp(s, &arr)
}

func dp(s string, arr *[]int) int {
	length := len(s)
	if length <= 0 {
		return 0
	}

	var count int
	// If last Digit Valid
	if s[length-1] != '0' {
		if length-1 > 0 {
			checkIdx := length - 1

			// DP Part
			if (*arr)[checkIdx] == -1 {
				(*arr)[checkIdx] = dp(s[:length-1], arr)
			}

			count += (*arr)[checkIdx]
		} else {
			count++
		}
	}

	// Check last two Digit
	if length-2 >= 0 {
		last2Ch := s[length-2]
		lastCh := s[length-1]

		if last2Ch == '1' || // Allow any number with leading 1 (10 - 19)
			last2Ch == '2' && lastCh <= '6' { // Only allow 20 - 26 (z)
			if length-2 > 0 {
				checkIdx := length - 2

				// DP Part
				if (*arr)[checkIdx] == -1 {
					(*arr)[checkIdx] = dp(s[:length-2], arr)
				}

				count += (*arr)[checkIdx]
			} else {
				count++
			}
		}
	}

	return count
}
