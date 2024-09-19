package main

import (
	"fmt"
)

func main() {
	fmt.Println("Result:", diffWaysToCompute("2-1-1"))
	fmt.Println("Result:", diffWaysToCompute("2*3-4*5"))
	fmt.Println("Result:", diffWaysToCompute("2*33-4*59"))
}

func diffWaysToCompute(expression string) []int {
	var result []int
	arr := splitNumAndOperator(expression)
	// fmt.Println(intArr, opArr)
	result = recursive(arr)

	return result
}

func splitNumAndOperator(expression string) []string {
	var arr []string
	lastDigit := ""
	for _, char := range expression {
		if isOp(string(char)) {
			arr = append(arr, lastDigit, string(char))
			lastDigit = ""
		} else {
			lastDigit = lastDigit + string(char)
		}
	}
	arr = append(arr, lastDigit)
	return arr
}

func recursive(arr []string) []int {
	var result []int

	if len(arr) == 1 {
		return []int{strToInt(arr[0])}
	}

	if len(arr) == 2 {
		return []int{calculate(strToInt(arr[0]), strToInt(arr[1]), arr[0])}
	}

	// var leftResult, rightResult []int
	for i := 0; i < len(arr); i++ {
		if !isOp(arr[i]) {
			continue
		}
		// fmt.Println(arr[:i], ":", string(arr[i]), ":", arr[i+1:])
		op := arr[i]
		leftResult := recursive(arr[:i])
		rightResult := recursive(arr[i+1:])

		for _, left := range leftResult {
			for _, right := range rightResult {
				result = append(result, calculate(left, right, op))
			}
		}
	}

	return result
}

func calculate(a, b int, opStr string) int {
	// func calculate(aStr, bStr, opStr string) int {
	// a := strToInt(aStr)
	// b := strToInt(bStr)
	switch opStr {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	}
	return 0
}

func strToInt(str string) (num int) {
	fmt.Sscanf(str, "%d", &num)
	return num
}

func isOp(char string) bool {
	switch char {
	case "+":
		return true
	case "-":
		return true
	case "*":
		return true
	}
	return false
}
