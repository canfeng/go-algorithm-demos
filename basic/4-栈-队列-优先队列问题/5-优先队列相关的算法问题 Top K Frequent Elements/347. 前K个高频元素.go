package main

import (
	"fmt"
	"github.com/Workiva/go-datastructures/queue"
	"sort"
)

/*
给定一个非空的整数数组，返回其中出现频率前 k 高的元素。

示例 1:

输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]
示例 2:

输入: nums = [1], k = 1
输出: [1]
说明：

你可以假设给定的 k 总是合理的，且 1 ≤ k ≤ 数组中不相同的元素的个数。
你的算法的时间复杂度必须优于 O(n log n) , n 是数组的大小。
*/
func main() {
	fmt.Println(topKFrequent_2([]int{1, 1, 1, 1, 2, 2, 3, 4, 4, 4}, 2))
	fmt.Println(topKFrequent_2([]int{1}, 1))
	fmt.Println(topKFrequent_2([]int{1, 1, 4, 4, 6, 6, 8, 8, 8}, 2))
}

type item struct {
	seq int
	num int
}

type itemArr []item

func (arr itemArr) Len() int {
	return len(arr)
}

func (it itemArr) Less(i, j int) bool {
	return it[i].seq > it[j].seq
}

func (it itemArr) Swap(i, j int) {
	it[i], it[j] = it[j], it[i]
}

/*
1. 先用map记录买个num出现的次数
2. 然后将构建一个item结构体，对齐排序，按照seq倒序

时间复杂度：O（nlogn）
空间复杂度：O(n)
*/
func topKFrequent(nums []int, k int) []int {
	dic := make(map[int]int)

	//统计出现频率
	for _, num := range nums {
		dic[num]++
	}
	var arr itemArr

	for key, value := range dic {
		arr = append(arr, item{seq: value, num: key})
	}

	//排序
	sort.Sort(arr)
	fmt.Println(arr)

	var res []int
	for i := 0; i < k; i++ {
		res = append(res, arr[i].num)
	}

	return res
}

/**
使用二维数组保存，第一维保存出现频率， 第二维保存出现同一频率的数字。
然后按照第一维倒序，第二维正序的方式遍历取k个num

时间复杂度：O(n)
时间复杂度：O(n)
*/
func topKFrequent_1(nums []int, k int) []int {
	m := make(map[int]int)
	for _, n := range nums {
		m[n]++
	}
	ret := make([]int, 0)
	res := make([][]int, len(nums)+1)
	for k, v := range m {
		res[v] = append(res[v], k)
	}
	for i := len(res) - 1; i >= 0; i-- {
		for _, n := range res[i] {
			ret = append(ret, n)
			if len(ret) >= k {
				return ret
			}
		}
	}
	return ret
}

/*
优先队列，不用手动排序
/// Time Complexity: O(nlogn)
/// Space Complexity: O(n)
*/
func topKFrequent_2(nums []int, k int) []int {
	m := make(map[int]int)
	for _, n := range nums {
		m[n]++
	}
	if k > len(m) {
		panic("wrong input ")
	}
	fmt.Println(m)

	priorityQueue := queue.NewPriorityQueue(k, false)
	for key, value := range m {
		if priorityQueue.Len() == k {
			peek := priorityQueue.Peek()
			if value > peek.(item).seq {
				priorityQueue.Put(item{seq: value, num: key})
				priorityQueue.Get(1)
			}
		} else {
			priorityQueue.Put(item{seq: value, num: key})
		}
	}

	var res []int
	items, _ := priorityQueue.Get(k)
	fmt.Println(items)

	for i := len(items) - 1; i >= 0; i-- {
		res = append(res, items[i].(item).num)
	}
	return res
}

func (mi item) Compare(other queue.Item) int {
	omi := other.(item)
	if mi.seq > omi.seq {
		return 1
	} else if mi == omi {
		return 0
	}
	return -1
}
