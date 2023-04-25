package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(nums, target)
	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {

	for x := 0; x < len(nums)-1; x++ {
		for j := x + 1; j < len(nums); j++ {
			if (nums[x] + nums[j]) == target {
				return []int{x, j}
			}

		}
	}
	return []int{-1, -1}
}
