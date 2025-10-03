package main

import (
	"fmt"
)

var testCases = []map[string]int{
	{
		"s": "babad",
	},
	{
		"s": "cbbd",
	},
}

func function(s string) string {
	return ""
}

func main() {
	for _, testCase := range testCases {
		fmt.Println(function(testCase["s"])
	}
}
