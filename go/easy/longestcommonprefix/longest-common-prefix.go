// https://leetcode.com/submissions/detail/458566543/

package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for i := range strs[0] {
		for _, str := range strs[1:] {
			if i == len(str) || strs[0][i] != str[i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}
