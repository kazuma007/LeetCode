// https://leetcode.com/submissions/detail/503797885/

package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	var ans = []string{}
	arr := strings.Split(s, " ")
	for _, s := range arr {
		runes := []rune(s)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		ans = append(ans, string(runes))
	}
	return strings.Join(ans, " ")
}

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}
