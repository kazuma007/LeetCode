// https://leetcode.com/submissions/detail/458327329/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isPalindrome(x int) bool {
	slice := strings.Split(strconv.Itoa(x), "")
	for i := 0; i < len(slice); i++ {
		if slice[i] != slice[len(slice)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(1234321))
}

// If you write like this, the data type can be int32.
// for _, char := range strconv.Itoa(x) {
// 	fmt.Println(char)
// }
//
// result
// 49
// 50
// 51
// 52
// 51
// 50
// 49
// true
