package main

import (
	"fmt"
	"math"
)

var mapping map[string]int

var modulo = 10

func main() {

	modulo = int(math.Pow10(9) + 7)
	fmt.Println(modulo)
	// numRollsToTarget(1, 6, 3) // 1
	// numRollsToTarget(1, 6, 1) // 1

	fmt.Println()
	fmt.Println()
	numRollsToTarget(2, 6, 2) //
	// numRollsToTarget(2, 6, 12) //
	// numRollsToTarget(2, 6, 11) //
	// numRollsToTarget(2, 6, 13) //

	// fmt.Println()
	// fmt.Println()
	// numRollsToTarget(3, 6, 18)
	// numRollsToTarget(3, 6, 17)
	// numRollsToTarget(3, 6, 16)

	// fmt.Println()
	// fmt.Println()

	// fmt.Println()
	// fmt.Println()
	numRollsToTarget(30, 30, 30)
	numRollsToTarget(30, 30, 900)
	// numRollsToTarget(30, 30, 889)
	numRollsToTarget(30, 30, 500)
}

func numRollsToTarget(n int, k int, target int) int {
	mapping = make(map[string]int)
	res := recursion(n, k, target)
	fmt.Printf("n: %d k: %d target: %d\n", n, k, target)
	fmt.Println("Res:", res)
	return res
}

func recursion(n int, k int, target int) int {
	// fmt.Println("n:", n, " k:", k, " Target:", target)
	if n <= 1 {
		if target > k {
			return 0
		}
		return 1
	}

	count := 0
	for i := 1; i <= k; i++ {
		if target-i < 1 {
			continue
		}
		key := fmt.Sprintf("%d:%d", n-1, target-i)
		var val int
		val, ok := mapping[key]
		if !ok {
			val = recursion(n-1, k, target-i) % modulo
			mapping[key] = val
		}

		count = ((count % modulo) + (val % modulo)) % modulo
	}

	return count
}
