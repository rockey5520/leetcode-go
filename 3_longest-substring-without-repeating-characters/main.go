package main

import (
	"log"
)

func main() {
	log.Println(lengthOfLongestSubstring("abcabcbb"))
}

func lengthOfLongestSubstring(s string) int {
	var res int
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if isDistinct(s, i, j) {
				res = (max(res, j-i+1))
			}
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func isDistinct(s string, i int, j int) bool {
	visited := make(map[string]bool)
	for k := i; k < j+1; k++ {
		_, ok := visited[string(s[k])]
		if ok {
			return false
		}
		visited[string(s[k])] = true
	}

	return true
}
