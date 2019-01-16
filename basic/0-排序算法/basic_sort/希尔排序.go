package basic_sort

import "sort"

func ShellSort(arr sort.IntSlice) {

	n := arr.Len()
	step := n / 2 //步长

	for step > 0 {
		for k := step; k < n; k++ {
			for i := k; i >= step && arr.Less(i, i-step); i -= step {
				arr.Swap(i, i-step)
			}
		}
		step = step / 2
	}
}
