package main

import (
	"fmt"
	"strconv"
)

func addDigits(num int) int {
	s := strconv.Itoa(num)

	n := num
	for n >= 10 {
		s = addStr(s)
		n, _ = strconv.Atoi(s)
	}

	return n
}

func addStr(s string) string {
	total := 0
	for _, c := range s {
		n, _ := strconv.Atoi(string(c))
		total += n
	}
	return strconv.Itoa(total)
}

func main() {
	fmt.Println(addDigits(38))
}
