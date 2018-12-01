package __数组问题

import "log"

/**
给定一个已按照升序排列 的有序数组，找到两个数使得它们相加之和等于目标数。

函数应该返回这两个下标值 index1 和 index2，其中 index1 必须小于 index2。

说明:

返回的下标值（index1 和 index2）不是从零开始的。
你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。
示例:

输入: numbers = [2, 7, 11, 15], target = 9
输出: [1,2]
解释: 2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。
 */
func main() {
	nums := []int{4, 7, 11, 15, 16, 22, 190}
	target := 26
	sum := twoSum_3(nums, target)
	log.Println(sum)
}

/**
暴力破解法
时间复杂度O(n^2)
 */
func twoSum_1(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		for k := i + 1; k < len(numbers); k++ {
			if numbers[i]+numbers[k] == target {
				return []int{i + 1, k + 1}
			}
		}
	}
	panic("not found correct solution")
}

/**
二分查找法
 */
func twoSum_2(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		//如果
		k := binarySearch(numbers, target-numbers[i], i)
		if k != -1 {
			return []int{i + 1, k + 1}
		}
	}
	panic("not found correct solution")
}

//二分查找法
//时间复杂度 O(nlogn)
//空间复杂度 O(1)
func binarySearch(nums []int, target int, index int) int {
	min := 0
	max := len(nums) - 1
	for min <= max {
		mid := min + (max-min)/2
		if nums[mid] == target && mid != index {
			return mid
		} else if nums[mid] > target {
			max = mid - 1
		} else {
			min = mid + 1
		}
	}
	return -1
}

//对撞指针法
//时间复杂度 O(n)
//空间复杂度 O(1)
func twoSum_3(numbers []int, target int) []int {
	x := 0
	y := len(numbers) - 1

	for x < y { //因为题目要求不能使用同一元素
		sum := numbers[x] + numbers[y]
		if sum == target {
			return []int{x + 1, y + 1}
		} else if sum > target {
			y --
		} else {
			x ++
		}
	}
	panic("not found correct solution")
}
