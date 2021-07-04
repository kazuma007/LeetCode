// https://leetcode.com/submissions/detail/517190798/

package main

import (
	"fmt"
	"sort"
)

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[len(nums)-1]
	}

	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i += 2 {
		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}

	return nums[len(nums)-1]
}

func main() {
	var nums = []int{2, 2, 1}
	fmt.Println(singleNumber(nums))
}
