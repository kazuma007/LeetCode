// https://leetcode.com/submissions/detail/458727648/

package main

import "fmt"

func removeDuplicates(nums []int) int {
	i := 0
	for j := range nums {
		if nums[i] < nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 1, 1, 2, 3, 4, 5, 5, 5, 5}))
}
