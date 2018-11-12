package main

/*
给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，使得 nums [i] = nums [j]，并且 i 和 j 的差的绝对值最大为 k。

示例 1:

输入: nums = [1,2,3,1], k = 3
输出: true

示例 2:

输入: nums = [1,0,1,1], k = 1
输出: true

示例 3:

输入: nums = [1,2,3,1,2,3], k = 2
输出: false
 */
func main() {
	println(containsNearbyDuplicate([]int{1,2,3,1},3))
	println(containsNearbyDuplicate([]int{1,0,1,1},1))
	println(containsNearbyDuplicate([]int{1,2,3,1,2,3},2))
}

/*
使用滑动窗口加查找表的方式实现
时间复杂度：O(n)
空间复杂度：O(k)
 */
func containsNearbyDuplicate(nums []int, k int) bool {
	//使用key，存放连续k个元素的值
	smap := map[int]int{}
	//表示 当前滑动窗口的最小索引
	start :=0
	for i := 0; i < len(nums); i++ {
		if _, ok := smap[nums[i]]; ok {
			return true
		}

		smap[nums[i]] = 1

		if len(smap) == k+1 {
			delete(smap, nums[start])
			//此时smap的长度为k+1，索引的最大差值为k，需要将第一个元素删除
			start++
		}
	}
	return false
}
