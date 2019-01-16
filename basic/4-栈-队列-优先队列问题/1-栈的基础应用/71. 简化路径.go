package __栈的基础应用

import (
	"algorithm-demo/util"
	"strings"
)

/*
给定一个文档 (Unix-style) 的完全路径，请进行路径简化。

例如，
path = "/home/", => "/home"
path = "/a/./b/../../c/", => "/c"

边界情况:

- 你是否考虑了 路径 = "/../" 的情况？
在这种情况下，你需返回 "/" 。

- 此外，路径中也可能包含多个斜杠 '/' ，如 "/home//foo/" 。
在这种情况下，你可忽略多余的斜杠，返回 "/home/foo" 。
*/
func main() {
	println(simplifyPath("/../"))
	println(simplifyPath("/../cin"))
	println(simplifyPath("/home/"))
	println(simplifyPath("/a/./b/../../c/"))
}

func simplifyPath(path string) string {
	//1. 先将path用“/”分割为数组
	arr := strings.Split(path, "/")

	stack := &util.Stack{}
	//2. 遍历数组
	for _, value := range arr {
		//3. 如果遇到的字符串不是“.”或“..”,则推入栈顶; 否则如果是“..”，需要移除栈顶元素，其它的不作处理
		if value == "." || value == "" {
			continue
		}
		if value == ".." {
			if !stack.Empty() {
				stack.Pop()
			}
		} else {
			stack.Push(value)
		}
	}
	//4. 最后用“/”从栈底开始连接栈中元素即可
	if stack.Empty() {
		return "/"
	}
	res := ""
	for _, v := range stack.Element {
		res += "/" + v.(string)
	}
	return res
}
