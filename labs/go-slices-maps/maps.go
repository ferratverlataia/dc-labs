package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	a := make(map[string]int)
	b := strings.Fields(s)
	for i := 0; i < len(b); i++ {
		a[b[i]]++
	}
	return a
}

func main() {
	wc.Test(WordCount)
}
