package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, n := range nums {
		complement := target - n

		if e, ok := m[complement]; ok {
			return []int{e, i}
		}

		m[nums[i]] = i
	}

	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
