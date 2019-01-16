package main

import (
	"testing"
	"fmt"
	"algorithm-demo/util"
	"time"
	"algorithm-demo/basic/0-排序算法/basic_sort"
	"sort"
)

func Test_SelectionSort(t *testing.T) {
	nums := util.GenerateRandomArray(1000, 0, 1000)
	basic_sort.SelectionSort(nums)
	fmt.Println(nums)
}

func Test_InsertionSort_SelectionSort(t *testing.T) {
	arr1 := util.GenerateNearlyOrderedArray(10000, 0, 10000, 100)
	arr2 := make(sort.IntSlice, len(arr1))
	copy(arr2, arr1)

	util.TestSort("SelectionSort", basic_sort.SelectionSort, arr1)
	util.TestSort("InsertionSort", basic_sort.InsertionSort_1, arr2)
}

func Test_InsertionSort(t *testing.T) {
	arr1 := util.GenerateRandomArray(10, 0, 10000)
	arr2 := make(sort.IntSlice, len(arr1))
	copy(arr2, arr1)

	util.TestSort("InsertionSort_1", basic_sort.InsertionSort_1, arr1)
	util.TestSort("InsertionSort_2", basic_sort.InsertionSort_2, arr2)
}

func Test_BubbleSort(t *testing.T) {
	arr1 := util.GenerateRandomArray(50000, 0, 10000)
	util.TestSort("BubbleSort", basic_sort.BubbleSort, arr1)
}

func Test_ShellSort(t *testing.T) {
	arr1 := util.GenerateRandomArray(50000, 0, 10000)
	util.TestSort("ShellSort", basic_sort.ShellSort, arr1)
}

func Test_Time(t *testing.T) {
	//获取毫秒
	fmt.Println(time.Microsecond)

	//获取月
	fmt.Println(time.Month(1))

	//当前时间
	fmt.Println(time.Now())
	fmt.Println(time.Now().String())
	//当前时间-小时
	fmt.Println(time.Now().Hour())

	//当前时间unix时间戳since 1970 -1- 1
	fmt.Println(time.Now().Unix())

	//当前时间unix时间戳(nanoseconds),since 1970 -1- 1,  1 seconds = 10^9 nanoseconds
	fmt.Println(time.Now().UnixNano())

	//当前时间加三个小时
	fmt.Println(time.Now().Add(time.Hour * 3))

	//时间戳转化成时间
	currentTime := time.Now().Unix()
	tm := time.Unix(currentTime, 0)
	fmt.Println(tm)
}

func Test_ASSERT(t *testing.T) {
	var a interface{} = Name{"szg"}
	i,ok := a.(Name)
	if ok {
		i.value = "ssss"
		fmt.Println(a)
		fmt.Println(i)
	}
}

type Name struct {
	value string
}