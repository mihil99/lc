package main

import (
	"fmt"
	"sort"
)

var testCases = []map[string][]int{
	{
		"s": []int{-1, 0, 1, 2, -1, -4},
	},
	{
		"s": []int{0, 1, 1},
	},
	{
		"s": []int{0, 0, 0},
	},
}

func threeSum(l []int) [][]int {
	sort.Ints(l)
	if len(l) < 3 {
		return [][]int{}
	}
	result := [][]int{}
	for i, v1 := range l[:len(l)-2] {
		if i > 0 && l[i] == l[i-1] {
			continue
		}
		left := i + 1
		right := len(l) - 1
		for left < right {
			s := v1 + l[left] + l[right]
			if s < 0 {
				left++
			} else if s > 0 {
				right--
			} else {
				result = append(result, []int{v1, l[left], l[right]})
				left++
				right--
				for left < right && l[left] == l[left-1] {
					left++
				}
			}
		}
	}
	return result
}

func main() {
	for _, testCase := range testCases {
		fmt.Println(threeSum(testCase["s"]))
	}
}
