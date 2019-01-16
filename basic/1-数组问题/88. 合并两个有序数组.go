package __数组问题

import "github.com/rs/zerolog/log"

/*

给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。

说明:

初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
示例:

输入:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

输出: [1,2,2,3,5,6]
*/
func main() {
	//case 1
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 3, 5}
	//case 2
	//nums1 := []int{2, 2, 3, 4, 0, 0, 0, 0, 0}
	//nums2 := []int{1, 3, 5, 6, 6}
	merge(nums1, 3, nums2, 3)
	log.Print(nums1)
}

//归并排序
//时间复杂度 O(n)
//空间复杂度 O(1)
func merge(nums1 []int, m int, nums2 []int, n int) {
	k := m + n
	for m > 0 || n > 0 {
		if m <= 0 {
			nums1[k-1] = nums2[n-1]
			n--
		} else if n <= 0 {
			nums1[k-1] = nums1[m-1]
			m--
		} else if nums1[m-1] < nums2[n-1] {
			nums1[k-1] = nums2[n-1]
			n--
		} else {
			nums1[k-1] = nums1[m-1]
			m--
		}
		k--
	}

}

//归并排序
//时间复杂度：O(n)
//空间复杂度：O(n)
func merge_2(nums1 []int, m int, nums2 []int, n int) {
	var temp []int
	i := 0
	j := 0
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			temp = append(temp, nums1[i])
			i++
		} else {
			temp = append(temp, nums2[j])
			j++
		}
	}
	for ; i < m; i++ {
		temp = append(temp, nums1[i])
	}

	for ; j < n; j++ {
		temp = append(temp, nums2[j])
	}

	copy(nums1, temp)
}
