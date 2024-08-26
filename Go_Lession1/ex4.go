package main

import "fmt"

func Exercise4() {
	//test
	var slice = []int{2, 3, 5, 7, 8, 10}
	target := 10

	fmt.Print("Result: ", twoSum(slice, target))
}

func twoSum(nums []int, target int) []int {
	var m = make(map[int]int)
	for i, v := range nums {
		m[v] = i
	}

	for i, v := range nums {
		complement := target - v
		if j, exist := m[complement]; exist && j != i {
			return []int{i, j}
		}
	}
	return nil
}
