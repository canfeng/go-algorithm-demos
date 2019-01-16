package basic_sort

import "sort"

//插入排序
//时间复杂度：O(n²)，当数组完全有序时，将退化到O(n)的时间复杂度
func InsertionSort_1(nums sort.IntSlice) {
	for i := 0; i < nums.Len(); i++ {
		//从后往前遍历[0,i-1]，一次将arr[i]和比他大的元素交换位置，知道已经处在合适的位置
		for j := i; j > 0 && nums.Less(j, j-1); j-- {
			nums.Swap(j, j-1)
		}
	}
}

//改进(减少交换的次数)
//运行时间基本相当与原来的1/3
func InsertionSort_2(arr sort.IntSlice) {

	for i := 0; i < arr.Len(); i++ {
		//从后往前遍历[0,i-1]，将arr[i]插到合适的位置

		//把arr[i]保存到临时变量temp，然后从后往前依次和temp比较，比之大则赋值到后一个位置
		temp := arr[i]
		j := i
		for ; j > 0 && arr.Less(temp, j); j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = temp
	}
}
