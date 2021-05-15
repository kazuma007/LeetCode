// https://leetcode.com/submissions/detail/459560221/

package main

import (
	"fmt"
)

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}

	for i := 1; i <= x/2; i++ {
		if i*i <= x && x < (i+1)*(i+1) {
			return i
		}
	}
	return 1
}

func main() {
	fmt.Println(mySqrt(4))
}
