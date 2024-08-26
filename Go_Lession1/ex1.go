package main

import "fmt"

func Exercise1() {
	var a, b float64

	fmt.Print("Input value a: ")
	fmt.Scan(&a)
	fmt.Print("Input value b: ")
	fmt.Scan(&b)

	fmt.Println("Area of rectangle: ", calculate_area_reactangle(a, b))
	fmt.Println("Perimeter of rectangle: ", calculate_perimeter_rectangle(a, b))
}

func calculate_area_reactangle(a float64, b float64) float64 {
	return a * b
}

func calculate_perimeter_rectangle(a float64, b float64) float64 {
	return (a + b) * 2
}
