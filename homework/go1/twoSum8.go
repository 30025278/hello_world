package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	for slow := 0; slow < len(nums)-1; slow++ {
		for fast := slow + 1; fast < len(nums); fast++ {
			if nums[slow]+nums[fast] == target {
				return []int{slow, fast}
			}
		}
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
