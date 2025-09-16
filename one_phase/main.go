package main

import (
	"fmt"
	"math"
	"strconv"
)

// https://leetcode.cn/problems/single-number/

// 回文数
func palindromic(x int) bool {
	str := strconv.Itoa(x)
	bytes := []byte(str)
	res := make([]byte, 0)
	for i := len(bytes); i > 0; i-- {
		s := byte(bytes[i-1])
		res = append(res, s)
	}
	fmt.Println(res)
	num := string(res)
	fmt.Println(str, num)
	if str == num {
		return true
	}
	return false
}

// 有效括号
func isValid(s string) bool {
	var brackets = [7]string{"", "(", ")", "[", "]", "{", "}"}
	var res = make([]string, 0)
	if s == brackets[0] || len(s) == 1 {
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
				if num == 0 || len(res) == 0 {
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

// 最长公共前缀
func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	charArray := []rune(strs[0])
	if len(charArray) == 0 {
		return ""
	}
	flag := make([]int, 0)

	for i := 1; i < len(strs); i++ {
		chars := []rune(strs[i])
		for j := 0; j < len(charArray); j++ {
			if j > len(chars)-1 {
				flag = append(flag, len(chars)-1)
				continue
			}
			if chars[j] != charArray[j] {
				flag = append(flag, j-1)
				break
			}
		}
	}

	return strs[0][:flag+1]
}

// 加一
func plusOne(digits []int) []int {
	var pow float64 = math.Pow10(len(digits) - 1)
	var sum int
	for i := 0; i < len(digits); i++ {
		sum = sum + digits[i]*int(pow)
		pow /= 10
	}
	sum += 1
	var res []int
	for {
		var remainder = sum % 10
		res = append(res, remainder)
		if sum < 10 {
			break
		}
		sum = sum / 10
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func plusOne2(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			fmt.Println("修改值:", digits)
			return digits
		} else {
			digits[i] = 0
		}
	}
	return append([]int{1}, digits...)
}

// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/

// https://leetcode.cn/problems/merge-intervals/

// 两数之和
func twoSum(nums []int, target int) []int {
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		tmp := target - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if tmp == nums[j] {
				return append(res, i, j)
			}
		}
	}
	return res
}

func main() {
	s := "(){}}{"
	b := isValid(s)
	fmt.Println(b)
}
