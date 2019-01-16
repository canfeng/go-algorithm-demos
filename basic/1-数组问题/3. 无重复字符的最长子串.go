package __数组问题

import (
	"algorithm-demo/util"
)

/**
给定一个字符串，找出不含有重复字符的最长子串的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 无重复字符的最长子串是 "abc"，其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 无重复字符的最长子串是 "b"，其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 无重复字符的最长子串是 "wke"，其长度为 3。
     请注意，答案必须是一个子串，"pwke" 是一个子序列 而不是子串。
*/
func main() {
	println(lengthOfLongestSubstring_1("abcabcbb"))
	println(lengthOfLongestSubstring_1("bbbbb4msddd"))
	println(lengthOfLongestSubstring_1("pwwkew"))
	println(lengthOfLongestSubstring_2("abcabcbb"))
	println(lengthOfLongestSubstring_2("bbbbb4msddd"))
	println(lengthOfLongestSubstring_2("pwwkew"))
}

//时间复杂度O(n)
//空间复杂度O(n)
func lengthOfLongestSubstring_1(s string) int {
	//记录在区间[l...r]中字符的出现频率
	var freq [256]int

	res := 0
	l := 0
	r := -1 //[l...r]表示最长不重复的字串

	step := 0

	for l < len(s) {
		if r+1 < len(s) && freq[s[r+1]] == 0 {
			r++
			freq[s[r]] = r + 1
		} else {
			freq[s[l]]--
			l++
		}
		res = util.Max(res, r-l+1)
		step++
	}

	println("func", util.GetFuncName(), "array length =>", len(s), "| total steps =>", step)
	return res
}

//时间复杂度O(n^2)
//空间复杂度O(1)
func lengthOfLongestSubstring_2(s string) int {
	step := 0

	res := 0
	l := 0
	r := 0 //滑动窗口为[l...r]
	for r < len(s) {
		//查找与s[r]重复的字符，l移动到该字符的索引+1的位置
		for i := l; i < r; i++ {
			if s[i] == s[r] {
				l = i + 1
			}
			step++
		}
		res = util.Max(res, r-l+1)
		r++
	}

	println("func", util.GetFuncName(), "array length =>", len(s), "| total steps =>", step)
	return res
}
