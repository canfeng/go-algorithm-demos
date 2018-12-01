package __查找表问题

import (
	"github.com/shopspring/decimal"
)

/*
给定一个二维平面，平面上有 n 个点，求最多有多少个点在同一条直线上。

示例 1:

输入: [[1,1],[2,2],[3,3]]
输出: 3
解释:
^
|
|        o
|     o
|  o
+------------->
0  1  2  3  4
示例 2:

输入: [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
输出: 4
解释:
^
|
|  o
|     o        o
|        o
|  o        o
+------------------->
0  1  2  3  4  5  6
 */
func main() {
	//println(maxPoints([]Point{{1, 1}, {2, 2}, {3, 3}}))
	//println(maxPoints([]Point{{2, 3}, {3, 3}, {-5, 3}}))
	//println(maxPoints([]Point{{0, 0}}))
	//println(maxPoints([]Point{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}))
	println(maxPoints([]Point{{-4,-4},{-8,-582},{-3,3},{-9,-651},{9,591}}))

	//println(gcd(2, 3))
	//println(gcd(2, 4))
}

type Point struct {
	X int
	Y int
}

/**
 * Definition for a point.
 * type Point struct {
 *     X int
 *     Y int
 * }
 */
func maxPoints_1(points []Point) int {
	/*
	时间复杂度：O(n²)
	空间复杂度：O(n)
	 */
	res := 0

	for i := 0; i < len(points); i++ {
		//使用斜率作为key，相同斜率的点数作为value
		//todo 这种方式存在除法运算，有精度缺失问题，可能有些case无法通过
		kMap := map[string]int{}
		pi := points[i]
		repeat := 1
		//直线一般式方程 Ax+By+C=0
		for j := 0; j < len(points); j++ {
			if i == j {
				continue
			}
			pj := points[j]
			//记录位置相同的点的数量，需要计入同一直线上
			if pi.X == pj.X && pi.Y == pj.Y {
				repeat ++
				continue
			}
			//如果 x 一样,单独计算
			if pi.X == pj.X {
				kMap["INF"]++
				continue
			}
			k := decimal.NewFromFloat(float64(pi.Y - pj.Y)).DivRound(decimal.NewFromFloat(float64(pi.X-pj.X)), 16)
			kMap[k.String()]++
		}

		max := 0
		for _, value := range kMap {
			if value > max {
				max = value
			}
		}
		if max+repeat > res {
			res = max + repeat
		}
	}

	return res
}

type F struct {
	fz int //分子
	fm int //分母
}

func maxPoints(points []Point) int {
	/*
	时间复杂度：O(n²)
	空间复杂度：O(n)
	 */
	res := 0

	for i := 0; i < len(points); i++ {
		//使用斜率作为key，相同斜率的点数作为value
		//TODO 但是存储的时候将斜率以分数的形式存储，比如1/2、5/6，分开存储分子fz分母fm，但是需要注意公约数问题，比如1/2 和 2/4是同样的斜率
		kMap := map[F]int{}
		pi := points[i]
		repeat := 0
		//直线一般式方程 Ax+By+C=0
		for j := 0; j < len(points); j++ {
			pj := points[j]
			f := F{pi.Y - pj.Y, pi.X - pj.X}
			if f.fz == 0 && f.fm == 0 {
				//重复的点
				repeat++
				continue
			}
			if f.fz == 0 {
				//y相同
				kMap[F{0, 1}]++
				continue
			}
			if f.fm == 0 {
				//x相同
				kMap[F{1, 0}]++
				continue
			}

			gc := gcd(f.fz, f.fm)
			if gc != 1 {
				f.fz /= gc
				f.fm /= gc
			}

			kMap[f]++
		}

		max := 0
		for _, value := range kMap {
			if value > max {
				max = value
			}
		}
		if max+repeat > res {
			res = max + repeat
		}
	}

	return res
}

func gcd(a, b int) int { //最大公约数
	if b != 0 {
		return gcd(b, a%b)
	}
	return a
}
