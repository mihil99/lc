// https://leetcode.com/problems/reverse-integer/
package main

import (
	"fmt"
	"math"
	// "strconv"
)

var testCases = []map[string]int{
	{
		"x": 123,
	},
	{
		"x": -123,
	},
	{
		"x": 120,
	},
}

// string based
// func reverse(x int) int {
// 	xs := strconv.Itoa(int(x))
// 	xx := ""
// 	// fmt.Printf("%d ", x)
// 	if xs[0] == '-' {
// 		xx += "-"
// 		// fmt.Printf("-> (-) %s ", xx)
// 		xs = xs[1:]
// 	}
// 	for i := 0; i < len(xs); i++ {
// 		xx += string(xs[len(xs)-1-i])
// 		// fmt.Printf("->(i) %s ", xx)
// 	}
//
// 	val, _ := strconv.Atoi(xx)
// 	// fmt.Printf("-> %d\n", val)
// 	if val > int(math.Pow(2,31))-1 || val < -int(math.Pow(2,31)) {
// 		return 0
// 	}
// 	return val
// }

// math based
func reverse(x int) int {
	positive := true
	if x < 0 {
		positive = false
		x *= -1
	}

	s := 0
	for x != 0 {
		d := x%10
		x = (x - d) / 10
		s = s*10+d
	}
	if s > int(math.Pow(2,31))-1 || s < -int(math.Pow(2,31)) {
		return 0
	}
	if !positive {
		return -s
	}
	return s
}


func main() {
	for _, testCase := range testCases {
		fmt.Println(reverse(testCase["x"]))
	}
}
