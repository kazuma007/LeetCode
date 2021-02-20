// https://leetcode.com/submissions/detail/458332080/

package main

import (
	"fmt"
)

var RomanNumeral = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	sum := 0
	largest := 0
	for i := len(s) - 1; i >= 0; i-- {
		num := RomanNumeral[string(s[i])]
		if num >= largest {
			sum += num
			largest = num
			continue
		}
		sum -= num
	}
	return sum
}

func main() {
	fmt.Println(romanToInt("XXX"))
}
