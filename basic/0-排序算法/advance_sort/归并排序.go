package advance_sort

import (
	"algorithm-demo/basic/0-排序算法/basic_sort"
	"sort"
)

//归并排序

/********************************递归法（Top-down）**********************************/

//Divide: 把长度为 n 的 array 分成两半。
//Conquer: 把左半边 array 及右半边 array 分别以 Merge Sort 进行 sorting。 (recursive)
//Combine: merge 左半边及右半边 sorted array。

//1. 申请空间，使其大小为两个已经排序序列之和，该空间用来存放合并后的序列
//2. 设定两个指针，最初位置分别为两个已经排序序列的起始位置
//3. 比较两个指针所指向的元素，选择相对小的元素放入到合并空间，并移动指针到下一位置
//4. 重复步骤3直到某一指针到达序列尾
//5. 将另一序列剩下的所有元素直接复制到合并序列尾

//时间复杂度：O(nlogn)
//空间复杂度：O(n)
func MergeSort_TopDown(arr sort.IntSlice) {
	n := arr.Len()
	if n <= 16{
		basic_sort.InsertionSort_1(arr)
		return
	}
	//[l,r]
	mergeSort_TopDown(arr, 0, n-1)
}
func mergeSort_TopDown(arr sort.IntSlice, L int, R int) {
	//[l,r]
	if L >= R {
		return
	}
	mid := (R-L)/2 + L
	mergeSort_TopDown(arr, L, mid)
	mergeSort_TopDown(arr, mid+1, R)
	merge(arr, L, mid, R)
}

// 将arr[l...mid]和arr[mid+1...r]两部分进行归并
func merge(arr sort.IntSlice, left int, mid int, right int) {
	//构建临时数组，复制从left到right的元素
	tempLen := right - left + 1
	tempArr := make(sort.IntSlice, tempLen)
	for i := left; i <= right; i++ {
		tempArr[i-left] = arr[i]
	}
	// 分别代表左右两个数组的当前索引位置
	leftIndex := 0
	rightIndex := mid + 1 - left
	for i := left; i <= right; i++ {
		if leftIndex > mid-left {
			arr[i] = tempArr[rightIndex]
			rightIndex++
		} else if rightIndex > right-left {
			arr[i] = tempArr[leftIndex]
			leftIndex++
		} else if tempArr[leftIndex] < tempArr[rightIndex] {
			arr[i] = tempArr[leftIndex]
			leftIndex++
		} else {
			arr[i] = tempArr[rightIndex]
			rightIndex++
		}
	}
}

//********************************递归法2（Top-down）**********************************/
//时间复杂度：O(nlogn)
func MergeSort_TopDown_2(arr sort.IntSlice) {
	//[l,r]
	mergeSort_TopDown2(arr, 0, arr.Len()-1)
}
func mergeSort_TopDown2(arr sort.IntSlice, l int, r int) {
	//[l,r]
	if l >= r {
		return
	}
	mid := (r-l)/2 + l
	mergeSort_TopDown2(arr, l, mid)
	mergeSort_TopDown2(arr, mid+1, r)
	merge2(arr, l, mid, r)
}

// 将arr[l...mid]和arr[mid+1...r]两部分进行归并
func merge2(arr sort.IntSlice, left int, mid int, right int) {
	//[0,mid]
	leftLen := mid - left + 1
	rightLen := right - mid
	leftArr := make(sort.IntSlice, leftLen) //TODO 注意：千万不要使用 arr[0:mid+1] 的方式截取左右数组,因为这样还是指向的原数组,这里需要copy两个新数组
	for i := 0; i < leftLen; i++ {
		leftArr[i] = arr[left+i]
	}
	//[mid+1,right]
	rightArr := make(sort.IntSlice, rightLen)
	for i := 0; i < rightLen; i++ {
		rightArr[i] = arr[mid+1+i]
	}

	//分别代表左右两个数组的当前索引位置
	leftIndex := 0
	rightIndex := 0

	i := left
	for ; leftIndex < leftArr.Len() && rightIndex < rightArr.Len(); i++ {
		if leftArr[leftIndex] < rightArr[rightIndex] {
			arr[i] = leftArr[leftIndex]
			leftIndex++
		} else {
			arr[i] = rightArr[rightIndex]
			rightIndex++
		}
	}
	for leftIndex < leftArr.Len() {
		arr[i] = leftArr[leftIndex]
		i++
		leftIndex++
	}
	for rightIndex < rightArr.Len() {
		arr[i] = rightArr[rightIndex]
		i++
		rightIndex++
	}

}

/**************** 迭代法（Bottom-up）****************/
//原理如下（假设序列共有 n个元素）：
//
//将序列每相邻两个数字进行归并操作，形成 ceil(n/2)个序列，排序后每个序列包含两/一个元素
//若此时序列数不是1个则将上述序列再次归并，形成 ceil(n/4)个序列，每个序列包含四/三个元素
//重复步骤2，直到所有元素排序完毕，即序列数为1
func MergeSort_BottomUp(arr sort.IntSlice) {
	n := arr.Len()
	for step := 1; step <= n; step *= 2 {
		for i := 0; i < n-step; i += step * 2 {
			left := i
			right := i + step*2 - 1
			if right > n-1 {
				right = n - 1
			}
			mid := i + step - 1
			merge(arr, left, mid, right)
		}
	}
}
