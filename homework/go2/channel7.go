package main

import (
	"fmt"
	"sync"
)

// 题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
func sendOnly(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		fmt.Printf("发送: %d\n", i)
		ch <- i
	}
	close(ch)
	fmt.Println("发送完成，通道已关闭")
}

func receiveOnly(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Printf("接收: %d\n", num)
	}
	fmt.Println("接收完成")
}

func main() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go sendOnly(ch, wg)
	go receiveOnly(ch, wg)
	wg.Wait()
	fmt.Println("所有操作完成")
}
