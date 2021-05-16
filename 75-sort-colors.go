package main

import "fmt"

// 題目指定要 in-place sort 且不能用 library
// merge sort O(nlogn)
// bubble sort O(n^2)
// counting sort O(n)
func sortColors(nums []int) {
	count := make([]int, 3)
	for _, v := range nums {
		count[v]++
	}

	color := 0
	for i := 0; i < len(nums); {
		if color > 2 {
			break
		}

		if count[color] > 0 {
			nums[i] = color
			count[color]--
			i++ // 要注意不能擺在頭部去遞增
		}

		if count[color] <= 0 {
			color++
		}
	}
}

func main() {
	for _, v := range [][]int{
		[]int{0},
		[]int{2, 0},
	} {
		sortColors(v)
		fmt.Println(v)
	}
}
