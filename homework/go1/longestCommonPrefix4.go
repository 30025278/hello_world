package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 || len(strs) == 0 {
		return strs[0]
	}
	for q, v := range strs[0] {
		for _, i := range strs {
			if q >= len(i) || string(v) != string(i[q]) {
				return strs[0][:q]
			}
		}
	}
	return strs[0]
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}
