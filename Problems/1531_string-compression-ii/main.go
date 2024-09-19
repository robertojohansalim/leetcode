package main

import (
	"fmt"
	"sort"
)

func main() {
	evaluate("aaabcccd", 2)
	evaluate("aabbaa", 2)
	evaluate("aaaaaaaaaaa", 2)
	evaluate("aaaaaaaaaaa", 0)
	evaluate("aaaaaaaaaaa", 10)
}

type pair struct {
	code  rune
	count int
}

func evaluate(s string, k int) int {
	res := getLengthOfOptimalCompression(s, k)
	// fmt.Println("Result", res)
	return res
}

type costValue struct {
	cost  int
	value int
}

func getLengthOfOptimalCompression(s string, k int) int {
	var pairs []pair

	lastPairIdx := -1
	var lastPairCode rune
	for _, char := range s {
		if lastPairCode == char {
			pairs[lastPairIdx].count++
		} else {
			pairs = append(pairs, pair{char, 1})
			lastPairCode = char
			lastPairIdx++
		}
		// fmt.Println(pairs)
	}

	// construct encoded:
	finalString := ""
	var neededToReduceLength [][]costValue
	for _, p := range pairs {

		var str string
		if p.count > 1 {
			str = fmt.Sprintf("%c%d", p.code, p.count)
		} else {
			str = string(p.code)
		}
		finalString += str
		var costs []costValue
		if p.count == 1 {
			costs = append(costs, costValue{cost: 1, value: 1})
		} else if p.count > 1 && p.count < 10 {
			costs = append(costs, costValue{cost: p.count - 1, value: 1})
		} else if p.count >= 10 {
			costs = append(costs, costValue{cost: p.count - 1 - (p.count % 10), value: 1})
		}

		neededToReduceLength = append(neededToReduceLength, costs)
	}
	fmt.Println(neededToReduceLength)

	sort.Slice(neededToReduceLength, func(i, j int) bool {
		return neededToReduceLength[i].cost < neededToReduceLength[j].cost
	})
	for _, val := range neededToReduceLength {
		fmt.Print(val.cost, " ")
	}
	fmt.Println()

	savedLength := 0
	kTmp := k
	for _, val := range neededToReduceLength {
		if kTmp < val.cost {
			break
		}
		kTmp -= val.cost
		savedLength += val.value
	}
	fmt.Println("S:", s, "K:", k, "encode:", finalString)
	fmt.Println("From", len(finalString), "To", len(finalString)-savedLength)
	return len(finalString) - savedLength
}
