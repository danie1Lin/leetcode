package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/3sum-closest/
// 1. 只要知道最接近 target 的結果
// 2. 本來 3Sum 是加總為零，只要減去 target 即可
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	closet := nums[0] + nums[1] + nums[2] // 不要 = 0 ，因為可能 target 是 0 但你組不出來
	for i := 0; i < len(nums)-2; i++ {
		// 三個時要實現 two pointer 最裡面是 two pointer，最外層只是遍歷
		k := len(nums) - 1
		for j := i + 1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			if sum > target {
				k--
			} else {
				j++
			}
			if Abs(target-sum) < Abs(target-closet) {
				closet = sum
			}
		}
	}
	return closet
}

func Abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

func main() {
	type Case struct {
		input  []interface{}
		answer int
	}
	cases := []Case{
		//{[]interface{}{[]int{-1, 2, 1, -4}, 2}, 2},
		//{[]interface{}{[]int{-1, 2, 1, -4}, 1}, 2},
		{[]interface{}{[]int{0, 2, 1, -3}, 1}, 0},
	}
	for i, v := range cases {
		fmt.Print(i, v.input)
		r := threeSumClosest(v.input[0].([]int), v.input[1].(int))
		fmt.Println("expect", v.answer, "got", r)
		if r == v.answer {
			fmt.Println("pass")
		} else {
			fmt.Println("failed")
		}
	}
}
