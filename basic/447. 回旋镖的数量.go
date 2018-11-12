package main

import "github.com/rs/zerolog/log"

/*
给定平面上 n 对不同的点，“回旋镖” 是由点表示的元组 (i, j, k) ，其中 i 和 j 之间的距离和 i 和 k 之间的距离相等（需要考虑元组的顺序）。

找到所有回旋镖的数量。你可以假设 n 最大为 500，所有点的坐标在闭区间 [-10000, 10000] 中。

示例:

输入:
[[0,0],[1,0],[2,0]]

输出:
2

解释:
两个回旋镖为 [[1,0],[0,0],[2,0]] 和 [[1,0],[2,0],[0,0]]
 */
func main() {
	var points = [][]int{{0, 0}, {1, 0}, {2, 0}}
	boomerangs := numberOfBoomerangs(points)
	log.Print(boomerangs)
}

func numberOfBoomerangs(points [][]int) int {
	// 解题思路：
	// 遍历所有的点，将每一个点作为枢纽点i，然后计算其它所有点到i的距离，存入map中，距离为key，点的数量为value
	// 然后遍历map，将其中value>=2的键值对取出来，排列组合所有的可能性。
	step := 0
	res := 0
	for i := 0; i < len(points); i++ {
		disMap := map[int]int{}

		for j := 0; j < len(points); j++ {
			if i != j {
				disMap[distance(points[i], points[j])]++
			}
			step++
		}

		for _, value := range disMap {
			if value >= 2 {
				res += value * (value - 1)
			}
			step++
		}
	}
	return res
}

/*
计算两个点的距离，公式根据勾股定理，dis = √((point1.x-point2.x)² + (point1.y-point2.y)²)

todo 1. 但是由于涉及到开根号，结果是浮点型，可能会出现进度缺失的情况，所以可以不开根号，直接用平方代替
todo 2. 两次平方和可能会出现整型溢出的情况，但是本题明确指出，坐标点的范围是[-10000, 10000]，所以最大的dis²=(10000 - (-10000))² + (10000 - (-10000))² = 800000000
todo	没有超过32位int类型的最大值，但是如果取值区间再大一点，可能就需要int64类型了
 */
func distance(point1 []int, point2 []int) int {
	return (point1[0]-point2[0])*(point1[0]-point2[0]) + (point1[1]-point2[1])*(point1[1]-point2[1])
}
