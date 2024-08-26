package main

import "fmt"

func Exercise2() {
	var s string
	fmt.Print("Input string: ")
	fmt.Scan(&s)

	fmt.Print("Check string Length: ", check_stringLength(s))
}

func check_stringLength(s string) bool {
	return len(s)%2 == 0
}
