package main

import (
	"fmt"
	"sync"
)

// 1 - 30
// 31 - 60
// 61 - 90
// 91 - 120

var wg sync.WaitGroup

func findPrime(num int) {
	for i := (num-1)*30 + 1; i < num*30; i++ {
		var flag = true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = false
				break
			}
		}
		if flag {
			fmt.Println(i)
		}

	}
	wg.Done()
}

func main() {
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go findPrime(i)
	}
	wg.Wait()
}
