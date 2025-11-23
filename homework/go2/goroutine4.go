package main

import (
	"fmt"
	"sync"
	"time"
)

// 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
func oddnum() time.Duration {
	startTime := time.Now()
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println("Odd:", i)
		}
	}
	return time.Since(startTime)
}

func evennum() time.Duration {
	startTime := time.Now()
	for i := 2; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Even:", i)
		}
	}
	return time.Since(startTime)
}
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	var oddDuration, evenDuration time.Duration
	go func() {
		defer wg.Done()
		oddDuration = oddnum()
	}()
	go func() {
		defer wg.Done()
		evenDuration = evennum()
	}()
	wg.Wait()
	fmt.Println("\n执行时间统计：")
	fmt.Printf("奇数打印函数执行时间: %v\n", oddDuration)
	fmt.Printf("偶数打印函数执行时间: %v\n", evenDuration)
}
