package main

import (
	"algorithm-demo/util"
)

/*
给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的连续子数组。如果不存在符合条件的连续子数组，返回 0。

示例:

输入: s = 7, nums = [2,3,1,2,4,3]
输出: 2
解释: 子数组 [4,3] 是该条件下的长度最小的连续子数组。
进阶:

如果你已经完成了O(n) 时间复杂度的解法, 请尝试 O(n log n) 时间复杂度的解法。
 */
func main() {
	println(minSubArrayLen_1(7, []int{2, 3, 1, 2, 4, 3}))
	println(minSubArrayLen_1(11, []int{2, 3, 1, 2, 4, 3, 23, 7, 5, 9, 2, 4}))
	println(minSubArrayLen_2(7, []int{2, 3, 1, 2, 4, 3}))
	println(minSubArrayLen_2(11, []int{2, 3, 1, 2, 4, 3, 23, 7, 5, 9, 2, 4}))
	println(minSubArrayLen_3(7, []int{2, 3, 1, 2, 4, 3}))
	println(minSubArrayLen_3(11, []int{2, 3, 1, 2, 4, 3, 23, 7, 5, 9, 2, 4}))
}

//暴力破解法
//时间复杂度O(n^3)
func minSubArrayLen_1(s int, nums []int) int {
	res := len(nums) + 1
	step := 0
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			sum := 0
			for k := i; k <= j; k++ {
				sum += nums[k]
				step ++
			}
			if sum >= s && j-i < res {
				res = j - i + 1
				step ++
			}
		}
	}

	println("func", util.GetFuncName(), "nums length =>", len(nums), "| total steps =>", step)

	//如果res始终都没有发生变化，则表示未找到正确的连续子数组，返回0
	if res == len(nums)+1 {
		res = 0
	}
	return res
}

//暴力破解法优化
//时间复杂度O(n^2)
func minSubArrayLen_2(s int, nums []int) int {
	res := len(nums) + 1
	step := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum >= s && j-i < res {
				res = j - i + 1
			}
			step ++
		}
	}
	println("func", util.GetFuncName(), "nums length =>", len(nums), "| total steps =>", step)
	//如果res始终都没有发生变化，则表示未找到正确的连续子数组，返回0
	if res == len(nums)+1 {
		res = 0
	}
	return res
}

//滑动窗口解法
//时间复杂度O(n)
func minSubArrayLen_3(s int, nums []int) int {
	step := 0
	res := len(nums) + 1
	l := 0
	r := -1
	sum := 0
	//最小连续子数组为[l...r]
	for l < len(nums) {
		if sum < s && r+1 < len(nums) {
			r++
			sum += nums[r]
		} else { //sum >= s

			res = r - l + 1
			sum -= nums[l]
			l++
		}
		step ++
	}

	println("func", util.GetFuncName(), "nums length =>", len(nums), "| total steps =>", step)
	if res == len(nums)+1 {
		res = 0
	}
	return res
}

func Min(a int,b int) int  {
	if a > b {
		return b
	}else{
		return a
	}
}