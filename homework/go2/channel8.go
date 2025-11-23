package main

import (
	"fmt"
	"sync"
)

// 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 100; i++ {
		fmt.Printf("发送: %d\n", i)
		ch <- i
	}
	close(ch)
	fmt.Println("发送完成，通道已关闭")
}

func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Printf("接收: %d\n", num)
	}
	fmt.Println("接收完成")
}

func main() {
	ch := make(chan int, 10) // 创建一个带有缓冲区大小为10的通道
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go producer(ch, wg)
	go consumer(ch, wg)
	wg.Wait()
	fmt.Println("所有操作完成")
}
