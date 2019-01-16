package __数组问题

import "log"

/*
给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2。
![image](https://aliyun-lc-upload.oss-cn-hangzhou.aliyuncs.com/aliyun-lc-upload/uploads/2018/07/25/question_11.jpg)
图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。

示例:

输入: [1,8,6,2,5,4,8,3,7]
输出: 49
*/
func main() {
	area := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	log.Println(area)
}

//对撞指针法
//时间复杂度O(n)
/*
算法

这种方法背后的思路在于，两线段之间形成的区域总是会受到其中较短那条长度的限制。此外，两线段距离越远，得到的面积就越大。

我们在由线段长度构成的数组中使用两个指针，一个放在开始，一个置于末尾。
此外，我们会使用变量 maxareamaxarea 来持续存储到目前为止所获得的最大面积。
在每一步中，我们会找出指针所指向的两条线段形成的区域，更新 maxareamaxarea，
并将指向<b>较短线段</b>的指针向较长线段那端移动一步。

*/
func maxArea(height []int) int {
	if len(height) < 2 {
		panic("height length not enough")
	}

	l := 0
	r := len(height) - 1
	maxarea := (r - l) * min(height[l], height[r])
	for l < r {
		minVal := min(height[l], height[r])
		tempArea := (r - l) * minVal
		if tempArea > maxarea {
			maxarea = tempArea
		}
		if minVal == height[l] {
			l++
		} else {
			r--
		}
	}
	return maxarea
}

func min(v1 int, v2 int) int {
	if v1 < v2 {
		return v1
	} else {
		return v2
	}
}
