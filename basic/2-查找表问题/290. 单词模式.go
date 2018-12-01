package __查找表问题

import "strings"

/*
给定一种 pattern(模式) 和一个字符串 str ，判断 str 是否遵循相同的模式。

这里的遵循指完全匹配，例如， pattern 里的每个字母和字符串 str 中的每个非空单词之间存在着双向连接的对应模式。

示例1:

输入: pattern = "abba", str = "dog cat cat dog"
输出: true
示例 2:

输入:pattern = "abba", str = "dog cat cat fish"
输出: false
示例 3:

输入: pattern = "aaaa", str = "dog cat cat dog"
输出: false
示例 4:

输入: pattern = "abba", str = "dog dog dog dog"
输出: false
说明:
你可以假设 pattern 只包含小写字母， str 包含了由单个空格分隔的小写字母。
 */
func main() {
	println(wordPattern("aaaa", "dog cat cat dog"))
	println(wordPattern("abba", "dog cat cat dog"))
	println(wordPattern("abba", "dog dog dog dog"))
}

func wordPattern(pattern string, str string) bool {
	patternMap := make(map[byte]int)
	wordMap := make(map[string]int)

	words := strings.Split(str, " ")

	if len(words) != len(pattern) {
		return false
	}

	for index, word := range words {
		vp, ok1 := patternMap[pattern[index]]
		vw, ok2 := wordMap[word]
		if ok1 && ok2 && vp == vw {
			continue
		} else if !ok1 && !ok2 {
			patternMap[pattern[index]] = index
			wordMap[word] = index
		} else {
			return false
		}
	}
	return true
}
