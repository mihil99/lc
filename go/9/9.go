// https://leetcode.com/problems/palindrome-number/submissions/1793267485/
package main

import (
	"fmt"
)

var testCases = []map[string]int{
	{
		"x": 121,
	},
	{
		"x": -121,
	},
	{
		"x": 10,
	},
}

func convertToInverse(s int) (int, bool) {
	if s < 0 {
		return 0, false
	}
	t := 0
	for s != 0 {
		t = t*10 + s%10
		s = s / 10
	}
	return t, true
}

func isPalindrome(s int) bool {
	inverse, ok := convertToInverse(s)
	if !ok {
		return false
	}
	return inverse == s
}

func main() {
	for _, testCase := range testCases {
		fmt.Println(isPalindrome(testCase["x"]))
	}
}
