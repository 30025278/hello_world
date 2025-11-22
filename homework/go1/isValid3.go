package main

import (
	"fmt"
	"strings"
)

func isValid(s string) bool {
	arr := strings.Split(s, "")
	arr2 := make([]string, 0, len(arr))
	arr3 := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}
	for _, v := range arr {
		if v == "(" || v == "{" || v == "[" {
			arr2 = append(arr2, v)
		} else if len(arr2) == 0 || arr2[len(arr2)-1] != arr3[v] {
			return false
		} else if arr2[len(arr2)-1] == arr3[v] {
			arr2 = arr2[:(len(arr2) - 1)]
		}
	}
	return len(arr2) == 0
}

func main() {
	str1 := "{[]}"
	aaa := isValid(str1)
	fmt.Println(aaa)
}
