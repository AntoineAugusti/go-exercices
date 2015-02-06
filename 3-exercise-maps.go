package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	counts := make(map[string]int)

	for _, word := range words {
		counts[strings.ToLower(word)]++
	}

	return counts
}

func main() {
	fmt.Println(WordCount("Hello Bob bob awesome!"))
}
