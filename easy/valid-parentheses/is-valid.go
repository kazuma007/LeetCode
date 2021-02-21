// https://leetcode.com/submissions/detail/458599684/

package main

import (
	"fmt"
)

var parentheses = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
}

func isValid(s string) bool {
	var stack []rune
	for _, char := range s {
		if _, ok := parentheses[char]; !ok {
			stack = append(stack, char)
			continue
		}
		if len(stack) == 0 || stack[len(stack)-1] != parentheses[char] {
			return false
		}
		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("{[()]}"))
}
