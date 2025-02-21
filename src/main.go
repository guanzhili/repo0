package main

import (
	"fmt"
)

func countChars(str string) map[rune]int {
	countMap := make(map[rune]int)
	for _, char := range str {
		countMap[char]++
	}
	return countMap
}

func main() {
	str := "hello world"
	result := countChars(str)

	for char, count := range result {
		fmt.Printf("字符 '%c' 出现了 %d 次\n", char, count)
	}
}
