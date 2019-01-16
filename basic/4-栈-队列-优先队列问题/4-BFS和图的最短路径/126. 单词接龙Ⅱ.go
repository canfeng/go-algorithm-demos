package main

import "container/list"

/*
给定两个单词（beginWord 和 endWord）和一个字典 wordList，找出所有从 beginWord 到 endWord 的最短转换序列。转换需遵循如下规则：

每次转换只能改变一个字母。
转换过程中的中间单词必须是字典中的单词。
说明:

如果不存在这样的转换序列，返回一个空列表。
所有单词具有相同的长度。
所有单词只由小写字母组成。
字典中不存在重复的单词。
你可以假设 beginWord 和 endWord 是非空的，且二者不相同。

示例 1:

输入:
beginWord = "hit",
endWord = "cog",
wordList = ["hot","dot","dog","lot","log","cog"]

输出:
[
  ["hit","hot","dot","dog","cog"],
  ["hit","hot","lot","log","cog"]
]

示例 2:

输入:
beginWord = "hit"
endWord = "cog"
wordList = ["hot","dot","dog","lot","log"]

输出: []

解释: endWord "cog" 不在字典中，所以不存在符合要求的转换序列。
*/
func main() {

}

type item1 struct {
	step int
	word string
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	var result [][]string
	if len(wordList) == 0 || beginWord == endWord {
		panic("wrong input")
	}
	wordLen := len(beginWord)
	//存放所有单词，访问过的为true，未访问过的为false
	dict := make(map[string]bool)
	for _, word := range wordList {
		if len(word) != wordLen {
			panic("wrong input")
		}
		dict[word] = true
	}
	//确保没有重复单词
	if len(dict) != len(wordList) {
		panic("wrong input")
	}
	queue := list.New()
	queue.PushBack(&item1{step: 1, word: beginWord})

	//转换序列
	var seq []string
	for queue.Len() > 0 {

		front := queue.Front() //取出队首单词
		queue.Remove(front)    //移除对首
		wordItem := front.Value.(*item1)

		seq = seq

		if wordItem.word == endWord {
			//找到一个完整序列
			result = append(result, seq[:wordItem.step+1])
			continue
		}

		chars := []rune(wordItem.word)

		for i := 0; i < len(chars); i++ {
			for c := 'a'; c <= 'z'; c++ {
				if chars[i] == c {
					continue
				}
				chars[i] = c
				newWord := string(chars)
				if enabled, ok := dict[newWord]; ok && enabled {
					//如果存在单词字典中，并且未被使用过
					queue.PushBack(&item1{word: newWord, step: wordItem.step + 1})
					dict[newWord] = true
				}
			}
		}
	}

	return result
}
