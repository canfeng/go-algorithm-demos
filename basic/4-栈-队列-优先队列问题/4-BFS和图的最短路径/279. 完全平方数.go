package main

import (
	"log"
	"math"
)

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
	//log.Println(numSquares(12))
	//log.Println(numSquares(13))
	log.Println(numSquares_2(12))
	log.Println(numSquares_2(13))
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
		//移除队首
		queue = queue[1:len(queue):len(queue)]
		num := m["num"]
		step := m["step"]
		log.Printf("num=%d; step=%d \n", num, step)

		if num == 0 {
			return step
		}

		for i := 1; num-i*i >= 0; i++ {
			a := num - i*i
			if !visited[a] {
				queue = append(queue, map[string]int{"num": a, "step": step + 1})
				visited[a] = true
			}
		}
	}

	panic("no solution")
}

//递归方式
func numSquares_2(n int) int {
	if n <= 0 {
		panic("wrong input")
	}

	mem := map[int]int{}
	return numSquaresCur(n, mem)

}

func numSquaresCur(n int, mem map[int]int) int {
	if n == 0 {
		return 0
	}
	if _, ok := mem[n]; ok {
		return mem[n]
	}
	log.Printf("num=%d \n", n)
	res := math.MaxInt32
	for i := 1; n-i*i >= 0; i++ {
		res = min(res, 1+numSquaresCur(n-i*i, mem))
		mem[n] = res
	}
	return mem[n]
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
