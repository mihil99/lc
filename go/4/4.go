package main

import (
	"fmt"
)

var testCases = []map[string][]int{
	{
		"nums1": {1, 3},
		"nums2": {2},
	},
	{
		"nums1": {1, 2},
		"nums2": {3, 4},
	},
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// binary approach
	return 0.0
}

// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
// 	i1 := 0
// 	i2 := 0
// 	curr := -1
// 	prev := -1
// 	total := len(nums1) + len(nums2)
//
// 	// fmt.Println("-------------------------------------------------")
// 	for i := 0; i <= total/2; i++ {
// 		// fmt.Printf("(%d/%d)\t[Total %d]  [i1 %d]  [i2 %d]  [curr %d]  [prev %d]\n", i, total/2, total, i1, i2, curr, prev)
// 		if i1 >= len(nums1) {
// 			prev = curr
// 			curr = nums2[i2]
// 			i2++
// 			continue
// 		}
// 		if i2 >= len(nums2) {
// 			prev = curr
// 			curr = nums1[i1]
// 			i1++
// 			continue
// 		}
// 		if nums1[i1] <= nums2[i2] {
// 			prev = curr
// 			curr = nums1[i1]
// 			i1++
// 		} else {
// 			prev = curr
// 			curr = nums2[i2]
// 			i2++
// 		}
// 	}
//
// 	if total%2 == 0 && prev != -1 {
// 		return float64(curr+prev) / 2
// 	}
// 	return float64(curr)
// }

func main() {
	for _, testCase := range testCases {
		fmt.Println(findMedianSortedArrays(testCase["nums1"], testCase["nums2"]))
	}
}
