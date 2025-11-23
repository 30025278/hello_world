package main

import "fmt"

//题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
func doublenum(p *[]int) {
	for i := range *p {
		(*p)[i] = (*p)[i] * 2
	}
}

func main() {
	nums := []int{1, 2, 3, 4}
	doublenum(&nums)
	fmt.Println(nums)
}
