package basic_sort

import (
	"sort"
)

//选择排序：
//使用
func SelectionSort(nums sort.IntSlice) {
	for i := 0; i < len(nums); i++ {
		minIndex := i
		//寻找[i,len(nums)) 区间内最小的元素
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		nums.Swap(i, minIndex)
	}
}
