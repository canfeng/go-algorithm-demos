package main

import (
	"algorithm-demo/util"
	"log"
	"math/rand"
)

//二分查找法——实现
//给定一个有序数组，返回指定数字所在的索引位置

func main() {
	arr := util.GenerateIntArrayWithSorted(10)
	log.Println("arr:", arr)
	for i := 0; i < len(arr); i++ {
		target := arr[rand.Intn(len(arr))]
		log.Println("target:", target)
		index := binarySearch4(arr, target, 0, len(arr))
		log.Println("target index:", index)
	}
}

//时间复杂度：O(logn)
//空间复杂度：O(1)
func binarySearch1(nums []int, target int) int {
	len := len(nums)
	//定义左右边界，意义是 从[l...r]的闭区间内找到target的位置
	l, r := 0, len-1

	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

//时间复杂度：O(logn)
func binarySearch2(nums []int, target int) int {
	len := len(nums)
	//定义左右边界，意义是 从[l...r)的闭区间内找到target的位置
	l, r := 0, len

	for l < r {
		mid := (l + r) / 2 //当数组长度足够大时，可能会导致整型溢出
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return -1
}

//时间复杂度：O(logn)
func binarySearch3(nums []int, target int) int {
	len := len(nums)
	//定义左右边界，意义是 从[l...r)的闭区间内找到target的位置
	l, r := 0, len

	for l < r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return -1
}

//时间复杂度：O(logn) 递归实现
func binarySearch4(nums []int, target int, l int, r int) int {
	mid := (l + r) / 2
	if nums[mid] == target {
		return mid
	}
	if nums[mid] > target {
		r = mid
	} else {
		l = mid + 1
	}
	return binarySearch4(nums, target, l, r)
}
