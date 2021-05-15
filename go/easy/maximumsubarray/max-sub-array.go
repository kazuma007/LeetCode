// https://leetcode.com/submissions/detail/459537648/

package main

import "fmt"

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		fmt.Println(nums[i])
	}
	return max(nums)
}

func max(nums []int) int {
	max := nums[0]
	for _, num := range nums {
		if max < num {
			max = num
		}
	}
	return max
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
