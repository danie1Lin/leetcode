package main

func runningSum(nums []int) []int {
	sum := make([]int, len(nums))
	if len(nums) == 0 {
		return sum
	}
	sum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sum[i] = sum[i-1] + nums[i]
	}
	return sum
}
