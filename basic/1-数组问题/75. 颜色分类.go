package __数组问题

import (
	"errors"
	"github.com/rs/zerolog/log"
)

/*
给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

注意:
不能使用代码库中的排序函数来解决这道题。

示例:

输入: [2,0,2,1,1,0]
输出: [0,0,1,1,2,2]
进阶：

一个直观的解决方案是使用计数排序的两趟扫描算法。
首先，迭代计算出0、1 和 2 元素的个数，然后按照0、1、2的排序，重写当前数组。
你能想出一个仅使用常数空间的一趟扫描算法吗？
 */
func main() {
	arr := []int{2, 1, 1, 2, 0, 0, 1, 1,2}
	sortColors_2(arr)
	log.Print(arr)
}

//时间复杂度O(n)
//空间复杂度O(n)
//但是需要遍历多次
func sortColors(nums []int) {
	var count [3]int
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 || nums[i] > 2 {
			panic(errors.New("error"))
		}
		count[nums[i]]++
	}

	index := 0
	for k := 0; k < len(count); k++ {
		for i := 0; i < count[k]; i++ {
			nums[index] = k
			index++
		}
	}
}

//三路快排
//时间复杂度:O(n)
//空间复杂度:O(n)
//只遍历一次
func sortColors_2(nums []int) {
	//标定元素为1，1左边的元素都为0，右边的元素都为2
	zero := -1       //表示在[0...zero]的元素都是0
	two := len(nums) //表示在[two...n-1]的元素都是2

	for i := 0; i < two; {
		if nums[i] == 1 {
			i++
		} else if nums[i] == 2 {
			two--
			nums[i], nums[two] = nums[two], nums[i] //i和two索引上的元素交换
		} else if nums[i] == 0 {
			zero++
			nums[i],nums[zero] = nums[zero],nums[i]
			i++
		}
	}

}
