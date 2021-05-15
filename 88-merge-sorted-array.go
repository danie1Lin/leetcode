package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := 0
	for len(nums2) != 0 {
		// 考慮到 m = 0 的情況，這個要擺前面
		// 其實通常只要把停止條件擺最前面即可
		if i == m { // 不是減 m-1 因為 j 會從 0 開始
			for j, v := range nums2 {
				nums1[i+j] = v
			}
			nums2 = []int{}
			break
		}
		if nums1[i] > nums2[0] {
			num := nums2[0]
			nums2 = nums2[1:]
			// 往後一位
			shift(nums1, i)
			nums1[i] = num
			m++ // 因為 nums1 後面的尾巴變長了
		}
		i++
	}
}

func shift(array []int, n int) {
	for i := len(array) - 2; i >= n; i-- {
		array[i+1] = array[i]
	}
}

func main() {
	nums1 := []int{1, 3, 5, 7, 9, 0, 0, 0, 0}
	nums2 := []int{2, 4, 6, 8}
	merge(nums1, 5, nums2, len(nums2))
	fmt.Println(nums1)

	nums1 = []int{0}
	nums2 = []int{1}
	merge(nums1, 0, nums2, len(nums2))
	fmt.Println(nums1)
}
