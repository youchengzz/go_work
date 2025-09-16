package main

import "fmt"

func isValid(s string) bool {
	var brackets = [7]string{"", "(", ")", "[", "]", "{", "}"}
	var res = make([]string, 0)
	if s == brackets[0] || len(s) == 1{
		return false
	}
	// 遍历目标字符串
	for num, str := range s {
		// 判断字符是左还是右
		for i, v := range brackets {
			// 匹配到左括号
			if string(str) == v && i%2 == 1 {
				res = append(res, v)
				fmt.Println("添加一个元素", res)
			} else if string(str) == v && i%2 == 0 {
				// 第一个元素是右括号
				if num == 0  || len(res) == 0{
					return false
				}
				// 匹配到右括号
				// 判断是否是一对
				fmt.Println("判断是否是一对：", res[len(res)-1], brackets[i])
				b := res[len(res)-1] == brackets[i-1]
				if b {
					// 删除已匹配成功
					fmt.Println("匹配成功，需要删除")
					res = res[:len(res)-1]
				} else {
					return false
				}
			}
		}
	}
	return len(res) == 0
}

func main() {
	s := "(){}}{"
	b := isValid(s)
	fmt.Println(b)
}
