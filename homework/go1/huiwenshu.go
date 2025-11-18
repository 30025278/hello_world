package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	str := fmt.Sprintf("%d", x)
	arr := make([]int, len(str))
	for i, ch := range str {
		arr[i] = int(ch - '0')
	}
	for i := range arr {
		arr[i] = arr[len(str)-1-i]
	}
	str1 := ""
	for _, v := range arr {
		str1 += strconv.Itoa(v)
	}
	num, _ := strconv.Atoi(str1)
	if num == x {
		return true
	} else {
		return false
	}
}

func main() {
	num := 121
	fmt.Println(isPalindrome(num))
}
