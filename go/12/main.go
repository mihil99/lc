// https://leetcode.com/problems/integer-to-roman/
package main

import (
	"fmt"
	"strings"
)

var testCases = []map[string]int{
	{
		"s": 58, //LVIII
	},
	{
		"s": 3749, //"MMMDCCXLIX"
	},
}

var m = []map[string]any{
	{
		"n": 1000,
		"s": "M",
	},
	{
		"n": 900,
		"s": "CM",
	},
	{
		"n": 500,
		"s": "D",
	},
	{
		"n": 400,
		"s": "CD",
	},
	{
		"n": 100,
		"s": "C",
	},
	{
		"n": 90,
		"s": "XC",
	},
	{
		"n": 50,
		"s": "L",
	},
	{
		"n": 40,
		"s": "XL",
	},
	{
		"n": 10,
		"s": "X",
	},
	{
		"n": 9,
		"s": "IX",
	},
	{
		"n": 5,
		"s": "V",
	},
	{
		"n": 4,
		"s": "IV",
	},
	{
		"n": 1,
		"s": "I",
	},
}

func intToRoman(num int) string {
	res := ""
	for _, obj := range m {
		n := obj["n"].(int)
		if num < n {
			continue
		}
		s := obj["s"].(string)
		d := num / n
		res += strings.Repeat(s, d)
		num -= d * n
	}
	return res
}

func main() {
	for _, testCase := range testCases {
		fmt.Println(intToRoman(testCase["s"]))
	}
}
