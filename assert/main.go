package main

import (
	"fmt"
	"runtime"
)

type A interface{}

func typeAssert(t interface{}) {
	switch t.(type) {
	case int:
		fmt.Println("int类型")
	case string:
		fmt.Println("string类型")
	case bool:
		fmt.Println("bool类型")
	}
	
}

func main() {
	// typeAssert("a")
	var num = runtime.NumCPU()
	fmt.Println(num)
}
