package main

import "github.com/rs/zerolog/log"

/*
在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

示例 1:

输入: [3,2,1,5,6,4] 和 k = 2
输出: 5
示例 2:

输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
输出: 4
说明:

你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。
 */
func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	val := findKthLargest(nums, 3)
	log.Print(val)
}

func findKthLargest(nums []int, k int) int {
	if (nums == null || nums.length < k) return 0;
	k--;
	lo = 0, hi = nums.length - 1;
	for lo < hi {
		int j = partition(nums, lo, hi);
		if (j < k) lo = j + 1;
		else if (j > k) hi = j - 1;
		else return nums[k];
	}
	return nums[lo];
}

func partition(nums [] int, len int, left int, right int) int {
	pivot := nums[left]

	for left < right {
		for left < right && pivot >= nums[left] {
			right--
		}
		nums[left] = nums[right]
		for left < right && pivot <= nums[left] {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = pivot

	return left
}
