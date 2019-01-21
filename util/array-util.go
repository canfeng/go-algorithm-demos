package util

import (
	"math/rand"
	"time"
	"sort"
	"fmt"
	"math"
)

//生成指定长度的有序int数组
func GenerateIntArrayWithSorted(size int) []int {
	var arr []int
	for i := 0; i < size; i++ {
		arr = append(arr, i)
	}
	return arr
}

//生成指定长度的有序int数组
func GenerateRandomIntArray(size int, rangeMax int) []int {
	var arr []int
	for i := 0; i < size; i++ {
		arr = append(arr, rand.Intn(rangeMax))
	}
	return arr
}

//生成指定长度和取值范围的随机数组
func GenerateRandomArray(size int, rangeMin int, rangeMax int) sort.IntSlice {
	fmt.Printf("GenerateRandomArray() - size:%d - range:[%d~%d] \n", size, rangeMin, rangeMax)
	var arr sort.IntSlice
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		arr = append(arr, rand.Int()%(rangeMax-rangeMin+1)+rangeMin)
	}
	return arr
}

// 生成指定长度和取值范围的近乎有序的数组
// 首先生成一个含有[0...n-1]的完全有序数组, 之后随机交换swapTimes对数据
// swapTimes定义了数组的无序程度:
// swapTimes == 0 时, 数组完全有序
// swapTimes 越大, 数组越趋向于无序
func GenerateNearlyOrderedArray(size int, rangeMin int, rangeMax int, swapTimes int) sort.IntSlice {
	fmt.Printf("GenerateNearlyOrderedArray() - size:%d - range:[%d~%d] - swapTimes:%d \n", size, rangeMin, rangeMax, swapTimes)
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = i
	}
	for ; swapTimes > 0; swapTimes-- {
		posX := rand.Int()%(rangeMax-rangeMin+1) + rangeMin
		posY := rand.Int()%(rangeMax-rangeMin+1) + rangeMin
		arr[posX], arr[posY] = arr[posY], arr[posX]
	}
	return arr
}

//数组排序
func SortArray(in []int) {
	for i := 0; i < len(in); i++ {
		for j := i; j < len(in); j++ {
			if in[i] > in[j] {
				in[i], in[j] = in[j], in[i]
			}
		}
	}
}

/**
反转数组
*/
func ReverseArray(arr []int) {
	len := len(arr)
	for i := 0; i < len/2; i++ {
		arr[i], arr[len-1-i] = arr[len-1-i], arr[i]
	}
}

//检验排序算法排序之后的结果正确性，以及算法运行的时间
func TestSort(sortName string, sortFunc func(arr sort.IntSlice), arr sort.IntSlice) {
	startTime := time.Now().UnixNano()
	sortFunc(arr)
	endTime := time.Now().UnixNano()

	if !sort.IsSorted(arr) {
		panic("array is not sorted ==>" + sortName)
	}
	fmt.Printf("%s : %f s \n", sortName, float64(endTime-startTime)/math.Pow10(9))
}
