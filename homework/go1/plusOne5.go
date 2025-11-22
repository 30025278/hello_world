package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	var a = make([]int, len(digits), len(digits)+1)
	copy(a, digits)
	for i := len(digits) - 1; i >= 0; i-- {
		a[i]++
		if a[i] == 10 {
			a[i] = 0
		} else {
			return a
		}
	}
	if a[0] == 0 {
		b := append([]int{1}, a...)
		return b
	}
	return a
}

func main() {
	digits := []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}
	fmt.Println(plusOne(digits))
}
