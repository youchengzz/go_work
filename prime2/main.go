package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func putNum(ch chan int) {
	for i := 2; i < 120000; i++ {
		ch <- i
	}
	close(ch)
	wg.Done()
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	for v := range intChan {
		var flag = true
		for i := 2; i < v; i++ {
			if v%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan <- v
		}
	}
	exitChan <- true
	wg.Done()
}

func printNum(primeNum chan int) {
	// for v := range primeNum {
	// 	fmt.Println(v)
	// }
	wg.Done()
}

func main() {
	start := time.Now().Unix()
	fmt.Println(start)
	size := 16
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 50000)
	exitChan := make(chan bool, size)
	wg.Add(1)
	go putNum(intChan)
	for i := 1; i <= size; i++ {
		wg.Add(1)
		go primeNum(intChan, primeChan, exitChan)
	}
	wg.Add(1)
	go printNum(primeChan)
	wg.Add(1)
	go func() {
		for i := 0; i < size; i++ {
			<-exitChan
		}
		close(primeChan)
		wg.Done()
	}()
	wg.Wait()
	end := time.Now().Unix()
	fmt.Println(end)
}
