package main

import (
	"fmt"
	"unicode"
)

func CountOccurr(s string) map[rune]int {
	countMap := make(map[rune]int)

	for _, v := range s {
		// remove whitespace
		if !unicode.IsSpace(v) {
			countMap[v] += 1
		}
	}
	return countMap
}

func Exercise2() {
	s := "I'm a Superman, Superman is me. Supperman, I will save you from this bug"
	rs := CountOccurr(s)

	for char, count := range rs {
		fmt.Printf("%c: %d\n", char, count)
	}
}
