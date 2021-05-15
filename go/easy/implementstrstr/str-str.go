// https://leetcode.com/submissions/detail/459487309/

package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		j := 0
		for ; j < len(needle); j++ {
			if needle[j] != haystack[i+j] {
				break
			}
		}
		if j == len(needle) {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr("hello", "ll"))
}
