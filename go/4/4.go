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
	var i1, i2 int
	if nums1[0] <= nums2[len(nums2)] {
		i1 = 0
		i2 = len(nums2) - 1
	} else {
		i1 = len(nums1) - 1
		i2 = 0
	}

	// for {
	// go to center of 2 arrays
	// incrementing and decrementing depending on relation of 2 values
	// }

	return 0.0
}

func main() {
	for _, testCase := range testCases {
		fmt.Println(findMedianSortedArrays(testCase["nums1"], testCase["nums2"]))
	}
}
