// https://leetcode.com/problems/container-with-most-water/

package main

import (
	"fmt"
)

var testCases = []map[string][]int{
	{
		"heights": {1, 8, 6, 2, 5, 4, 8, 3, 7},
	},
	{
		"heights": {1, 1},
	},
}

func maxArea(heights []int) int {
	area := func(left int, right int) int {
		return (right - left) * min(heights[left], heights[right])

	}
	left, right := 0, len(heights)-1
	maxArea := area(left, right)
	for right > left {
		if heights[left] < heights[right] {
			left++
			maxArea = max(maxArea, area(left, right))
		} else {
			right--
			maxArea = max(maxArea, area(left, right))
		}
	}
	return maxArea
}

func main() {
	for _, testCase := range testCases {
		fmt.Println(maxArea(testCase["heights"]))
	}
}
