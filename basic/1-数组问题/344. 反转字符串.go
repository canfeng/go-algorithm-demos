package __数组问题

import "log"

/*
编写一个函数，其作用是将输入的字符串反转过来。

示例 1:

输入: "hello"
输出: "olleh"
示例 2:

输入: "A man, a plan, a canal: Panama"
输出: "amanaP :lanac a ,nalp a ,nam A"
 */
func main() {
	s := reverseString("A man, a plan, a canal: Panama")
	s1 := reverseString("hello")
	log.Println(s)
	log.Println(s1)
}

func reverseString(s string) string {
	arr := []rune(s)
	len := len(arr)
	for i := 0; i < len/2; i++ {
		arr[i], arr[len-1-i] = arr[len-1-i], arr[i]
	}
	return string(arr)
}
