// https://leetcode.com/submissions/detail/458307703/

package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	words := strings.Fields(s)
	if len(words) == 0 {
		return 0
	}
	target := words[len(words)-1]
	return len(target)
}

func main() {
	fmt.Println(lengthOfLastWord("Hello, World!"))
}
