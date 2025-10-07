// https://leetcode.com/problems/regular-expression-matching/
package main

import (
	"fmt"
)

var testCases = []map[string]string{
	{
		"s": "aa",
		"p": "a",
	},
	{
		"s": "a",
		"p": "a*",
	},
	{
		"s": "ab",
		"p": ".*",
	},
}

var m = map[string]bool{}

func isMatch(s string, p string) bool {
	if p == "" {
		return s == ""
	}
	if cacheVal, ok := m[s+"|"+p]; ok {
		return cacheVal

	}
	firstMatch := s != "" && (p[0] == s[0] || p[0] == '.')
	if len(p) >= 2 && p[1] == '*' {
		result := isMatch(s, p[2:]) || (firstMatch && isMatch(s[1:], p))
		m[s+"|"+p] = result
		return result
	}
	result := firstMatch && isMatch(s[1:], p[1:])
	m[s+"|"+p] = result
	return result
}

func main() {
	for _, testCase := range testCases {
		fmt.Println(isMatch(testCase["s"], testCase["p"]))
	}
}
