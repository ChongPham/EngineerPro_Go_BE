package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	myMap := make(map[int]int)

	for i, v := range nums {
		complement := target - v
		if j, exist := myMap[complement]; exist && j != i {
			return []int{i, j}
		}
		myMap[v] = i
	}

	return nil
}

func Exercise3() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 10
	fmt.Println(twoSum(nums, target))
}
