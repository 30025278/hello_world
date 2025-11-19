package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	str := fmt.Sprintf("%d", x)
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	num := 121
	fmt.Println(isPalindrome(num))
}
