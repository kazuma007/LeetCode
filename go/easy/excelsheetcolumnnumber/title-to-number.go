// https://leetcode.com/submissions/detail/520006940/

package main

import "fmt"

func titleToNumber(columnTitle string) int {
	var total = 0
	for _, c := range columnTitle {
		total = total*26 + int(c-'A'+1)
	}
	return total
}

func main() {
	fmt.Println(titleToNumber("AB"))
}
