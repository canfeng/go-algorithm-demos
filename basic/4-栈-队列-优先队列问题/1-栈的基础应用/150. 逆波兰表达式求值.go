package __栈的基础应用

import (
	"algorithm-demo/util"
	"strconv"
)

/*
根据逆波兰表示法，求表达式的值。

有效的运算符包括 +, -, *, / 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。

说明：

整数除法只保留整数部分。
给定逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。

示例 1：

输入: ["2", "1", "+", "3", "*"]
输出: 9
解释: ((2 + 1) * 3) = 9

示例 2：

输入: ["4", "13", "5", "/", "+"]
输出: 6
解释: (4 + (13 / 5)) = 6

示例 3：

输入: ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
输出: 22
解释:
  ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
= ((10 * (6 / (12 * -11))) + 17) + 5
= ((10 * (6 / -132)) + 17) + 5
= ((10 * 0) + 17) + 5
= (0 + 17) + 5
= 17 + 5
= 22
*/
func main() {
	println(evalRPN([]string{"2", "1", "+", "3", "*"}))
	println(evalRPN([]string{"4", "13", "5", "/", "+"}))
	println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}

func evalRPN(tokens []string) int {
	stack := &util.Stack{}
	for _, value := range tokens {
		if value == "+" || value == "-" || value == "*" || value == "/" {
			//2. 当遇到运算符的时候，连续从栈顶取出两个数字进行运算
			num2 := stack.Top()
			if num2 == nil {
				panic("wrong input")
			}
			stack.Pop()
			num1 := stack.Top()
			if num1 == nil {
				panic("wrong input")
			}
			stack.Pop()

			res := calc(num1.(int), num2.(int), value)
			//3. 将运算结果压入栈顶
			stack.Push(res)
		} else {
			//1. 当遇到数字是直接压入栈顶
			if n, ok := strconv.Atoi(value); ok == nil {
				stack.Push(n)
			} else {
				panic("wrong input")
			}
		}
	}
	if stack.Empty() {
		return 0
	}
	result := stack.Top()
	stack.Pop()
	return result.(int)
}

func calc(num1 int, num2 int, op string) int {
	switch op {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		return num1 / num2
	}
	panic("wrong operator")
}
