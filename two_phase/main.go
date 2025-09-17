package main

import (
	"fmt"
	"sync"
	"time"
)

/*
* ✅指针
题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
考察点 ：指针的使用、值传递与引用传递的区别。
*
*/
func point(num *int) int {
	a := *num
	for i := 0; i < 10; i++ {
		a++
	}
	return a
}

/*
*
题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
考察点 ：指针运算、切片操作。
*
*/
func slices(s []int) []int {
	for i, v := range s {
		s[i] = v * 2
	}
	return s
}

/*
*
✅Goroutine
题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
考察点 ： go 关键字的使用、协程的并发执行。
*
*/
var wg sync.WaitGroup

func goroutine1() {
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			time.Sleep(time.Second)
			fmt.Println("奇数", i)
		}
	}
	wg.Done()

}

func goroutine2() {
	for i := 2; i <= 10; i++ {
		if i%2 == 0 {
			time.Sleep(time.Second * 2)
			fmt.Println("偶数", i)
		}
	}
	wg.Done()
}

/*
*
题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
考察点 ：协程原理、并发任务调度。
*
*/
func shedule() {

}

/*
*
✅面向对象
题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
考察点 ：接口的定义与实现、面向对象编程风格。
*
*/
type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
}

func (r Rectangle) Area() {
	fmt.Println("调用Rectangle-Area方法")

}

func (r Rectangle) Perimeter() {
	fmt.Println("调用Rectangle-Perimeter方法")
}

type Circle struct {
}

func (c Circle) Area() {
	fmt.Println("调用Circle-Area方法")

}

func (c Circle) Perimeter() {
	fmt.Println("调用Circle-Perimeter方法")
}

/*
*
题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
考察点 ：组合的使用、方法接收者。
*
*/
type Person struct {
	Name string
	Age  int
}
type Employee struct {
	Person     Person
	EmployeeID string
}

func (e Employee) PrintInfo() {
	fmt.Printf("员工信息:%#v", e)
}

/*
*
✅Channel
题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
考察点 ：通道的基本使用、协程间通信。
*
*/
func Channel1(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
	close(ch)
	wg.Done()
}

func Channel2(ch chan int) {
	for v := range ch {
		fmt.Println(v)
	}
	wg.Done()
}

/**
题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
考察点 ：通道的缓冲机制。
**/

/*
*
✅锁机制
题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ： sync.Mutex 的使用、并发数据安全。
*
*/
var mutex sync.Mutex

var count = 0

func lock(num int) {
	for i := 0; i < 1000; i++ {
		mutex.Lock()
		count++
		mutex.Unlock()
		fmt.Printf("协程%v count= %v\n", num, count)
		time.Sleep(time.Second)
	}

	wg.Done()
}

/**
题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ：原子操作、并发数据安全。
**/

func main() {
	// var a = 0
	// b := &a
	// fmt.Println(point(b))

	// var sli []int = []int{1, 2, 3}
	// fmt.Println(slices(sli))

	// wg.Add(1)
	// go goroutine1()
	// wg.Add(1)
	// go goroutine2()
	// wg.Wait()

	// rectangle := &Rectangle{}
	// rectangle.Area()
	// rectangle.Perimeter()

	// circle := &Circle{}
	// circle.Area()
	// circle.Perimeter()

	// employee := &Employee{
	// 	EmployeeID: "1",
	// 	Person: Person{
	// 		Name: "张三",
	// 		Age:  20,
	// 	},
	// }
	// employee.PrintInfo()

	// ch := make(chan int, 10)
	// wg.Add(1)
	// go Channel1(ch)
	// wg.Add(1)
	// go Channel2(ch)
	// wg.Wait()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		fmt.Println("开启协程", i)
		go lock(i)
	}
	wg.Wait()
}
