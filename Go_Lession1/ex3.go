package main

import (
	"fmt"
	"sort"
)

func Exercise3() {
	var n int
	var slice []int

	fmt.Print("Input numnber of element: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var num int
		fmt.Print("Input numnber: ")
		fmt.Scan(&num)
		slice = append(slice, num)
	}

	fmt.Println("Sum: ", getSum(slice))
	fmt.Println("Max: ", getMax(slice))
	fmt.Println("Min: ", getMin(slice))
	fmt.Println("Average: ", getAverage(slice))
	fmt.Println("Sorted slice:", sortSlice(slice))
}

func getSum(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

func getMax(slice []int) int {
	max := slice[0]
	for _, v := range slice {
		if v > max {
			max = v
		}
	}
	return max
}

func getMin(slice []int) int {
	min := slice[0]
	for _, v := range slice {
		if v < min {
			min = v
		}
	}
	return min
}

func getAverage(slice []int) float64 {
	return float64(getSum(slice)) / float64(len(slice))
}

func sortSlice(slice []int) []int {
	sort.Ints(slice)
	return slice
}
