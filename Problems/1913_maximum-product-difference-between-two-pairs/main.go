package main

import "fmt"

func main() {
	var nums []int
	var result int
	nums = []int{5, 6, 2, 7, 4}
	fmt.Println(nums)
	result = maxProductDifference(nums)
	fmt.Println(result)

	nums = []int{4, 2, 5, 9, 7, 4, 8}
	fmt.Println(nums)
	result = maxProductDifference(nums)
	fmt.Println(result)

	nums = []int{2, 9, 5, 9, 1}
	fmt.Println(nums)
	result = maxProductDifference(nums)
	fmt.Println(result)
}
