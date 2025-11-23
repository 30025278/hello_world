package main

import (
	"fmt"
	"time"
)

// 题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
func oddNum() {
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println("Odd:", i)
		}
	}
}

func evenNum() {
	for i := 2; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Even:", i)
		}
	}
}
func main() {
	go oddNum()
	go evenNum()
	time.Sleep(200 * time.Millisecond)
}
