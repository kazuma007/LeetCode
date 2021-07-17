// https://leetcode.com/submissions/detail/523799954/

package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.TrimSpace(s)

	var fixedS string
	for _, c := range s {
		if (c >= rune('A') && c <= rune('Z')) || (c >= rune('a') && c <= rune('z')) || (c >= '0' && c <= '9') {
			fixedS += strings.ToLower(string(c))
		}
	}

	for i := 0; i < len(fixedS); i++ {
		if fixedS[i] != fixedS[len(fixedS)-i-1] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome(" "))
	fmt.Println(isPalindrome("a"))
	fmt.Println(isPalindrome("0P"))
}
