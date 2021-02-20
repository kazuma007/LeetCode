// Reverse Integer
// https://leetcode.com/submissions/detail/458256738/
package easy

import "math"

func reverse(x int) int {
	if x < 0 {
		return -(reverseInt(-x))
	} else {
		return reverseInt(x)
	}
}

func reverseInt(x int) int {
	ans := 0
	for x > 0 {
		ans = ans*10 + x%10
		x /= 10
	}
	if -(ans) < math.MinInt32 || ans > math.MaxInt32 {
		return 0
	}
	return ans
}
