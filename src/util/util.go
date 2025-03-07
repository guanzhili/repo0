package util

import (
	"fmt"
	"log"
	"os"
)

func countChars(str string) map[rune]int {
	countMap := make(map[rune]int)
	for _, char := range str {
		countMap[char]++
	}
	return countMap
}