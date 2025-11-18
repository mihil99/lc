// https://leetcode.com/problems/longest-common-prefix
package main

import (
	"fmt"
)

var testCases = []map[string][]string{
	{
		"s": []string{"flower", "flow", "flight"},
	},
	{
		"s": []string{"dog", "racecar", "car"},
	},
}

func longestCommonPrefix(l []string) string {
	if len(l) == 0 {
		return ""
	}
	i := 0
	for {
		if len(l[0]) == i {
			return l[0]
		}
		comp := l[0][i]
		for _, s := range l[1:] {
			if len(s) <= i {
				return s
			}
			if s[i] != comp {
				return l[0][:i]
			}
		}
		i++
	}
}

func main() {
	for _, testCase := range testCases {
		fmt.Println(longestCommonPrefix(testCase["s"]))
	}
}
