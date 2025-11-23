package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
type Counter struct {
	count int32
}

// 增加计数
func (c *Counter) Increment() {
	atomic.AddInt32(&c.count, 1)
}

// 获取当前计数
func (c *Counter) GetCount() int {
	return int(atomic.LoadInt32(&c.count))
}

func main() {
	counter := Counter{}
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
			fmt.Printf("协程 %d 完成了1000次递增操作\n", id)
		}(i)
	}
	wg.Wait()
	fmt.Printf("Final count: %d\n", counter.GetCount())
}
