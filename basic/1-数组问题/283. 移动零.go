package __数组问题

import (
	"algorithm-demo/util"
	"log"
)

//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//
//示例:
//
//输入: [0,1,0,3,12]
//输出: [1,3,12,0,0]
//说明:
//
//必须在原数组上操作，不能拷贝额外的数组。
//尽量减少操作次数。

func main() {
	//arr := []int{0, 1, 0, 3, 12}
	arr1 := util.GenerateRandomIntArray(40, 20)
	log.Println("source arr=>", arr1)
	moveZeroes_3(arr1)
	log.Println("result arr=>", arr1)
}

//时间复杂度O(n)
//空间复杂度O(n)
func moveZeroes_1(nums []int) {
	var tmp []int
	//将非0元素放到临时数组
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			tmp = append(tmp, nums[i])
		}
	}
	//将临时数组复制回原数组的 0-len(tmp)
	for i := 0; i < len(tmp); i++ {
		nums[i] = tmp[i]
	}

	//剩余元素重置为0
	for i := len(tmp); i < len(nums); i++ {
		nums[i] = 0
	}
}

//时间复杂度O(n)
//空间复杂度O(1)
func moveZeroes_2(nums []int) {
	//不使用临时数组，依次将非0元素移动到最后一个0元素位置
	// k表示新的非0元素的索引
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[k] = nums[i]
			k++
		}
	}
	//把剩余位置重置为0
	for i := k; i < len(nums); i++ {
		nums[i] = 0
	}
}

//时间复杂度O(n)
//空间复杂度O(1)
func moveZeroes_3(nums []int) {
	//不使用临时数组，依次将非0元素和0元素交换
	// k表示新的非0元素的索引
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if i != k {
				nums[i], nums[k] = nums[k], nums[i]
			}
			k++
		}
	}
}
