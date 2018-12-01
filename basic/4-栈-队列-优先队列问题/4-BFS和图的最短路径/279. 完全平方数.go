package main

import "log"

/*
给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。

示例 1:

输入: n = 12
输出: 3
解释: 12 = 4 + 4 + 4.
示例 2:

输入: n = 13
输出: 2
解释: 13 = 4 + 9.
 */
func main() {
	log.Println(numSquares(12))
	log.Println(numSquares(13))
}

func numSquares(n int) int {
	if n <= 0 {
		panic("wrong input")
	}
	var queue []map[string]int
	queue = append(queue, map[string]int{"num": n, "step": 0})

	//保存已经用过的数
	visited := make(map[int]bool)

	for len(queue) > 0 {
		m := queue[0]
		//移除对首
		queue = queue[1:len(queue):len(queue)]
		num := m["num"]
		step := m["step"]
		log.Printf("num=%d; step=%d \n", num, step)

		if num == 0 {
			return step
		}

		for i := 1; num-i*i >= 0; i++ {
			a := num - i*i
			if _, ok := visited[a]; !ok {
				queue = append(queue, map[string]int{"num": a, "step": step + 1})
				visited[a] = true
			}
		}
	}

	panic("no solution")
}
