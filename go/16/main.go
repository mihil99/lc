// https://leetcode.com/problems/3sum-closest
package main

import (
	"fmt"
	"sort"
)

var testCases = []map[string]any{
	{
		"s": []int{-1, 2, 1, -4},
		"t": 1,
	},
	{
		"s": []int{0, 0, 0},
		"t": 1,
	},
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func threeSumClosest(l []int, t int) int {
	sort.Ints(l)
	a := []int{}
	for i, v1 := range l[:len(l)-2] {
		if i > 0 && l[i] == l[i-1] {
			continue
		}
		left := i + 1
		right := len(l) - 1
		for left < right {
			// fmt.Printf("i: %d(%d), left: %d(%d), right: %d(%d), s: %d\n", i, v1, left, l[left], right, l[right], v1+l[left]+l[right])
			s := v1 + l[left] + l[right] - t
			a = append(a, v1+l[left]+l[right])
			if s < 0 {
				left++
			} else if s > 0 {
				right--
			} else {
				return t
			}
		}
	}
	// fmt.Println(a)
	maxDiff := 100000
	best := 100000
	for _, s := range a {
		if abs(t-s) < maxDiff {
			best = s
			maxDiff = abs(t - s)
		}
	}
	return best
}

func main() {
	for _, testCase := range testCases {
		fmt.Println(threeSumClosest(testCase["s"].([]int), testCase["t"].(int)))
	}
}
