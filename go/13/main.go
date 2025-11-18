// https://leetcode.com/problems/roman-to-integer
package main

import (
	"fmt"
)

var testCases = []map[string]string{
	{
		"s": "LVIII", // 58
	},
	{
		"s": "MMMDCCXLIX", // 3749
	},
}

var m = map[rune]int{
	'M': 1000,
	'D': 500,
	'C': 100,
	'L': 50,
	'X': 10,
	'V': 5,
	'I': 1,
}

func romanToInt(roman string) int {
	sum := 0
	for idx, c := range roman {
		if idx+1 < len(roman) && m[c] < m[rune(roman[idx+1])] {
			sum -= m[c]
		} else {
			sum += m[c]
		}
	}
	return sum
}

func main() {
	for _, testCase := range testCases {
		fmt.Println(romanToInt(testCase["s"]))
	}
}
