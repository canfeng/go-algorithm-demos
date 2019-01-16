package __查找表问题

import (
	"algorithm-demo/util"
	"fmt"
)

/*
给定两个数组，编写一个函数来计算它们的交集。

示例 1:

输入: nums1 = [1,2,2,1], nums2 = [2,2]
输出: [2]
示例 2:

输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出: [9,4]
说明:

输出结果中的每个元素一定是唯一的。
我们可以不考虑输出结果的顺序。
*/
func main() {
	//arr := []int{12,33,22}
	//println([]int{123})
	//fmt.Printf("%v\n",arr)
	//fmt.Printf("%t\n",arr)
	//println(reflect.TypeOf(arr))

	arr := intersection([]int{1, 2, 2, 1}, []int{2, 2})
	fmt.Printf("%v\n", arr)
}

func intersection(nums1 []int, nums2 []int) []int {
	set := util.NewSet(nums1)

	resSet := util.NewSet()

	for _, value := range nums2 {
		if set.Contains(value) {
			resSet.Add(value)
		}
	}

	var arr []int
	for _, va := range resSet.ToArray() {
		if value, ok := va.(int); ok {
			arr = append(arr, value)
		}
	}
	return arr
}
