package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}

func main() {
	nums := []int{2, 11, 15, 7}
	target := 9
	fmt.Println(twoSum(nums, target))
}
