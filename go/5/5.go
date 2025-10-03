package main

import (
	"fmt"
)

var testCases = []map[string]string{
	{
		"s": "babad",
	},
	{
		"s": "cbbd",
	},
	{
		"s": "bb",
	},
	{
		"s": "abcba",
	},
}

// func isPalindrome(s string) bool {
// 	for i := 0; i < len(s); i++ {
// 		if s[i] != s[len(s)-i-1] {
// 			return false
// 		}
// 	}
// 	return true
// }

func expandFromCenter(s string, left int, right int) string {
	cur := ""
	for left >= 0 && right < len(s) {
		if s[left] == s[right] {
			cur = s[left : right+1]
			left--
			right++
			continue
		}
		return cur
	}
	return cur
}

func longestPalindrome(s string) string {
	best := ""
	for i := 0; i < len(s); i++ {
		oddPal := expandFromCenter(s, i, i)
		evenPal := expandFromCenter(s, i, i+1)
		if len(oddPal) > len(best) {
			best = oddPal
		}
		if len(evenPal) > len(best) {
			best = evenPal
		}
	}
	return best
}

func main() {
	for _, testCase := range testCases {
		fmt.Println(longestPalindrome(testCase["s"]))
	}
}
