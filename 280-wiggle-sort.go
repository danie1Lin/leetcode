package main

import "fmt"

func wiggleSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		if i%2 == 0 {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		} else {
			if nums[i] < nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}
}

func main() {
	for _, v := range [][]int{
		{6, 6, 5, 6, 3, 8},
		{3, 5, 2, 1, 6, 4},
	} {
		fmt.Print(v, "->")
		wiggleSort(v)
		fmt.Println(v)
	}
}
