package util

import (
	"math/rand"
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
func GenerateRandomIntArray(size int,rangeMax int) []int {
	var arr []int
	for i := 0; i < size; i++ {
		arr = append(arr, rand.Intn(rangeMax))
	}
	return arr
}
//生成指定长度的有序int数组
func GenerateRandomIntArrayWithSorted(size int,rangeMax int) []int {
	var arr []int
	for i := 0; i < size; i++ {
		arr = append(arr, rand.Intn(rangeMax))
	}
	SortArray(arr)
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


