package __数组问题

import (
	"log"
	"unicode"
)

/*
编写一个函数，以字符串作为输入，反转该字符串中的元音字母。

示例 1:

输入: "hello"
输出: "holle"
示例 2:

输入: "leetcode"
输出: "leotcede"
说明:
元音字母不包含字母"y"。
*/
func main() {
	log.Println(reverseVowels("hello"))
	log.Println(reverseVowels("leetcode"))
}

func reverseVowels(s string) string {
	vowels := map[rune]rune{
		'a': 'a',
		'e': 'e',
		'i': 'i',
		'o': 'o',
		'u': 'u',
	}

	charArr := []rune(s)

	l := 0
	r := len(charArr) - 1
	for l < r {
		if _, ok := vowels[unicode.ToLower(charArr[l])]; !ok {
			l++
		} else if _, ok := vowels[unicode.ToLower(charArr[r])]; !ok {
			r--
		} else {
			charArr[l], charArr[r] = charArr[r], charArr[l]
			l++
			r--
		}
	}
	return string(charArr)
}
