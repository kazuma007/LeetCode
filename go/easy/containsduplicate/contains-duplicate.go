// https://leetcode.com/submissions/detail/510459664/

package main

import (
	"fmt"
	"sort"
)

func containsDuplicate(nums []int) bool {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			return true
		}
	}

	return false
}

func main() {
	nums := []int{1, 6, 3, 6}
	fmt.Println(containsDuplicate(nums))
}
