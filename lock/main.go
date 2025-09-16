package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mutex sync.Mutex

var map1 = make(map[int]int, 0)

func calcSum(num int) {
	mutex.Lock()
	var sum = 1
	for i := 1; i <= num; i++ {
		sum *= i
	}
	map1[num] = sum
	fmt.Printf("key:%v, value:%v \n", num, sum)
	mutex.Unlock()
	time.Sleep(time.Millisecond)
	wg.Done()
}

func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go calcSum(i)
	}
	wg.Wait()
}
