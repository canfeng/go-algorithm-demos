package __查找表问题

/*
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的一个字母异位词。

示例 1:

输入: s = "anagram", t = "nagaram"
输出: true
示例 2:

输入: s = "rat", t = "car"
输出: false
说明:
你可以假设字符串只包含小写字母。

进阶:
如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
 */
func main() {
	println(isAnagram("anagram", "nagaram"))
	println(isAnagram("rat", "car"))
	println(isAnagram("ab", "a"))
}

func isAnagram(s string, t string) bool {
	m := make(map[rune]int)
	for _, ch := range s {
		m[ch]++
	}
	for _, ch := range t {
		if v, ok := m[ch]; ok {
			if v == 1 {
				delete(m, ch)
			} else {
				m[ch]--
			}
		} else {
			return false
		}
	}

	return len(m) == 0
}
