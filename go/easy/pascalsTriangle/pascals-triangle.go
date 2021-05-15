// https://leetcode.com/submissions/detail/459560221/

package main

import (
	"fmt"
)

func pascalsTriangles(num int) [][]int {
	ans := [][]int{}

	if num == 0 {
		return ans
	}
	for i := 0; i < num; i++ {
		ans = append(ans, []int{})
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				ans[i] = append(ans[i], 1)
			} else {
				ans[i] = append(ans[i], ans[i-1][j-1]+ans[i-1][j])
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(pascalsTriangles(5))
}
