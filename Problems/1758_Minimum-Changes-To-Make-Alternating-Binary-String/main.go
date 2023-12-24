package main

import "fmt"

func main() {
	var res int
	res = minOperations("0100")
	fmt.Println(res)
	res = minOperations("10")
	fmt.Println(res)
	res = minOperations("1111")
	fmt.Println(res)
}

func minOperations(s string) int {
	var changesForStartsWithZero, changesForStartsWithOne int

	for idx, ch := range s {
		if idx%2 == 0 {
			// Handling Starts with 0
			if ch == '0' {
				changesForStartsWithOne++
			} else {
				changesForStartsWithZero++
			}
		} else {
			if ch != '0' {
				changesForStartsWithOne++
			} else {
				changesForStartsWithZero++
			}
		}
	}

	if changesForStartsWithOne > changesForStartsWithZero {
		return changesForStartsWithZero
	}
	return changesForStartsWithOne
}
