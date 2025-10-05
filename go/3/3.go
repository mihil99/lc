// https://leetcode.com/problems/longest-substring-without-repeating-characters/
package main

import (
	"fmt"
)

// Double for loop
// func lengthOfLongestSubstring(s string) int {
// 	if l := len(s); l <= 1 {
// 		return l
// 	}
// 	r := 1
// 	for i, c1 := range s {
// 		m := make(map[rune]bool)
// 		m[c1] = true
// 		for j, c2 := range s[i+1:] {
// 			if _, ok := m[c2]; ok {
// 				r = max(r, j+1)
// 				break
// 			}
// 			m[c2] = true
// 		}
// 		r = max(r, len(m))
// 	}
// 	return r
// }

func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int)
	l := 0
	start := 0

	for end, char := range s {
		if idx, ok := m[char]; ok {
			if idx >= start {
				start = idx + 1
			}
		}
		m[char] = end
		l = max(end-start+1, l)
	}
	return l
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("au"))
}
