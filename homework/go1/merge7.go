package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	newin := make([][]int, 0, len(intervals))
	newin = append(newin, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= newin[len(newin)-1][1] {
			if intervals[i][1] >= newin[len(newin)-1][1] {
				newin = append(newin, []int{newin[len(newin)-1][0], intervals[i][1]})
				newin = append(newin[:len(newin)-2], newin[len(newin)-1:]...)
			}
		} else {
			newin = append(newin, intervals[i])
		}
	}
	return newin
}
func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(intervals))
}
