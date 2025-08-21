package main

import (
	"fmt"
)

// func twoSum(nums []int, target int) []int {
// 	for idx1, val1 := range nums {
// 		for idx2, val2 := range nums[idx1+1:] {
// 			if val1+val2 == target {
// 				return []int{idx1, idx2 + idx1 + 1}
// 			}
// 		}
// 	}
// 	return []int{}
// }

// func twoSum(nums []int, target int) []int {
// 	for idx1, val1 := range nums {
// 		for idx2, val2 := range nums[idx1+1:] {
// 			if val2 == target-val1 {
// 				return []int{idx1, idx2 + idx1 + 1}
// 			}
// 		}
// 	}
// 	return []int{}
// }

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for idx, val := range nums {
		m[val] = idx
	}
	for idx, val := range nums {
		if t, ok := m[target-val]; ok && t != idx {
			return []int{idx, t}

		}
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}
