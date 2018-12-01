package __栈的基础应用

import (
	"algorithm-demo/util"
)

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true

示例 2:

输入: "()[]{}"
输出: true

示例 3:

输入: "(]"
输出: false

示例 4:
输入: "([)]"
输出: false

示例 5:

输入: "{[]}"
输出: true
 */
func main() {
	println(isValid("]"))
	println(isValid("()"))
	println(isValid("()[]{}"))
	println(isValid("(]"))
	println(isValid("([)]"))
	println(isValid("{[]}"))
}

func isValid(s string) bool {

	stack := &util.Stack{}
	for _, ch := range s {
		if ch == '(' || ch == '[' || ch == '{' {
			//1. 把遇到的所有的左边的括号推入栈中
			stack.Push(ch)
		} else {
			//2. 当遇到右边括号的时候，则去和栈顶的元素去匹配，如果匹配，则删除栈顶元素，继续遍历；否则，返回false
			if stack.Top() == nil {
				return false
			}
			left := stack.Top().(rune)
			if match(left, ch) {
				stack.Pop()
			} else {
				return false
			}
		}
	}
	//3. 最后判断栈是否为空，为空则true，否则false
	return stack.Empty()
}

func match(left rune, right rune) bool {
	if left == '(' && right == ')' {
		return true
	}
	if left == '[' && right == ']' {
		return true
	}
	if left == '{' && right == '}' {
		return true
	}
	return false
}
