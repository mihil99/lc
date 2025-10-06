// https://leetcode.com/problems/string-to-integer-atoi/
package main

import (
	"fmt"
	"math"
	"strings"
)

var testCases = []map[string]string{
	{
		"s": "42",
	},
	{
		"s": "-042",
	},
	{
		"s": "1337c0d3",
	},
	{
		"s": "0-1",
	},
	{
		"s": "words and 987",
	},
}

func clamp(n int) int {
	MIN := int(-math.Pow(2, 31))
	if n < MIN {
		return MIN
	}
	MAX := int(math.Pow(2, 31) - 1)
	if n > MAX {
		return MAX
	}
	return n
}

func myAtoi(s string) int {
	INTS := "0123456789"
	SIGNS := "-+"
	WHITESPACE := "\n\t\r "
	VALID := INTS + SIGNS + WHITESPACE
	t := 0
	sign := 1
	runes := []rune(s)
	// whitespace
	for idx, r := range runes {
		// fmt.Printf("%s: handling %c in whitespace step\n", s, r)
		if !strings.ContainsRune(VALID, r) {
			// fmt.Printf("%s: invalid in whitespace step\n", s)
			return 0
		}
		if !strings.ContainsRune(WHITESPACE, r) {
			// fmt.Printf("%s: no whitespace char (%c) during whitespace step", s, r)
			runes = runes[idx:]
			break
		}
		// fmt.Printf("%s: whitespace char (%c) during whitespace step", s, r)
	}
	// sign
	if len(runes) > 0 && strings.ContainsRune(SIGNS, runes[0]) {
		if len(runes) == 1 {
			return 0
		}
		if runes[0] == '-' {
			sign = -1
		}
		runes = runes[1:]
	}
	// numbers
	for _, r := range runes {
		if !strings.ContainsRune(INTS, r) {
			return clamp(t)
		}
		if r == '0' && t == 0 {
			continue
		}
		for idx, sn := range INTS {
			if r == sn {
				t = clamp(t*10 + sign*idx)
				break
			}
		}
	}
	return clamp(t * sign)
}

func main() {
	for _, testCase := range testCases {
		fmt.Println(myAtoi(testCase["s"]))
	}
}
