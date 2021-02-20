package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == target-nums[i] {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{1, 2, 3}, 3))
}
