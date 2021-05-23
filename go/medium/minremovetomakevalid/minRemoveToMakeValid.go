// https://leetcode.com/submissions/detail/497107014/

package main

import (
	"fmt"
	"strings"
)

func minRemoveToMakeValid(s string) string {

	list := []int{}
	leftParenthesis := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			list = append(list, i)
			leftParenthesis++
		} else if s[i] == ')' {
			if leftParenthesis > 0 {
				leftParenthesis--
				list = list[:len(list)-1]
			} else {
				list = append(list, i)
			}
		}
	}

	if len(list) == 0 {
		return s
	}

	ans := []string{}
	first := 0
	last := 0
	for j := 0; j < len(list); j++ {
		last = list[j]
		if last != first {
			ans = append(ans, s[first:last])
		}
		first = last + 1
	}
	if first < len(s) {
		ans = append(ans, s[first:len(s)])
	}

	return strings.Join(ans, "")
}

func main() {
	fmt.Println(minRemoveToMakeValid(")()())"))
	fmt.Println(minRemoveToMakeValid(")))t((u)"))
	fmt.Println(minRemoveToMakeValid("(((t))u("))
}
