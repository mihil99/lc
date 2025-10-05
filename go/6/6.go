package main

import (
	"fmt"
	"strings"
)

var testCases = []map[string]any{
	{
		"s": "PAYPALISHIRING",
		"numrows": 3,
	},
	{
		"s": "PAYPALISHIRING",
		"numrows": 4,
	},
	{
		"s": "A",
		"numrows": 1,
	},
}

func convert(s string, numrows int) string {
	if numrows == 1 || numrows >= len(s) {
		return s
	}
	sl := make([]string, numrows)
	for i := 0; i < numrows; i++ {
		sl[i] = ""
	}
	down := true
	r := 0
	for _, char := range s {
		// char placement logic
		sl[r] += string(char)
		// boundary checks
		if down && r == numrows - 1 {
			down = false
		}
		if !down && r == 0 {
			down = true
		}
		// moving cursor
		var dr int
		if down {
			dr = 1
		} else {
			dr = -1
		}
		r += dr
	}
	return strings.Join(sl, "")
}

func main() {
	for _, testCase := range testCases {
		s, ok := testCase["s"].(string)
		if !ok {
			continue
		}
		numrows, ok := testCase["numrows"].(int)
		if !ok {
			continue
		}
		fmt.Println(convert(s, numrows))
	}
}
