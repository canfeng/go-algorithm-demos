package advance_sort

import (
	"math/rand"
	"sort"
)

//快速排序采用“分而治之、各个击破”的观念，此为原地（In-place）分割版本。
//快速排序使用分治法（Divide and conquer）策略来把一个序列（list）分为两个子序列（sub-lists）。
//
//步骤为：
//
//1. 从数列中挑出一个元素，称为“基准”（pivot），
//2. 重新排序数列，所有比基准值小的元素摆放在基准前面，所有比基准值大的元素摆在基准后面（相同的数可以到任何一边）。在这个分割结束之后，该基准就处于数列的中间位置。这个称为分割（partition）操作。
//3. 递归地（recursively）把小于基准值元素的子数列和大于基准值元素的子数列排序。
//递归到最底部时，数列的大小是零或一，也就是已经排序好了。这个算法一定会结束，因为在每次的迭代（iteration）中，它至少会把一个元素摆到它最后的位置去。
//TODO 取第一个元素为基准点，当数组是完全倒序的时候，
//TODO partition 之后的分为 <v 和 >=v 两部分，当有大量 =v 的元素时，会导致左右两边分布不平衡
func QuickSort(arr sort.IntSlice) {
	quickSort(arr, 0, arr.Len()-1)
}

func quickSort(arr sort.IntSlice, l int, r int) {
	if l >= r {
		return
	}
	//基准
	v := arr[l]
	//arr[l+1...j] < v ; arr[j+1...i) > vn
	j := l
	for i := l + 1; i <= r; i++ {
		if arr[i] < v {
			arr.Swap(j+1, i)
			j++
		}
	}
	arr.Swap(l, j)
	quickSort(arr, l, j-1)
	quickSort(arr, j+1, r)
}

// 双路快排
//todo 优化1==>随机选取基准的位置,使得基准左右两边的元素个数保持平衡
//todo 优化2==>将等于基准点的元素，分散到<v 和 >v的两边
func QuickSort2Ways(arr sort.IntSlice) {
	quickSort2Ways(arr, 0, arr.Len()-1)
}

func quickSort2Ways(arr sort.IntSlice, l int, r int) {
	if l >= r {
		return
	}
	//随机在arr[l...r]的范围中, 选择一个数值作为标定点pivot
	arr.Swap(l, rand.Int()%(r-l+1)+l)
	v := arr[l]
	//arr[l+1...i) <= v ; arr(j...r] >= v
	i := l + 1
	j := r
	for {
		for i <= r && arr[i] < v {
			i++
		}
		for j >= l+1 && arr[j] > v {
			j--
		}
		if i > j {
			break
		}
		arr.Swap(i, j)
		i++
		j--
	}
	arr.Swap(l, j)
	quickSort2Ways(arr, l, j-1)
	quickSort2Ways(arr, j+1, r)
}

// 三路快排
//todo 优化1==>将数组分为<v; =v; >v 三个部分，然后分别对 <v 和 >v 的部分继续排序
func QuickSort3Ways(arr sort.IntSlice) {
	quickSort3Ways(arr, 0, arr.Len()-1)
}

func quickSort3Ways(arr sort.IntSlice, l int, r int) {
	if l >= r {
		return
	}
	arr.Swap(l, rand.Int()%(r-l+1)+l)
	v := arr[l] //基准点

	lt := l     //arr[l+1...lt] < v
	gt := r + 1 //arr[gt...r] > v
	// arr[lt+1...i) == v
	for i := l + 1; i < gt; {
		if arr[i] < v {
			//将arr[i] 和 arr[l+1...lt] 之后的第一个元素交换位置
			arr.Swap(i, lt+1)
			lt++
			i++
		} else if arr[i] > v {
			//将arr[i] 和 arr[gt...r] 之前的第一个元素交换位置
			arr.Swap(i, gt-1)
			gt--
		} else {
			//arr[i] == v
			i++
		}
	}
	//最后把基准点放到正确的位置，就是等于v的区间的前一个位置，arr[lt+1...i) 之前
	arr.Swap(l, lt)

	//将小于v和大于v的部分分别快速排序
	quickSort3Ways(arr, l, lt-1)
	quickSort3Ways(arr, gt, r)
}
