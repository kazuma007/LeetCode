// https://leetcode.com/submissions/detail/459493394/

package main

import "fmt"

func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
		if i > 0 && nums[i-1] < target && nums[i] > target {
			return i
		}
		if i == len(nums)-1 && nums[i] < target {
			return i + 1
		}
	}
	return 0
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 4, 5}, 2))
}
