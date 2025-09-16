package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func fn1(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("写入数据", i)
		time.Sleep(time.Millisecond * 100)
	}
	close(ch)
	wg.Done()
}

func fn2(ch chan int) {
	for v := range ch {
		fmt.Println("读取数据:", v)
		time.Sleep(time.Millisecond * 10)
	}
	wg.Done()
}

func main() {
	ch := make(chan int, 10)
	wg.Add(1)
	go fn1(ch)
	wg.Add(1)
	go fn2(ch)
	wg.Wait()
}
