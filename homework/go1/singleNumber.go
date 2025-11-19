package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	for i, v := range nums {
		nums2 := make([]int, len(nums)-1)
		copy(nums2, nums[:i])
		copy(nums2[i:], nums[i+1:])
		b := 1
		for _, a := range nums2 {
			if a == v {
				b = 0
				break
			}
		}
		if b == 1 {
			return v
		}
	}
	return nums[len(nums)-1]
}

func main() {
	num := []int{1, 0, 1}
	fmt.Println(singleNumber(num))
}
