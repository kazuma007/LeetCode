// https://leetcode.com/submissions/detail/500242165/

package main

import "fmt"

func findTheDifference(s string, t string) byte {
	m := make(map[int32]int)

	for _, i := range s {
		m[i]++
	}
	for _, j := range t {
		m[j]--
		if m[j] < 0 {
			return byte(j)
		}
	}

	return byte('0')
}

func main() {
	s := "abcd"
	t := "abcde"
	fmt.Println(findTheDifference(s, t))
}
