package main

import (
	"fmt"
	"math"
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
	{
		"nums1": func() []int {
			nums := make([]int, 1000000) // 1 million elements
			for i := 0; i < 1000000; i++ {
				nums[i] = i // Numbers from 0 to 999,999
			}
			return nums
		}(),
		"nums2": func() []int {
			nums := make([]int, 1500000) // 1.5 million elements
			for i := 0; i < 1500000; i++ {
				nums[i] = i + 500000 // Numbers from 500,000 to 1,999,999
			}
			return nums
		}(),
	},
}

func newIndex(left int, right int) int {
	return (left + right) / 2
}

func getValueFromArray(index int, nums []int) int {
	if 0 <= index && index < len(nums) {
		return nums[index]
	}
	if index < 0 {
		return int(-math.Pow(10, 6) - 1)
	}
	return int(math.Pow(10, 6) + 1)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// lengthts
	l1 := len(nums1)
	l2 := len(nums2)

	low := 0
	high := len(nums1)

	// partition points,
	// with the partition point meaning the last value of the left slice of the array
	b1 := newIndex(low, high)
	b2 := (l1+l2+1)/2 - b1

	for {
		maxLeft1 := getValueFromArray(b1-1, nums1)
		minRight1 := getValueFromArray(b1, nums1)
		maxLeft2 := getValueFromArray(b2-1, nums2)
		minRight2 := getValueFromArray(b2, nums2)

		if maxLeft1 <= minRight2 && maxLeft2 <= minRight1 {
			if (l1+l2)%2 == 0 {
				return float64(max(maxLeft1, maxLeft2)+min(minRight1, minRight2)) / 2
			}
			return float64(max(maxLeft1, maxLeft2))
		}
		if maxLeft1 > minRight2 {
			high = b1
			b1 = newIndex(low, high)
			b2 = (l1+l2+1)/2 - b1
		} else {
			low = b1 + 1
			b1 = newIndex(low, high)
			b2 = (l1+l2+1)/2 - b1
		}
	}
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
