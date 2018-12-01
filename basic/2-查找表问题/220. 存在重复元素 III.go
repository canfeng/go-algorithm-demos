package __查找表问题

import "math"

/*

给定一个整数数组，判断数组中是否有两个不同的索引 i 和 j，使得 nums [i] 和 nums [j] 的差的绝对值最大为 t，并且 i 和 j 之间的差的绝对值最大为 ķ。

示例 1:

输入: nums = [1,2,3,1], k = 3, t = 0
输出: true
示例 2:

输入: nums = [1,0,1,1], k = 1, t = 2
输出: true
示例 3:

输入: nums = [1,5,9,1,5,9], k = 2, t = 3
输出: false
 */
func main() {
	println(containsNearbyAlmostDuplicate([]int{2,1}, 1, 1))
	println(containsNearbyAlmostDuplicate([]int{-1, -1}, 1, 0))
	println(containsNearbyAlmostDuplicate([]int{1, 2, 3, 1}, 3, 0))
	println(containsNearbyAlmostDuplicate([]int{1, 0, 1, 1}, 1, 2))
	println(containsNearbyAlmostDuplicate([]int{1, 5, 9, 1, 5, 9}, 2, 3))
}

//时间复杂度：O(nlogn)
//空间复杂度：O(k)
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	//使用key，存放连续k个元素的值
	smap := map[int]int{}
	//表示 当前滑动窗口的最小索引
	start := 0
	for i := 0; i < len(nums); i++ {
		//取出 
		if math.Abs(float64(nums[i]-value)) <= float64(t) {
			return true
		}

		smap[nums[i]] = nums[i]

		if len(smap) == k+1 {
			delete(smap, nums[start])
			//此时smap的长度为k+1，索引的最大差值为k，需要将第一个元素删除
			start++
		}
	}
	return false
}
