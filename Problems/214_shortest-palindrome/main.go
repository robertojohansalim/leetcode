package main

import "fmt"

func main(){
	fmt.Println("RES:", shortestPalindrome("aacecaaa")) // aaacecaaa
	fmt.Println("RES:", shortestPalindrome("abcd")) // dcbabcd
	fmt.Println("RES:", shortestPalindrome("aa")) // aa
	fmt.Println("RES:", shortestPalindrome("a")) // a
}

func shortestPalindrome(s string) string {
    // reverseStartIdx := len(s) - 1
    reverse := Reverse(s)

    for pivotReverseIdx := range reverse{
        rIdx := pivotReverseIdx
        // fmt.Println("Move Pivot", reverse[rIdx:])
        for i := range s{
            if rIdx >= len(reverse) -1{
                return reverse[:pivotReverseIdx] + s
            }
            // fmt.Println("Compare:", string(s[i]), string(reverse[rIdx]))
            if s[i] != reverse[rIdx]{
                break
            }
            rIdx++
        }
    }

    return ""
}


func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}