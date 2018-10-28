package main

import (
	"log"
)

//1. 两数之和
//
//给定一个整数数组和一个目标值，找出数组中和为目标值的两个数。
//
//你可以假设每个输入只对应一种答案，且同样的元素不能被重复利用。
//
//示例:
//
//给定 nums = [2, 7, 11, 15], target = 9
//
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]

func main() {
	sum1 := twoSum1([]int{7, 6, 8, 3}, 11)
	log.Println("sum1:", sum1)
	sum2 := twoSum2([]int{7, 6, 8, 3}, 11)
	log.Println("sum2:", sum2)
	sum3 := twoSum3([]int{7, 6, 8, 3}, 11)
	log.Println("sum3:", sum3)
}

//O(n^2)
func twoSum1(nums []int, target int) []int {
	var out = make([]int, 2)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				out[0] = i
				out[1] = j
			}
		}
	}
	return out
}

//O(n)
func twoSum2(nums []int, target int) []int {
	var out []int
	firstValueMap := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		firstValueMap[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		sub := target - nums[i]
		v, ok := firstValueMap[sub]
		if ok && v != i {
			out = []int{i, v}
		}
	}
	return out
}

//O(n)
func twoSum3(nums []int, target int) []int {
	var out []int
	firstValueMap := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		sub := target - nums[i]
		v, ok := firstValueMap[sub]
		if ok && v != i {
			out = []int{i, v}
			break
		}
		firstValueMap[nums[i]] = i
	}
	return out
}
